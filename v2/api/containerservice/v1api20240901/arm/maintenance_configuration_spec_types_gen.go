// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type MaintenanceConfiguration_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of a default maintenance configuration.
	Properties *MaintenanceConfigurationProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &MaintenanceConfiguration_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-09-01"
func (configuration MaintenanceConfiguration_Spec) GetAPIVersion() string {
	return "2024-09-01"
}

// GetName returns the Name of the resource
func (configuration *MaintenanceConfiguration_Spec) GetName() string {
	return configuration.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/maintenanceConfigurations"
func (configuration *MaintenanceConfiguration_Spec) GetType() string {
	return "Microsoft.ContainerService/managedClusters/maintenanceConfigurations"
}

// Properties used to configure planned maintenance for a Managed Cluster.
type MaintenanceConfigurationProperties struct {
	// MaintenanceWindow: Maintenance window for the maintenance configuration.
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// NotAllowedTime: Time slots on which upgrade is not allowed.
	NotAllowedTime []TimeSpan `json:"notAllowedTime"`

	// TimeInWeek: If two array entries specify the same day of the week, the applied configuration is the union of times in
	// both entries.
	TimeInWeek []TimeInWeek `json:"timeInWeek"`
}

// Maintenance window used to configure scheduled auto-upgrade for a Managed Cluster.
type MaintenanceWindow struct {
	// DurationHours: Length of maintenance window range from 4 to 24 hours.
	DurationHours *int `json:"durationHours,omitempty"`

	// NotAllowedDates: Date ranges on which upgrade is not allowed. 'utcOffset' applies to this field. For example, with
	// 'utcOffset: +02:00' and 'dateSpan' being '2022-12-23' to '2023-01-03', maintenance will be blocked from '2022-12-22
	// 22:00' to '2023-01-03 22:00' in UTC time.
	NotAllowedDates []DateSpan `json:"notAllowedDates"`

	// Schedule: Recurrence schedule for the maintenance window.
	Schedule *Schedule `json:"schedule,omitempty"`

	// StartDate: The date the maintenance window activates. If the current date is before this date, the maintenance window is
	// inactive and will not be used for upgrades. If not specified, the maintenance window will be active right away.
	StartDate *string `json:"startDate,omitempty"`

	// StartTime: The start time of the maintenance window. Accepted values are from '00:00' to '23:59'. 'utcOffset' applies to
	// this field. For example: '02:00' with 'utcOffset: +02:00' means UTC time '00:00'.
	StartTime *string `json:"startTime,omitempty"`

	// UtcOffset: The UTC offset in format +/-HH:mm. For example, '+05:30' for IST and '-07:00' for PST. If not specified, the
	// default is '+00:00'.
	UtcOffset *string `json:"utcOffset,omitempty"`
}

// Time in a week.
type TimeInWeek struct {
	// Day: The day of the week.
	Day *WeekDay `json:"day,omitempty"`

	// HourSlots: Each integer hour represents a time range beginning at 0m after the hour ending at the next hour
	// (non-inclusive). 0 corresponds to 00:00 UTC, 23 corresponds to 23:00 UTC. Specifying [0, 1] means the 00:00 - 02:00 UTC
	// time range.
	HourSlots []int `json:"hourSlots"`
}

// For example, between 2021-05-25T13:00:00Z and 2021-05-25T14:00:00Z.
type TimeSpan struct {
	// End: The end of a time span
	End *string `json:"end,omitempty"`

	// Start: The start of a time span
	Start *string `json:"start,omitempty"`
}

// For example, between '2022-12-23' and '2023-01-05'.
type DateSpan struct {
	// End: The end date of the date span.
	End *string `json:"end,omitempty"`

	// Start: The start date of the date span.
	Start *string `json:"start,omitempty"`
}

// One and only one of the schedule types should be specified. Choose either 'daily', 'weekly', 'absoluteMonthly' or
// 'relativeMonthly' for your maintenance schedule.
type Schedule struct {
	// AbsoluteMonthly: For schedules like: 'recur every month on the 15th' or 'recur every 3 months on the 20th'.
	AbsoluteMonthly *AbsoluteMonthlySchedule `json:"absoluteMonthly,omitempty"`

	// Daily: For schedules like: 'recur every day' or 'recur every 3 days'.
	Daily *DailySchedule `json:"daily,omitempty"`

	// RelativeMonthly: For schedules like: 'recur every month on the first Monday' or 'recur every 3 months on last Friday'.
	RelativeMonthly *RelativeMonthlySchedule `json:"relativeMonthly,omitempty"`

	// Weekly: For schedules like: 'recur every Monday' or 'recur every 3 weeks on Wednesday'.
	Weekly *WeeklySchedule `json:"weekly,omitempty"`
}

// The weekday enum.
// +kubebuilder:validation:Enum={"Friday","Monday","Saturday","Sunday","Thursday","Tuesday","Wednesday"}
type WeekDay string

const (
	WeekDay_Friday    = WeekDay("Friday")
	WeekDay_Monday    = WeekDay("Monday")
	WeekDay_Saturday  = WeekDay("Saturday")
	WeekDay_Sunday    = WeekDay("Sunday")
	WeekDay_Thursday  = WeekDay("Thursday")
	WeekDay_Tuesday   = WeekDay("Tuesday")
	WeekDay_Wednesday = WeekDay("Wednesday")
)

// Mapping from string to WeekDay
var weekDay_Values = map[string]WeekDay{
	"friday":    WeekDay_Friday,
	"monday":    WeekDay_Monday,
	"saturday":  WeekDay_Saturday,
	"sunday":    WeekDay_Sunday,
	"thursday":  WeekDay_Thursday,
	"tuesday":   WeekDay_Tuesday,
	"wednesday": WeekDay_Wednesday,
}

// For schedules like: 'recur every month on the 15th' or 'recur every 3 months on the 20th'.
type AbsoluteMonthlySchedule struct {
	// DayOfMonth: The date of the month.
	DayOfMonth *int `json:"dayOfMonth,omitempty"`

	// IntervalMonths: Specifies the number of months between each set of occurrences.
	IntervalMonths *int `json:"intervalMonths,omitempty"`
}

// For schedules like: 'recur every day' or 'recur every 3 days'.
type DailySchedule struct {
	// IntervalDays: Specifies the number of days between each set of occurrences.
	IntervalDays *int `json:"intervalDays,omitempty"`
}

// For schedules like: 'recur every month on the first Monday' or 'recur every 3 months on last Friday'.
type RelativeMonthlySchedule struct {
	// DayOfWeek: Specifies on which day of the week the maintenance occurs.
	DayOfWeek *WeekDay `json:"dayOfWeek,omitempty"`

	// IntervalMonths: Specifies the number of months between each set of occurrences.
	IntervalMonths *int `json:"intervalMonths,omitempty"`

	// WeekIndex: Specifies on which week of the month the dayOfWeek applies.
	WeekIndex *RelativeMonthlySchedule_WeekIndex `json:"weekIndex,omitempty"`
}

// For schedules like: 'recur every Monday' or 'recur every 3 weeks on Wednesday'.
type WeeklySchedule struct {
	// DayOfWeek: Specifies on which day of the week the maintenance occurs.
	DayOfWeek *WeekDay `json:"dayOfWeek,omitempty"`

	// IntervalWeeks: Specifies the number of weeks between each set of occurrences.
	IntervalWeeks *int `json:"intervalWeeks,omitempty"`
}

// +kubebuilder:validation:Enum={"First","Fourth","Last","Second","Third"}
type RelativeMonthlySchedule_WeekIndex string

const (
	RelativeMonthlySchedule_WeekIndex_First  = RelativeMonthlySchedule_WeekIndex("First")
	RelativeMonthlySchedule_WeekIndex_Fourth = RelativeMonthlySchedule_WeekIndex("Fourth")
	RelativeMonthlySchedule_WeekIndex_Last   = RelativeMonthlySchedule_WeekIndex("Last")
	RelativeMonthlySchedule_WeekIndex_Second = RelativeMonthlySchedule_WeekIndex("Second")
	RelativeMonthlySchedule_WeekIndex_Third  = RelativeMonthlySchedule_WeekIndex("Third")
)

// Mapping from string to RelativeMonthlySchedule_WeekIndex
var relativeMonthlySchedule_WeekIndex_Values = map[string]RelativeMonthlySchedule_WeekIndex{
	"first":  RelativeMonthlySchedule_WeekIndex_First,
	"fourth": RelativeMonthlySchedule_WeekIndex_Fourth,
	"last":   RelativeMonthlySchedule_WeekIndex_Last,
	"second": RelativeMonthlySchedule_WeekIndex_Second,
	"third":  RelativeMonthlySchedule_WeekIndex_Third,
}