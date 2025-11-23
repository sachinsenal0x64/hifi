<script lang="ts">
	import Label from '$lib/components/ui/label.svelte';
	import { useLocale } from '$lib/hooks/use-locale.svelte';
	import { cn } from '$lib/utils.js';

	import { type DateValue, getLocalTimeZone, isWeekend, today } from '@internationalized/date';
	import Calendar from '@lucide/svelte/icons/calendar';
	import ChevronLeft from '@lucide/svelte/icons/chevron-left';
	import ChevronRight from '@lucide/svelte/icons/chevron-right';
	import { type DateRange, DateRangePicker, type DateRangeValidator } from 'bits-ui';

	let now = today(getLocalTimeZone());
	let value: DateRange = $state({ end: undefined, start: undefined });
	let locale = useLocale();

	// Define disabled date ranges
	const disabledRanges = [
		[now, now.add({ days: 5 })],
		[now.add({ days: 14 }), now.add({ days: 16 })],
		[now.add({ days: 23 }), now.add({ days: 24 })]
	];

	// Check if a date is unavailable
	function isDateUnavailable(date: DateValue) {
		return (
			isWeekend(date, locale.locale) ||
			disabledRanges.some(
				(interval) => date.compare(interval[0]) >= 0 && date.compare(interval[1]) <= 0
			)
		);
	}

	// Validate the selected range
	const validate: DateRangeValidator = (value: DateRange) => {
		if (!value?.start || !value?.end) return;
		const hasOverlap = disabledRanges.some((interval) => {
			const rangeEnd = value.end!.compare(interval[0]) >= 0;
			const rangeStart = value.start!.compare(interval[1]) <= 0;
			return rangeEnd && rangeStart;
		});
		if (!hasOverlap) return;

		return 'Selected date range may not include unavailable dates.';
	};
</script>

<DateRangePicker.Root
	bind:value
	locale={locale.locale}
	minValue={now}
	{validate}
	{isDateUnavailable}
	weekdayFormat="short"
	fixedWeeks={true}
	class="*:not-first:mt-2"
