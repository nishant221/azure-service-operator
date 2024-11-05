// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServer_Spec struct {
	// Identity: The cmk identity for the server.
	Identity *MySQLServerIdentity `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *MySQLServerSku `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServer_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-06-30"
func (server FlexibleServer_Spec) GetAPIVersion() string {
	return "2023-06-30"
}

// GetName returns the Name of the resource
func (server *FlexibleServer_Spec) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers"
func (server *FlexibleServer_Spec) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers"
}

// Properties to configure Identity for Bring your Own Keys
type MySQLServerIdentity struct {
	// Type: Type of managed service identity.
	Type                   *MySQLServerIdentity_Type              `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Billing information related properties of a server.
type MySQLServerSku struct {
	// Name: The name of the sku, e.g. Standard_D32s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. GeneralPurpose.
	Tier *MySQLServerSku_Tier `json:"tier,omitempty"`
}

// The properties of a server.
type ServerProperties struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// AvailabilityZone: availability Zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup related properties of a server.
	Backup *Backup `json:"backup,omitempty"`

	// CreateMode: The mode to create a new MySQL server.
	CreateMode *ServerProperties_CreateMode `json:"createMode,omitempty"`

	// DataEncryption: The Data Encryption for CMK.
	DataEncryption *DataEncryption `json:"dataEncryption,omitempty"`

	// HighAvailability: High availability related properties of a server.
	HighAvailability *HighAvailability `json:"highAvailability,omitempty"`

	// ImportSourceProperties: Source properties for import from storage.
	ImportSourceProperties *ImportSourceProperties `json:"importSourceProperties,omitempty"`

	// MaintenanceWindow: Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// Network: Network related properties of a server.
	Network *Network `json:"network,omitempty"`

	// ReplicationRole: The replication role.
	ReplicationRole *ReplicationRole `json:"replicationRole,omitempty"`

	// RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime     *string `json:"restorePointInTime,omitempty"`
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// Storage: Storage related properties of a server.
	Storage *Storage `json:"storage,omitempty"`

	// Version: Server version.
	Version *ServerVersion `json:"version,omitempty"`
}

// Storage Profile properties of a server
type Backup struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: Whether or not geo redundant backup is enabled.
	GeoRedundantBackup *EnableStatusEnum `json:"geoRedundantBackup,omitempty"`
}

// The date encryption for cmk.
type DataEncryption struct {
	// GeoBackupKeyURI: Geo backup key uri as key vault can't cross region, need cmk in same region as geo backup
	GeoBackupKeyURI                 *string `json:"geoBackupKeyURI,omitempty"`
	GeoBackupUserAssignedIdentityId *string `json:"geoBackupUserAssignedIdentityId,omitempty"`

	// PrimaryKeyURI: Primary key uri
	PrimaryKeyURI                 *string `json:"primaryKeyURI,omitempty"`
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// Type: The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
	Type *DataEncryption_Type `json:"type,omitempty"`
}

// Network related properties of a server
type HighAvailability struct {
	// Mode: High availability mode for a server.
	Mode *HighAvailability_Mode `json:"mode,omitempty"`

	// StandbyAvailabilityZone: Availability zone of the standby server.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`
}

// Import source related properties.
type ImportSourceProperties struct {
	// DataDirPath: Relative path of data directory in storage.
	DataDirPath *string `json:"dataDirPath,omitempty"`

	// SasToken: Sas token for accessing source storage. Read and list permissions are required for sas token.
	SasToken *string `json:"sasToken,omitempty"`

	// StorageType: Storage type of import source.
	StorageType *ImportSourceProperties_StorageType `json:"storageType,omitempty"`

	// StorageUrl: Uri of the import source storage.
	StorageUrl *string `json:"storageUrl,omitempty"`
}

// Maintenance window of a server.
type MaintenanceWindow struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

// +kubebuilder:validation:Enum={"UserAssigned"}
type MySQLServerIdentity_Type string

const MySQLServerIdentity_Type_UserAssigned = MySQLServerIdentity_Type("UserAssigned")

// Mapping from string to MySQLServerIdentity_Type
var mySQLServerIdentity_Type_Values = map[string]MySQLServerIdentity_Type{
	"userassigned": MySQLServerIdentity_Type_UserAssigned,
}

// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type MySQLServerSku_Tier string

const (
	MySQLServerSku_Tier_Burstable       = MySQLServerSku_Tier("Burstable")
	MySQLServerSku_Tier_GeneralPurpose  = MySQLServerSku_Tier("GeneralPurpose")
	MySQLServerSku_Tier_MemoryOptimized = MySQLServerSku_Tier("MemoryOptimized")
)

// Mapping from string to MySQLServerSku_Tier
var mySQLServerSku_Tier_Values = map[string]MySQLServerSku_Tier{
	"burstable":       MySQLServerSku_Tier_Burstable,
	"generalpurpose":  MySQLServerSku_Tier_GeneralPurpose,
	"memoryoptimized": MySQLServerSku_Tier_MemoryOptimized,
}

