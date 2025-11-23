#!/usr/bin/env bash

echo "🔍 Scanning for compromised NPM packages..."

# Define compromised packages and versions
declare -A compromised=(
  [ansi-regex]="6.2.1"
  [ansi-styles]="6.2.2"
  [backslash]="0.2.1"
  [chalk]="5.6.1"
  [chalk-template]="1.1.1"
  [color-convert]="3.1.1"
  [color-name]="2.0.1"
  [color-string]="2.1.1"
  [debug]="4.4.2"
  [error-ex]="1.3.3"
  [has-ansi]="6.0.1"
  [is-arrayish]="0.3.3"
  [proto-tinker-wc]="0.1.87"
  [simple-swizzle]="0.2.3"
  [slice-ansi]="7.1.1"
  [strip-ansi]="7.1.1"
  [supports-color]="10.2.1"
  [supports-hyperlinks]="4.1.1"
  [wrap-ansi]="9.0.1"
)

declare -a findings=()

scan_file() {
  local file="$1"
  for pkg in "${!compromised[@]}"; do
    version="${compromised[$pkg]}"
    if grep -q "$pkg" "$file" && grep -q "$version" "$file"; then
      findings+=("Found $pkg@$version in $file")
    fi
  done
}

scan_npm_cache() {
  local cache_dir="$1"
  for pkg in "${!compromised[@]}"; do
    version="${compromised[$pkg]}"
    while IFS= read -r match; do
      findings+=("Found $pkg@$version in cached file: $match")
    done < <(find "$cache_dir" -type f -exec grep -l "$pkg@$version" {} + 2>/dev/null)
  done
}

scan_nvm_versions() {
  local nvm_dir="${NVM_DIR:-$HOME/.nvm}"
  if [ ! -d "$nvm_dir" ]; then return; fi

  echo "🧠 Scanning NVM-managed Node versions..."
  count=0
  find "$nvm_dir/versions/node" -type d \( -name "node_modules" -o -name ".npm" \) | while read -r dir; do
    scan_npm_cache "$dir"
    count=$((count + 1))
    if (( count % 5 == 0 )); then
      echo "  ...scanned $count NVM directories"
    fi
  done
}

scan_dockerfiles() {
  echo "🐳 Scanning Dockerfiles..."
  count=0
  while IFS= read -r file; do
    scan_file "$file"
    count=$((count + 1))
    if (( count % 5 == 0 )); then
      echo "  ...scanned $count Dockerfiles"
    fi
  done < <(find . -type f -iname "Dockerfile")
}

scan_ci_configs() {
  echo "⚙️ Scanning CI/CD config files..."
  count=0
  while IFS= read -r file; do
    scan_file "$file"
    count=$((count + 1))
    if (( count % 5 == 0 )); then
      echo "  ...scanned $count CI config files"
    fi
  done < <(find . -type f \( -name "*.yml" -o -name "*.yaml" \) -path "*/.github/*" -o -path "*/.gitlab/*")
}

scan_vendored_dirs() {
  echo "📁 Scanning vendored folders..."
  for dir in vendor third_party static assets; do
    [ -d "$dir" ] || continue
    while IFS= read -r file; do
      scan_file "$file"
    done < <(find "$dir" -type f \( -name "*.js" -o -name "*.json" -o -name "*.tgz" \))
  done
}

scan_lockfile_resolved() {
  local file="$1"
  for pkg in "${!compromised[@]}"; do
    version="${compromised[$pkg]}"
    # Match resolved tarball URLs for compromised versions
    if grep -q "$pkg/-/$pkg-$version.tgz" "$file"; then
      findings+=("Resolved $pkg@$version in $file")
    fi
  done
}

echo "🔒 Scanning project lockfiles and package.json..."
count=0
while IFS= read -r file; do
  scan_lockfile_resolved "$file"  # new accurate match
  count=$((count + 1))
  if (( count % 10 == 0 )); then
    echo "  ...scanned $count files"
  fi
done < <(find . -type f \( -name "package-lock.json" -o -name "yarn.lock" -o -name "pnpm-lock.yaml" \))

# Global caches
echo "📦 Scanning global npm caches..."
[ -d "$HOME/.npm/_cacache" ] && scan_npm_cache "$HOME/.npm/_cacache"
[ -d "$HOME/.npm-packages" ] && scan_npm_cache "$HOME/.npm-packages"

echo "📦 Scanning Yarn global cache..."
if command -v yarn >/dev/null 2>&1; then
  yarn_cache=$(yarn cache dir 2>/dev/null)
  [ -n "$yarn_cache" ] && [ -d "$yarn_cache" ] && scan_npm_cache "$yarn_cache"
fi

echo "📦 Scanning pnpm global store..."
if command -v pnpm >/dev/null 2>&1; then
  pnpm_cache=$(pnpm store path 2>/dev/null)
  [ -n "$pnpm_cache" ] && [ -d "$pnpm_cache" ] && scan_npm_cache "$pnpm_cache"
fi

scan_nvm_versions
scan_dockerfiles
scan_ci_configs
scan_vendored_dirs

echo ""
echo "📊 Summary of Findings:"
if [ ${#findings[@]} -eq 0 ]; then
  echo "✅ No compromised packages found."
else
  declare -A grouped
  declare -A remediation

  # Group findings by remediation directory
  for line in "${findings[@]}"; do
    file=$(echo "$line" | awk -F'in ' '{print $2}')
    dir=$(dirname "$file")

    # Trim path before node_modules if present
    if [[ "$dir" == *"/node_modules/"* ]]; then
      dir="${dir%%/node_modules/*}"
    elif [[ "$dir" == *"/node_modules" ]]; then
      dir="${dir%/node_modules}"
    fi

    # Walk up to find the remediation root
    while [ "$dir" != "/" ]; do
      if [ -f "$dir/package-lock.json" ]; then
        remediation["$dir"]="npm"
        break
      elif [ -f "$dir/yarn.lock" ]; then
        remediation["$dir"]="yarn"
        break
      elif [ -f "$dir/pnpm-lock.yaml" ]; then
        remediation["$dir"]="pnpm"
        break
      fi
      dir=$(dirname "$dir")
    done

    grouped["$dir"]+="$line"$'\n'
  done

  # Print grouped findings
  for dir in "${!grouped[@]}"; do
    echo "📁 $dir"
    echo "${grouped[$dir]}"
  done

  echo ""
  echo "🛠️ Suggested Remediation Commands:"
  for dir in "${!remediation[@]}"; do
    tool="${remediation[$dir]}"
    echo "💡 cd \"$dir\" && rm -rf node_modules ${tool}-lock.yaml yarn.lock package-lock.json && $tool install"
  done
fi
