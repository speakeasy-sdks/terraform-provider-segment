// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ReplaceAdvancedSyncScheduleForWarehouseV1InputAdvancedWarehouseSyncScheduleV1Input - Defines the advanced sync schedule for a Warehouse.
type ReplaceAdvancedSyncScheduleForWarehouseV1InputAdvancedWarehouseSyncScheduleV1Input struct {
	// A list that contains the times when syncs should occur.
	Times []WarehouseAdvancedSyncV1 `json:"times"`
	// A TZ-database timezone for this sync schedule.
	Timezone string `json:"timezone"`
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseV1InputAdvancedWarehouseSyncScheduleV1Input) GetTimes() []WarehouseAdvancedSyncV1 {
	if o == nil {
		return []WarehouseAdvancedSyncV1{}
	}
	return o.Times
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseV1InputAdvancedWarehouseSyncScheduleV1Input) GetTimezone() string {
	if o == nil {
		return ""
	}
	return o.Timezone
}

// ReplaceAdvancedSyncScheduleForWarehouseV1Input - Replaces the advanced sync schedule for a Warehouse.
type ReplaceAdvancedSyncScheduleForWarehouseV1Input struct {
	// Enable to turn on an advanced sync schedule for the Warehouse.
	Enabled bool `json:"enabled"`
	// The full sync schedule for the Warehouse.
	Schedule *ReplaceAdvancedSyncScheduleForWarehouseV1InputAdvancedWarehouseSyncScheduleV1Input `json:"schedule,omitempty"`
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseV1Input) GetSchedule() *ReplaceAdvancedSyncScheduleForWarehouseV1InputAdvancedWarehouseSyncScheduleV1Input {
	if o == nil {
		return nil
	}
	return o.Schedule
}