// Network related properties of a server
type Network struct {
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneResourceId  *string `json:"privateDnsZoneResourceId,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is 'Disabled' when server
	// has VNet integration.
	PublicNetworkAccess *EnableStatusEnum `json:"publicNetworkAccess,omitempty"`
}

// The replication role.
// +kubebuilder:validation:Enum={"None","Replica","Source"}
type ReplicationRole string

const (
	ReplicationRole_None    = ReplicationRole("None")
	ReplicationRole_Replica = ReplicationRole("Replica")
	ReplicationRole_Source  = ReplicationRole("Source")
)

// Mapping from string to ReplicationRole
var replicationRole_Values = map[string]ReplicationRole{
	"none":    ReplicationRole_None,
	"replica": ReplicationRole_Replica,
	"source":  ReplicationRole_Source,
}

// +kubebuilder:validation:Enum={"Default","GeoRestore","PointInTimeRestore","Replica"}
type ServerProperties_CreateMode string

const (
	ServerProperties_CreateMode_Default            = ServerProperties_CreateMode("Default")
	ServerProperties_CreateMode_GeoRestore         = ServerProperties_CreateMode("GeoRestore")
	ServerProperties_CreateMode_PointInTimeRestore = ServerProperties_CreateMode("PointInTimeRestore")
	ServerProperties_CreateMode_Replica            = ServerProperties_CreateMode("Replica")
)

// Mapping from string to ServerProperties_CreateMode
var serverProperties_CreateMode_Values = map[string]ServerProperties_CreateMode{
	"default":            ServerProperties_CreateMode_Default,
	"georestore":         ServerProperties_CreateMode_GeoRestore,
	"pointintimerestore": ServerProperties_CreateMode_PointInTimeRestore,
	"replica":            ServerProperties_CreateMode_Replica,
}

// The version of a server.
// +kubebuilder:validation:Enum={"5.7","8.0.21"}
type ServerVersion string

const (
	ServerVersion_57   = ServerVersion("5.7")
	ServerVersion_8021 = ServerVersion("8.0.21")
)

// Mapping from string to ServerVersion
var serverVersion_Values = map[string]ServerVersion{
	"5.7":    ServerVersion_57,
	"8.0.21": ServerVersion_8021,
}

// Storage Profile properties of a server
type Storage struct {
	// AutoGrow: Enable Storage Auto Grow or not.
	AutoGrow *EnableStatusEnum `json:"autoGrow,omitempty"`

	// AutoIoScaling: Enable IO Auto Scaling or not.
	AutoIoScaling *EnableStatusEnum `json:"autoIoScaling,omitempty"`

	// Iops: Storage IOPS for a server.
	Iops *int `json:"iops,omitempty"`

	// LogOnDisk: Enable Log On Disk or not.
	LogOnDisk *EnableStatusEnum `json:"logOnDisk,omitempty"`

	// StorageSizeGB: Max storage size allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
}

// +kubebuilder:validation:Enum={"AzureKeyVault","SystemManaged"}
type DataEncryption_Type string

const (
	DataEncryption_Type_AzureKeyVault = DataEncryption_Type("AzureKeyVault")
	DataEncryption_Type_SystemManaged = DataEncryption_Type("SystemManaged")
)

// Mapping from string to DataEncryption_Type
var dataEncryption_Type_Values = map[string]DataEncryption_Type{
	"azurekeyvault": DataEncryption_Type_AzureKeyVault,
	"systemmanaged": DataEncryption_Type_SystemManaged,
}

// Enum to indicate whether value is 'Enabled' or 'Disabled'
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type EnableStatusEnum string

const (
	EnableStatusEnum_Disabled = EnableStatusEnum("Disabled")
	EnableStatusEnum_Enabled  = EnableStatusEnum("Enabled")
)

// Mapping from string to EnableStatusEnum
var enableStatusEnum_Values = map[string]EnableStatusEnum{
	"disabled": EnableStatusEnum_Disabled,
	"enabled":  EnableStatusEnum_Enabled,
}

// +kubebuilder:validation:Enum={"Disabled","SameZone","ZoneRedundant"}
type HighAvailability_Mode string

const (
	HighAvailability_Mode_Disabled      = HighAvailability_Mode("Disabled")
	HighAvailability_Mode_SameZone      = HighAvailability_Mode("SameZone")
	HighAvailability_Mode_ZoneRedundant = HighAvailability_Mode("ZoneRedundant")
)

// Mapping from string to HighAvailability_Mode
var highAvailability_Mode_Values = map[string]HighAvailability_Mode{
	"disabled":      HighAvailability_Mode_Disabled,
	"samezone":      HighAvailability_Mode_SameZone,
	"zoneredundant": HighAvailability_Mode_ZoneRedundant,
}

// +kubebuilder:validation:Enum={"AzureBlob"}
type ImportSourceProperties_StorageType string

const ImportSourceProperties_StorageType_AzureBlob = ImportSourceProperties_StorageType("AzureBlob")

// Mapping from string to ImportSourceProperties_StorageType
var importSourceProperties_StorageType_Values = map[string]ImportSourceProperties_StorageType{
	"azureblob": ImportSourceProperties_StorageType_AzureBlob,
}