>
	<Label class="text-foreground text-sm font-medium">Date range picker (unavailable dates)</Label>
	<div class="flex">
		<div
			class="border-input bg-background ring-offset-background focus-within:border-ring focus-within:ring-ring/30 inline-flex h-9 w-full items-center overflow-hidden rounded-lg border px-3 py-2 pe-9 text-sm whitespace-nowrap shadow-xs shadow-black/[.04] transition-shadow focus-within:ring-2 focus-within:ring-offset-2 focus-within:outline-hidden disabled:opacity-50"
		>
			{#each ['start', 'end'] as const as type (type)}
				<DateRangePicker.Input {type}>
					{#snippet children({ segments })}
						<!-- eslint-disable-next-line svelte/require-each-key -->
						{#each segments as { part, value }}
							<DateRangePicker.Segment
								{part}
								class={[
									'text-foreground focus:bg-accent data-invalid:focused:bg-destructive focused:aria-[valuetext=Empty]:text-foreground focused:text-foreground data-invalid:aria-[valuetext=Empty]:text-destructive data-invalid:text-destructive aria-[valuetext=Empty]:text-muted-foreground/70 data-invalid:focused:text-white data-invalid:focused:aria-[valuetext=Empty]:text-white inline rounded p-0.5 caret-transparent outline-hidden disabled:cursor-not-allowed disabled:opacity-50',
									'data-[segment=literal]:text-muted-foreground/70  data-[segment=literal]:px-0'
								]}
							>
								{value}
							</DateRangePicker.Segment>
						{/each}
					{/snippet}
				</DateRangePicker.Input>
				{#if type === 'start'}
					<span aria-hidden="true" class="text-muted-foreground/70 px-2">-</span>
				{/if}
			{/each}
		</div>

		<DateRangePicker.Trigger
			class="text-muted-foreground/80 ring-offset-background hover:text-foreground focus-visible:text-foreground data-focus-visible:border-ring data-focus-visible:ring-ring/30 z-10 -ms-9 -me-px flex w-9 items-center justify-center rounded-e-lg transition-shadow focus-visible:outline-hidden data-focus-visible:border data-focus-visible:ring-2 data-focus-visible:ring-offset-2"
		>
			<Calendar size={16} />
		</DateRangePicker.Trigger>
	</div>

	<DateRangePicker.Content
		class="border-input bg-background text-foreground z-50 rounded-lg border shadow-lg shadow-black/[.04] outline-hidden"
	>
		<DateRangePicker.Calendar class="w-fit p-2">
			{#snippet children({ months, weekdays })}
				<!-- Calendar header and navigation -->
				<DateRangePicker.Header class="flex w-full items-center gap-1 pb-1">
					<DateRangePicker.PrevButton
						class="text-muted-foreground/80 ring-offset-background hover:bg-accent hover:text-foreground flex size-9 items-center justify-center rounded-lg transition-shadow"
					>
						<ChevronLeft size={16} />
					</DateRangePicker.PrevButton>
					<DateRangePicker.Heading class="grow text-center text-sm font-medium" />
					<DateRangePicker.NextButton
						class="text-muted-foreground/80 ring-offset-background hover:bg-accent hover:text-foreground flex size-9 items-center justify-center rounded-lg transition-shadow"
					>
						<ChevronRight size={16} />
					</DateRangePicker.NextButton>
				</DateRangePicker.Header>

				<!-- Calendar grid -->
				<div class="flex flex-col space-y-4 pt-4 sm:flex-row sm:space-y-0 sm:space-x-4">
					{#each months as month (month.value)}
						<DateRangePicker.Grid class="w-fit border-collapse space-y-1 select-none">
							<DateRangePicker.GridHead>
								<DateRangePicker.GridRow class="flex w-full justify-between">
									{#each weekdays as day (day)}
										<DateRangePicker.HeadCell
											class="text-muted-foreground/80 size-9 rounded-lg p-0 text-xs font-medium"
										>
											{day.slice(0, 2)}
										</DateRangePicker.HeadCell>
									{/each}
								</DateRangePicker.GridRow>
							</DateRangePicker.GridHead>

							<DateRangePicker.GridBody class="[&_td]:px-0">
								{#each month.weeks as weekDates (weekDates.join('-'))}
									<DateRangePicker.GridRow class="flex w-full">
										{#each weekDates as date (date.day)}
											<DateRangePicker.Cell
												{date}
												month={month.value}
												class={cn(
													'text-foreground ring-offset-background data-focus-visible:border-ring hover:bg-accent data-selected:bg-accent hover:text-foreground data-selected:text-foreground data-focus-visible:ring-ring/30 data-invalid:data-selection-end:[&:not([data-hover])]:bg-destructive data-invalid:data-selection-start:[&:not([data-hover])]:bg-destructive data-selection-end:[&:not([data-hover])]:bg-primary data-selection-start:[&:not([data-hover])]:bg-primary data-invalid:data-selection-end:[&:not([data-hover])]:text-destructive-foreground data-invalid:data-selection-start:[&:not([data-hover])]:text-destructive-foreground data-selection-end:[&:not([data-hover])]:text-primary-foreground data-selection-start:[&:not([data-hover])]:text-primary-foreground relative flex size-9 items-center justify-center rounded-lg border border-transparent p-0 text-sm font-normal whitespace-nowrap [transition-property:border-radius,box-shadow] duration-150 focus-visible:outline-hidden disabled:pointer-events-none disabled:opacity-30 data-focus-visible:z-10 data-focus-visible:ring-2 data-focus-visible:ring-offset-2 data-focus-visible:outline-hidden data-invalid:bg-red-100 data-selected:rounded-none data-selection-end:rounded-e-lg data-selection-start:rounded-s-lg data-unavailable:pointer-events-none data-unavailable:line-through data-unavailable:opacity-30',
													date.compare(now) === 0 &&
														'after:bg-primary data-selection-end:[&:not([data-hover])]:after:bg-background data-selection-start:[&:not([data-hover])]:after:bg-background after:pointer-events-none after:absolute after:start-1/2 after:bottom-1 after:z-10 after:size-[3px] after:-translate-x-1/2 after:rounded-full'
												)}
											>
												<DateRangePicker.Day
													class={cn(
														'text-foreground ring-offset-background relative flex size-9 items-center justify-center rounded-lg border border-transparent p-0 text-sm font-normal whitespace-nowrap [transition-property:border-radius,box-shadow] duration-150',
														'disabled:pointer-events-none data-outside-month:pointer-events-none',
														'data-highlighted:bg-accent data-selected:bg-accent',
														'data-selection-end:bg-primary data-selection-start:bg-primary',
														'data-selection-end:text-primary-foreground data-selection-start:text-primary-foreground',
														'data-highlighted:rounded-none data-selection-end:rounded-e-lg data-selection-start:rounded-s-lg',
														'focus-visible:ring-ring/30 focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden'
													)}
												>
													{date.day}
												</DateRangePicker.Day>
											</DateRangePicker.Cell>
										{/each}
									</DateRangePicker.GridRow>
								{/each}
							</DateRangePicker.GridBody>
						</DateRangePicker.Grid>
					{/each}
				</div>
			{/snippet}
		</DateRangePicker.Calendar>
	</DateRangePicker.Content>
	<p class="text-muted-foreground mt-2 text-xs" role="region" aria-live="polite">
		Built with
		<a
			class="hover:text-foreground underline"
			href="https://next.bits-ui.com/docs/components/date-range-picker"
			target="_blank"
			rel="noopener nofollow"
		>
			Bits UI
		</a>
	</p>
</DateRangePicker.Root>
