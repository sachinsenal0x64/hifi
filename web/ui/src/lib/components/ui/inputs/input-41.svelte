<script lang="ts">
	import type { DateValue } from '@internationalized/date';

	import Label from '../ui/label.svelte';
	import { useLocale } from '$lib/hooks/use-locale.svelte';

	import CalendarIcon from '@lucide/svelte/icons/calendar';
	import ChevronLeft from '@lucide/svelte/icons/chevron-left';
	import ChevronRight from '@lucide/svelte/icons/chevron-right';
	import { cn } from '$lib/utils';
	import { DatePicker } from 'bits-ui';

	let value: DateValue | undefined = $state(undefined);
	let locale = useLocale();
</script>

<DatePicker.Root locale={locale.locale} bind:value weekdayFormat="short" fixedWeeks={true}>
	<div class="*:not-first:mt-2">
		<Label class="text-foreground text-sm font-medium">Date picker</Label>
		<div class="flex">
			<div
				class="bg-background ring-offset-background focus-within:border-ring focus-within:ring-ring/30 inline-flex h-9 w-full items-center overflow-hidden rounded-lg border px-3 py-2 pe-9 text-sm whitespace-nowrap shadow-xs shadow-black/[.04] transition-shadow focus-within:ring-2 focus-within:ring-offset-2 focus-within:outline-hidden disabled:opacity-50"
			>
				<DatePicker.Input>
					{#snippet children({ segments })}
						{#each segments as { part, value }, i (part + i)}
							<DatePicker.Segment
								{part}
								class={[
									'text-foreground focus:bg-accent data-invalid:focused:bg-destructive focused:aria-[valuetext=Empty]:text-foreground focused:text-foreground data-invalid:aria-[valuetext=Empty]:text-destructive data-invalid:text-destructive aria-[valuetext=Empty]:text-muted-foreground/70 data-invalid:focused:text-white data-invalid:focused:aria-[valuetext=Empty]:text-white inline rounded p-0.5 caret-transparent outline-hidden disabled:cursor-not-allowed disabled:opacity-50',
									'data-[segment=literal]:text-muted-foreground/70 data-[segment=literal]:px-0'
								]}
							>
								{value}
							</DatePicker.Segment>
						{/each}
					{/snippet}
				</DatePicker.Input>
			</div>
			<DatePicker.Trigger
				class="text-muted-foreground/80 hover:text-foreground data-focus-visible:border-ring data-focus-visible:ring-ring/50 z-10 -ms-9 -me-px flex w-9 items-center justify-center rounded-e-md transition-[color,box-shadow] outline-none data-focus-visible:ring-[3px]"
			>
				<CalendarIcon size={16} />
			</DatePicker.Trigger>
		</div>
		<DatePicker.Content
			sideOffset={6}
			class="border-input bg-background text-foreground z-50 rounded-lg border shadow-lg shadow-black/[.04] outline-hidden"
		>
			<DatePicker.Calendar class="w-fit p-2">
				{#snippet children({ months, weekdays })}
					<DatePicker.Header class="flex w-full items-center gap-1 pb-1">
						<DatePicker.PrevButton
							class="text-muted-foreground/80 ring-offset-background hover:bg-accent hover:text-foreground flex size-9 items-center justify-center rounded-lg transition-shadow"
						>
							<ChevronLeft size={16} />
						</DatePicker.PrevButton>
						<DatePicker.Heading class="grow text-center text-sm font-medium" />
						<DatePicker.NextButton
							class="text-muted-foreground/80 ring-offset-background hover:bg-accent hover:text-foreground flex size-9 items-center justify-center rounded-lg transition-shadow"
						>
							<ChevronRight size={16} />
						</DatePicker.NextButton>
					</DatePicker.Header>
					<div class="flex flex-col space-y-4 pt-4 sm:flex-row sm:space-y-0 sm:space-x-4">
						{#each months as month (month.value)}
							<DatePicker.Grid class="w-fit border-collapse space-y-1 select-none">
								<DatePicker.GridHead>
									<DatePicker.GridRow class="flex w-full justify-between">
										{#each weekdays as day (day)}
											<DatePicker.HeadCell
												class="text-muted-foreground/80 size-9 rounded-lg p-0 text-xs font-medium"
											>
												<div>{day.slice(0, 2)}</div>
											</DatePicker.HeadCell>
										{/each}
									</DatePicker.GridRow>
								</DatePicker.GridHead>
								<DatePicker.GridBody class="[&_td]:px-0">
									{#each month.weeks as weekDates (weekDates)}
										<DatePicker.GridRow class="flex w-full">
											{#each weekDates as date (date)}
												<DatePicker.Cell
													{date}
													month={month.value}
													class="relative size-10 p-0! text-center text-sm"
												>
													<DatePicker.Day
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
														<div
															class="bg-primary data-selected:bg-background absolute start-1/2 bottom-1 hidden size-[3px] -translate-x-1/2 rounded-full transition-all group-data-today:block"
														></div>
														{date.day}
													</DatePicker.Day>
												</DatePicker.Cell>
											{/each}
										</DatePicker.GridRow>
									{/each}
								</DatePicker.GridBody>
							</DatePicker.Grid>
						{/each}
					</div>
				{/snippet}
			</DatePicker.Calendar>
		</DatePicker.Content>
	</div>
</DatePicker.Root>
