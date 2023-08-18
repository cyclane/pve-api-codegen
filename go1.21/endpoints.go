package pve

func StrPtr[T ~string](s T) *T {
	return &s
}

type ClusterReplicationRemoveJob string

const (
	ClusterReplicationRemoveJobLocal ClusterReplicationRemoveJob = "local"
	ClusterReplicationRemoveJobFull  ClusterReplicationRemoveJob = "full"
)

type ClusterReplicationType string

const (
	ClusterReplicationTypeLocal ClusterReplicationType = "local"
)

type ClusterMetricsServerInfluxdbproto string

const (
	ClusterMetricsServerInfluxdbprotoUdp   ClusterMetricsServerInfluxdbproto = "udp"
	ClusterMetricsServerInfluxdbprotoHttp  ClusterMetricsServerInfluxdbproto = "http"
	ClusterMetricsServerInfluxdbprotoHttps ClusterMetricsServerInfluxdbproto = "https"
)

type ClusterMetricsServerProto string

const (
	ClusterMetricsServerProtoUdp ClusterMetricsServerProto = "udp"
	ClusterMetricsServerProtoTcp ClusterMetricsServerProto = "tcp"
)

type ClusterMetricsServerType string

const (
	ClusterMetricsServerTypeGraphite ClusterMetricsServerType = "graphite"
	ClusterMetricsServerTypeInfluxdb ClusterMetricsServerType = "influxdb"
)

type FirewallLog string

const (
	FirewallLogEmerg   FirewallLog = "emerg"
	FirewallLogAlert   FirewallLog = "alert"
	FirewallLogCrit    FirewallLog = "crit"
	FirewallLogErr     FirewallLog = "err"
	FirewallLogWarning FirewallLog = "warning"
	FirewallLogNotice  FirewallLog = "notice"
	FirewallLogInfo    FirewallLog = "info"
	FirewallLogDebug   FirewallLog = "debug"
	FirewallLogNolog   FirewallLog = "nolog"
)

type FirewallType string

const (
	FirewallTypeIn    FirewallType = "in"
	FirewallTypeOut   FirewallType = "out"
	FirewallTypeGroup FirewallType = "group"
)

type FirewallOptionsPolicy string

const (
	FirewallOptionsPolicyACCEPT FirewallOptionsPolicy = "ACCEPT"
	FirewallOptionsPolicyREJECT FirewallOptionsPolicy = "REJECT"
	FirewallOptionsPolicyDROP   FirewallOptionsPolicy = "DROP"
)

type FirewallRefsType string

const (
	FirewallRefsTypeAlias FirewallRefsType = "alias"
	FirewallRefsTypeIpset FirewallRefsType = "ipset"
)

type Compress string

const (
	Compress0    Compress = "0"
	Compress1    Compress = "1"
	CompressGzip Compress = "gzip"
	CompressLzo  Compress = "lzo"
	CompressZstd Compress = "zstd"
)

type Mailnotification string

const (
	MailnotificationAlways  Mailnotification = "always"
	MailnotificationFailure Mailnotification = "failure"
)

type Mode string

const (
	ModeSnapshot Mode = "snapshot"
	ModeSuspend  Mode = "suspend"
	ModeStop     Mode = "stop"
)

type ClusterHaResourcesType string

const (
	ClusterHaResourcesTypeCt ClusterHaResourcesType = "ct"
	ClusterHaResourcesTypeVm ClusterHaResourcesType = "vm"
)

type ClusterHaResourceState string

const (
	ClusterHaResourceStateStarted  ClusterHaResourceState = "started"
	ClusterHaResourceStateStopped  ClusterHaResourceState = "stopped"
	ClusterHaResourceStateEnabled  ClusterHaResourceState = "enabled"
	ClusterHaResourceStateDisabled ClusterHaResourceState = "disabled"
	ClusterHaResourceStateIgnored  ClusterHaResourceState = "ignored"
)

type ClusterHaGroupsType string

const (
	ClusterHaGroupsTypeGroup ClusterHaGroupsType = "group"
)

type ClusterAcmePluginsType string

const (
	ClusterAcmePluginsTypeDns        ClusterAcmePluginsType = "dns"
	ClusterAcmePluginsTypeStandalone ClusterAcmePluginsType = "standalone"
)

type ClusterAcmePluginApi string

const (
	ClusterAcmePluginApi1984hosting   ClusterAcmePluginApi = "1984hosting"
	ClusterAcmePluginApiAcmedns       ClusterAcmePluginApi = "acmedns"
	ClusterAcmePluginApiAcmeproxy     ClusterAcmePluginApi = "acmeproxy"
	ClusterAcmePluginApiActive24      ClusterAcmePluginApi = "active24"
	ClusterAcmePluginApiAd            ClusterAcmePluginApi = "ad"
	ClusterAcmePluginApiAli           ClusterAcmePluginApi = "ali"
	ClusterAcmePluginApiAnx           ClusterAcmePluginApi = "anx"
	ClusterAcmePluginApiArvan         ClusterAcmePluginApi = "arvan"
	ClusterAcmePluginApiAurora        ClusterAcmePluginApi = "aurora"
	ClusterAcmePluginApiAutodns       ClusterAcmePluginApi = "autodns"
	ClusterAcmePluginApiAws           ClusterAcmePluginApi = "aws"
	ClusterAcmePluginApiAzion         ClusterAcmePluginApi = "azion"
	ClusterAcmePluginApiAzure         ClusterAcmePluginApi = "azure"
	ClusterAcmePluginApiBunny         ClusterAcmePluginApi = "bunny"
	ClusterAcmePluginApiCf            ClusterAcmePluginApi = "cf"
	ClusterAcmePluginApiClouddns      ClusterAcmePluginApi = "clouddns"
	ClusterAcmePluginApiCloudns       ClusterAcmePluginApi = "cloudns"
	ClusterAcmePluginApiCn            ClusterAcmePluginApi = "cn"
	ClusterAcmePluginApiConoha        ClusterAcmePluginApi = "conoha"
	ClusterAcmePluginApiConstellix    ClusterAcmePluginApi = "constellix"
	ClusterAcmePluginApiCpanel        ClusterAcmePluginApi = "cpanel"
	ClusterAcmePluginApiCuranet       ClusterAcmePluginApi = "curanet"
	ClusterAcmePluginApiCyon          ClusterAcmePluginApi = "cyon"
	ClusterAcmePluginApiDa            ClusterAcmePluginApi = "da"
	ClusterAcmePluginApiDdnss         ClusterAcmePluginApi = "ddnss"
	ClusterAcmePluginApiDesec         ClusterAcmePluginApi = "desec"
	ClusterAcmePluginApiDf            ClusterAcmePluginApi = "df"
	ClusterAcmePluginApiDgon          ClusterAcmePluginApi = "dgon"
	ClusterAcmePluginApiDnshome       ClusterAcmePluginApi = "dnshome"
	ClusterAcmePluginApiDnsimple      ClusterAcmePluginApi = "dnsimple"
	ClusterAcmePluginApiDnsservices   ClusterAcmePluginApi = "dnsservices"
	ClusterAcmePluginApiDo            ClusterAcmePluginApi = "do"
	ClusterAcmePluginApiDoapi         ClusterAcmePluginApi = "doapi"
	ClusterAcmePluginApiDomeneshop    ClusterAcmePluginApi = "domeneshop"
	ClusterAcmePluginApiDp            ClusterAcmePluginApi = "dp"
	ClusterAcmePluginApiDpi           ClusterAcmePluginApi = "dpi"
	ClusterAcmePluginApiDreamhost     ClusterAcmePluginApi = "dreamhost"
	ClusterAcmePluginApiDuckdns       ClusterAcmePluginApi = "duckdns"
	ClusterAcmePluginApiDurabledns    ClusterAcmePluginApi = "durabledns"
	ClusterAcmePluginApiDyn           ClusterAcmePluginApi = "dyn"
	ClusterAcmePluginApiDynu          ClusterAcmePluginApi = "dynu"
	ClusterAcmePluginApiDynv6         ClusterAcmePluginApi = "dynv6"
	ClusterAcmePluginApiEasydns       ClusterAcmePluginApi = "easydns"
	ClusterAcmePluginApiEdgedns       ClusterAcmePluginApi = "edgedns"
	ClusterAcmePluginApiEuserv        ClusterAcmePluginApi = "euserv"
	ClusterAcmePluginApiExoscale      ClusterAcmePluginApi = "exoscale"
	ClusterAcmePluginApiFornex        ClusterAcmePluginApi = "fornex"
	ClusterAcmePluginApiFreedns       ClusterAcmePluginApi = "freedns"
	ClusterAcmePluginApiGandiLivedns  ClusterAcmePluginApi = "gandi_livedns"
	ClusterAcmePluginApiGcloud        ClusterAcmePluginApi = "gcloud"
	ClusterAcmePluginApiGcore         ClusterAcmePluginApi = "gcore"
	ClusterAcmePluginApiGd            ClusterAcmePluginApi = "gd"
	ClusterAcmePluginApiGeoscaling    ClusterAcmePluginApi = "geoscaling"
	ClusterAcmePluginApiGoogledomains ClusterAcmePluginApi = "googledomains"
	ClusterAcmePluginApiHe            ClusterAcmePluginApi = "he"
	ClusterAcmePluginApiHetzner       ClusterAcmePluginApi = "hetzner"
	ClusterAcmePluginApiHexonet       ClusterAcmePluginApi = "hexonet"
	ClusterAcmePluginApiHostingde     ClusterAcmePluginApi = "hostingde"
	ClusterAcmePluginApiHuaweicloud   ClusterAcmePluginApi = "huaweicloud"
	ClusterAcmePluginApiInfoblox      ClusterAcmePluginApi = "infoblox"
	ClusterAcmePluginApiInfomaniak    ClusterAcmePluginApi = "infomaniak"
	ClusterAcmePluginApiInternetbs    ClusterAcmePluginApi = "internetbs"
	ClusterAcmePluginApiInwx          ClusterAcmePluginApi = "inwx"
	ClusterAcmePluginApiIonos         ClusterAcmePluginApi = "ionos"
	ClusterAcmePluginApiIpv64         ClusterAcmePluginApi = "ipv64"
	ClusterAcmePluginApiIspconfig     ClusterAcmePluginApi = "ispconfig"
	ClusterAcmePluginApiJd            ClusterAcmePluginApi = "jd"
	ClusterAcmePluginApiJoker         ClusterAcmePluginApi = "joker"
	ClusterAcmePluginApiKappernet     ClusterAcmePluginApi = "kappernet"
	ClusterAcmePluginApiKas           ClusterAcmePluginApi = "kas"
	ClusterAcmePluginApiKinghost      ClusterAcmePluginApi = "kinghost"
	ClusterAcmePluginApiKnot          ClusterAcmePluginApi = "knot"
	ClusterAcmePluginApiLa            ClusterAcmePluginApi = "la"
	ClusterAcmePluginApiLeaseweb      ClusterAcmePluginApi = "leaseweb"
	ClusterAcmePluginApiLexicon       ClusterAcmePluginApi = "lexicon"
	ClusterAcmePluginApiLinode        ClusterAcmePluginApi = "linode"
	ClusterAcmePluginApiLinodeV4      ClusterAcmePluginApi = "linode_v4"
	ClusterAcmePluginApiLoopia        ClusterAcmePluginApi = "loopia"
	ClusterAcmePluginApiLua           ClusterAcmePluginApi = "lua"
	ClusterAcmePluginApiMaradns       ClusterAcmePluginApi = "maradns"
	ClusterAcmePluginApiMe            ClusterAcmePluginApi = "me"
	ClusterAcmePluginApiMiab          ClusterAcmePluginApi = "miab"
	ClusterAcmePluginApiMisaka        ClusterAcmePluginApi = "misaka"
	ClusterAcmePluginApiMyapi         ClusterAcmePluginApi = "myapi"
	ClusterAcmePluginApiMydevil       ClusterAcmePluginApi = "mydevil"
	ClusterAcmePluginApiMydnsjp       ClusterAcmePluginApi = "mydnsjp"
	ClusterAcmePluginApiMythicBeasts  ClusterAcmePluginApi = "mythic_beasts"
	ClusterAcmePluginApiNamecheap     ClusterAcmePluginApi = "namecheap"
	ClusterAcmePluginApiNamecom       ClusterAcmePluginApi = "namecom"
	ClusterAcmePluginApiNamesilo      ClusterAcmePluginApi = "namesilo"
	ClusterAcmePluginApiNanelo        ClusterAcmePluginApi = "nanelo"
	ClusterAcmePluginApiNederhost     ClusterAcmePluginApi = "nederhost"
	ClusterAcmePluginApiNeodigit      ClusterAcmePluginApi = "neodigit"
	ClusterAcmePluginApiNetcup        ClusterAcmePluginApi = "netcup"
	ClusterAcmePluginApiNetlify       ClusterAcmePluginApi = "netlify"
	ClusterAcmePluginApiNic           ClusterAcmePluginApi = "nic"
	ClusterAcmePluginApiNjalla        ClusterAcmePluginApi = "njalla"
	ClusterAcmePluginApiNm            ClusterAcmePluginApi = "nm"
	ClusterAcmePluginApiNsd           ClusterAcmePluginApi = "nsd"
	ClusterAcmePluginApiNsone         ClusterAcmePluginApi = "nsone"
	ClusterAcmePluginApiNsupdate      ClusterAcmePluginApi = "nsupdate"
	ClusterAcmePluginApiNw            ClusterAcmePluginApi = "nw"
	ClusterAcmePluginApiOci           ClusterAcmePluginApi = "oci"
	ClusterAcmePluginApiOne           ClusterAcmePluginApi = "one"
	ClusterAcmePluginApiOnline        ClusterAcmePluginApi = "online"
	ClusterAcmePluginApiOpenprovider  ClusterAcmePluginApi = "openprovider"
	ClusterAcmePluginApiOpenstack     ClusterAcmePluginApi = "openstack"
	ClusterAcmePluginApiOpnsense      ClusterAcmePluginApi = "opnsense"
	ClusterAcmePluginApiOvh           ClusterAcmePluginApi = "ovh"
	ClusterAcmePluginApiPdns          ClusterAcmePluginApi = "pdns"
	ClusterAcmePluginApiPleskxml      ClusterAcmePluginApi = "pleskxml"
	ClusterAcmePluginApiPointhq       ClusterAcmePluginApi = "pointhq"
	ClusterAcmePluginApiPorkbun       ClusterAcmePluginApi = "porkbun"
	ClusterAcmePluginApiRackcorp      ClusterAcmePluginApi = "rackcorp"
	ClusterAcmePluginApiRackspace     ClusterAcmePluginApi = "rackspace"
	ClusterAcmePluginApiRage4         ClusterAcmePluginApi = "rage4"
	ClusterAcmePluginApiRcode0        ClusterAcmePluginApi = "rcode0"
	ClusterAcmePluginApiRegru         ClusterAcmePluginApi = "regru"
	ClusterAcmePluginApiScaleway      ClusterAcmePluginApi = "scaleway"
	ClusterAcmePluginApiSchlundtech   ClusterAcmePluginApi = "schlundtech"
	ClusterAcmePluginApiSelectel      ClusterAcmePluginApi = "selectel"
	ClusterAcmePluginApiSelfhost      ClusterAcmePluginApi = "selfhost"
	ClusterAcmePluginApiServercow     ClusterAcmePluginApi = "servercow"
	ClusterAcmePluginApiSimply        ClusterAcmePluginApi = "simply"
	ClusterAcmePluginApiTele3         ClusterAcmePluginApi = "tele3"
	ClusterAcmePluginApiTransip       ClusterAcmePluginApi = "transip"
	ClusterAcmePluginApiUdr           ClusterAcmePluginApi = "udr"
	ClusterAcmePluginApiUltra         ClusterAcmePluginApi = "ultra"
	ClusterAcmePluginApiUnoeuro       ClusterAcmePluginApi = "unoeuro"
	ClusterAcmePluginApiVariomedia    ClusterAcmePluginApi = "variomedia"
	ClusterAcmePluginApiVeesp         ClusterAcmePluginApi = "veesp"
	ClusterAcmePluginApiVercel        ClusterAcmePluginApi = "vercel"
	ClusterAcmePluginApiVscale        ClusterAcmePluginApi = "vscale"
	ClusterAcmePluginApiVultr         ClusterAcmePluginApi = "vultr"
	ClusterAcmePluginApiWebsupport    ClusterAcmePluginApi = "websupport"
	ClusterAcmePluginApiWorld4you     ClusterAcmePluginApi = "world4you"
	ClusterAcmePluginApiYandex        ClusterAcmePluginApi = "yandex"
	ClusterAcmePluginApiYc            ClusterAcmePluginApi = "yc"
	ClusterAcmePluginApiZilore        ClusterAcmePluginApi = "zilore"
	ClusterAcmePluginApiZone          ClusterAcmePluginApi = "zone"
	ClusterAcmePluginApiZonomi        ClusterAcmePluginApi = "zonomi"
)

type ClusterCephMetadataScope string

const (
	ClusterCephMetadataScopeAll      ClusterCephMetadataScope = "all"
	ClusterCephMetadataScopeVersions ClusterCephMetadataScope = "versions"
)

type ClusterCephFlagFlag string

const (
	ClusterCephFlagFlagNobackfill  ClusterCephFlagFlag = "nobackfill"
	ClusterCephFlagFlagNodeepScrub ClusterCephFlagFlag = "nodeep-scrub"
	ClusterCephFlagFlagNodown      ClusterCephFlagFlag = "nodown"
	ClusterCephFlagFlagNoin        ClusterCephFlagFlag = "noin"
	ClusterCephFlagFlagNoout       ClusterCephFlagFlag = "noout"
	ClusterCephFlagFlagNorebalance ClusterCephFlagFlag = "norebalance"
	ClusterCephFlagFlagNorecover   ClusterCephFlagFlag = "norecover"
	ClusterCephFlagFlagNoscrub     ClusterCephFlagFlag = "noscrub"
	ClusterCephFlagFlagNotieragent ClusterCephFlagFlag = "notieragent"
	ClusterCephFlagFlagNoup        ClusterCephFlagFlag = "noup"
	ClusterCephFlagFlagPause       ClusterCephFlagFlag = "pause"
)

type SyncScope string

const (
	SyncScopeUsers  SyncScope = "users"
	SyncScopeGroups SyncScope = "groups"
	SyncScopeBoth   SyncScope = "both"
)

type ClusterSdnVnetsType string

const (
	ClusterSdnVnetsTypeVnet ClusterSdnVnetsType = "vnet"
)

type ClusterSdnVnetSubnetsType string

const (
	ClusterSdnVnetSubnetsTypeSubnet ClusterSdnVnetSubnetsType = "subnet"
)

type ClusterSdnZonesType string

const (
	ClusterSdnZonesTypeEvpn   ClusterSdnZonesType = "evpn"
	ClusterSdnZonesTypeFaucet ClusterSdnZonesType = "faucet"
	ClusterSdnZonesTypeQinq   ClusterSdnZonesType = "qinq"
	ClusterSdnZonesTypeSimple ClusterSdnZonesType = "simple"
	ClusterSdnZonesTypeVlan   ClusterSdnZonesType = "vlan"
	ClusterSdnZonesTypeVxlan  ClusterSdnZonesType = "vxlan"
)

type ClusterSdnZoneVlanProtocol string

const (
	ClusterSdnZoneVlanProtocol8021q  ClusterSdnZoneVlanProtocol = "802.1q"
	ClusterSdnZoneVlanProtocol8021ad ClusterSdnZoneVlanProtocol = "802.1ad"
)

type ClusterSdnControllersType string

const (
	ClusterSdnControllersTypeBgp    ClusterSdnControllersType = "bgp"
	ClusterSdnControllersTypeEvpn   ClusterSdnControllersType = "evpn"
	ClusterSdnControllersTypeFaucet ClusterSdnControllersType = "faucet"
)

type ClusterSdnIpamsType string

const (
	ClusterSdnIpamsTypeNetbox  ClusterSdnIpamsType = "netbox"
	ClusterSdnIpamsTypePhpipam ClusterSdnIpamsType = "phpipam"
	ClusterSdnIpamsTypePve     ClusterSdnIpamsType = "pve"
)

type ClusterSdnDnsType string

const (
	ClusterSdnDnsTypePowerdns ClusterSdnDnsType = "powerdns"
)

type ClusterResourcesType string

const (
	ClusterResourcesTypeVm      ClusterResourcesType = "vm"
	ClusterResourcesTypeStorage ClusterResourcesType = "storage"
	ClusterResourcesTypeNode    ClusterResourcesType = "node"
	ClusterResourcesTypeSdn     ClusterResourcesType = "sdn"
)

type ClusterOptionsConsole string

const (
	ClusterOptionsConsoleApplet  ClusterOptionsConsole = "applet"
	ClusterOptionsConsoleVv      ClusterOptionsConsole = "vv"
	ClusterOptionsConsoleHtml5   ClusterOptionsConsole = "html5"
	ClusterOptionsConsoleXtermjs ClusterOptionsConsole = "xtermjs"
)

type ClusterOptionsFencing string

const (
	ClusterOptionsFencingWatchdog ClusterOptionsFencing = "watchdog"
	ClusterOptionsFencingHardware ClusterOptionsFencing = "hardware"
	ClusterOptionsFencingBoth     ClusterOptionsFencing = "both"
)

type Keyboard string

const (
	KeyboardDe   Keyboard = "de"
	KeyboardDeCh Keyboard = "de-ch"
	KeyboardDa   Keyboard = "da"
	KeyboardEnGb Keyboard = "en-gb"
	KeyboardEnUs Keyboard = "en-us"
	KeyboardEs   Keyboard = "es"
	KeyboardFi   Keyboard = "fi"
	KeyboardFr   Keyboard = "fr"
	KeyboardFrBe Keyboard = "fr-be"
	KeyboardFrCa Keyboard = "fr-ca"
	KeyboardFrCh Keyboard = "fr-ch"
	KeyboardHu   Keyboard = "hu"
	KeyboardIs   Keyboard = "is"
	KeyboardIt   Keyboard = "it"
	KeyboardJa   Keyboard = "ja"
	KeyboardLt   Keyboard = "lt"
	KeyboardMk   Keyboard = "mk"
	KeyboardNl   Keyboard = "nl"
	KeyboardNo   Keyboard = "no"
	KeyboardPl   Keyboard = "pl"
	KeyboardPt   Keyboard = "pt"
	KeyboardPtBr Keyboard = "pt-br"
	KeyboardSv   Keyboard = "sv"
	KeyboardSl   Keyboard = "sl"
	KeyboardTr   Keyboard = "tr"
)

type ClusterOptionsLanguage string

const (
	ClusterOptionsLanguageCa   ClusterOptionsLanguage = "ca"
	ClusterOptionsLanguageDa   ClusterOptionsLanguage = "da"
	ClusterOptionsLanguageDe   ClusterOptionsLanguage = "de"
	ClusterOptionsLanguageEn   ClusterOptionsLanguage = "en"
	ClusterOptionsLanguageEs   ClusterOptionsLanguage = "es"
	ClusterOptionsLanguageEu   ClusterOptionsLanguage = "eu"
	ClusterOptionsLanguageFa   ClusterOptionsLanguage = "fa"
	ClusterOptionsLanguageFr   ClusterOptionsLanguage = "fr"
	ClusterOptionsLanguageHe   ClusterOptionsLanguage = "he"
	ClusterOptionsLanguageIt   ClusterOptionsLanguage = "it"
	ClusterOptionsLanguageJa   ClusterOptionsLanguage = "ja"
	ClusterOptionsLanguageNb   ClusterOptionsLanguage = "nb"
	ClusterOptionsLanguageNn   ClusterOptionsLanguage = "nn"
	ClusterOptionsLanguagePl   ClusterOptionsLanguage = "pl"
	ClusterOptionsLanguagePtBR ClusterOptionsLanguage = "pt_BR"
	ClusterOptionsLanguageRu   ClusterOptionsLanguage = "ru"
	ClusterOptionsLanguageSl   ClusterOptionsLanguage = "sl"
	ClusterOptionsLanguageSv   ClusterOptionsLanguage = "sv"
	ClusterOptionsLanguageTr   ClusterOptionsLanguage = "tr"
	ClusterOptionsLanguageZhCN ClusterOptionsLanguage = "zh_CN"
	ClusterOptionsLanguageZhTW ClusterOptionsLanguage = "zh_TW"
)

type NodeQemuArch string

const (
	NodeQemuArchX8664   NodeQemuArch = "x86_64"
	NodeQemuArchAarch64 NodeQemuArch = "aarch64"
)

type NodeQemuBios string

const (
	NodeQemuBiosSeabios NodeQemuBios = "seabios"
	NodeQemuBiosOvmf    NodeQemuBios = "ovmf"
)

type NodeQemuCitype string

const (
	NodeQemuCitypeConfigdrive2 NodeQemuCitype = "configdrive2"
	NodeQemuCitypeNocloud      NodeQemuCitype = "nocloud"
	NodeQemuCitypeOpennebula   NodeQemuCitype = "opennebula"
)

type NodeQemuHugepages string

const (
	NodeQemuHugepagesAny  NodeQemuHugepages = "any"
	NodeQemuHugepages2    NodeQemuHugepages = "2"
	NodeQemuHugepages1024 NodeQemuHugepages = "1024"
)

type NodeQemuLock string

const (
	NodeQemuLockBackup         NodeQemuLock = "backup"
	NodeQemuLockClone          NodeQemuLock = "clone"
	NodeQemuLockCreate         NodeQemuLock = "create"
	NodeQemuLockMigrate        NodeQemuLock = "migrate"
	NodeQemuLockRollback       NodeQemuLock = "rollback"
	NodeQemuLockSnapshot       NodeQemuLock = "snapshot"
	NodeQemuLockSnapshotDelete NodeQemuLock = "snapshot-delete"
	NodeQemuLockSuspending     NodeQemuLock = "suspending"
	NodeQemuLockSuspended      NodeQemuLock = "suspended"
)

type NodeQemuOstype string

const (
	NodeQemuOstypeOther   NodeQemuOstype = "other"
	NodeQemuOstypeWxp     NodeQemuOstype = "wxp"
	NodeQemuOstypeW2k     NodeQemuOstype = "w2k"
	NodeQemuOstypeW2k3    NodeQemuOstype = "w2k3"
	NodeQemuOstypeW2k8    NodeQemuOstype = "w2k8"
	NodeQemuOstypeWvista  NodeQemuOstype = "wvista"
	NodeQemuOstypeWin7    NodeQemuOstype = "win7"
	NodeQemuOstypeWin8    NodeQemuOstype = "win8"
	NodeQemuOstypeWin10   NodeQemuOstype = "win10"
	NodeQemuOstypeWin11   NodeQemuOstype = "win11"
	NodeQemuOstypeL24     NodeQemuOstype = "l24"
	NodeQemuOstypeL26     NodeQemuOstype = "l26"
	NodeQemuOstypeSolaris NodeQemuOstype = "solaris"
)

type NodeQemuScsihw string

const (
	NodeQemuScsihwLsi              NodeQemuScsihw = "lsi"
	NodeQemuScsihwLsi53c810        NodeQemuScsihw = "lsi53c810"
	NodeQemuScsihwVirtioScsiPci    NodeQemuScsihw = "virtio-scsi-pci"
	NodeQemuScsihwVirtioScsiSingle NodeQemuScsihw = "virtio-scsi-single"
	NodeQemuScsihwMegasas          NodeQemuScsihw = "megasas"
	NodeQemuScsihwPvscsi           NodeQemuScsihw = "pvscsi"
)

type NodeQemuAgentCommand string

const (
	NodeQemuAgentCommandFsfreezeFreeze       NodeQemuAgentCommand = "fsfreeze-freeze"
	NodeQemuAgentCommandFsfreezeStatus       NodeQemuAgentCommand = "fsfreeze-status"
	NodeQemuAgentCommandFsfreezeThaw         NodeQemuAgentCommand = "fsfreeze-thaw"
	NodeQemuAgentCommandFstrim               NodeQemuAgentCommand = "fstrim"
	NodeQemuAgentCommandGetFsinfo            NodeQemuAgentCommand = "get-fsinfo"
	NodeQemuAgentCommandGetHostName          NodeQemuAgentCommand = "get-host-name"
	NodeQemuAgentCommandGetMemoryBlockInfo   NodeQemuAgentCommand = "get-memory-block-info"
	NodeQemuAgentCommandGetMemoryBlocks      NodeQemuAgentCommand = "get-memory-blocks"
	NodeQemuAgentCommandGetOsinfo            NodeQemuAgentCommand = "get-osinfo"
	NodeQemuAgentCommandGetTime              NodeQemuAgentCommand = "get-time"
	NodeQemuAgentCommandGetTimezone          NodeQemuAgentCommand = "get-timezone"
	NodeQemuAgentCommandGetUsers             NodeQemuAgentCommand = "get-users"
	NodeQemuAgentCommandGetVcpus             NodeQemuAgentCommand = "get-vcpus"
	NodeQemuAgentCommandInfo                 NodeQemuAgentCommand = "info"
	NodeQemuAgentCommandNetworkGetInterfaces NodeQemuAgentCommand = "network-get-interfaces"
	NodeQemuAgentCommandPing                 NodeQemuAgentCommand = "ping"
	NodeQemuAgentCommandShutdown             NodeQemuAgentCommand = "shutdown"
	NodeQemuAgentCommandSuspendDisk          NodeQemuAgentCommand = "suspend-disk"
	NodeQemuAgentCommandSuspendHybrid        NodeQemuAgentCommand = "suspend-hybrid"
	NodeQemuAgentCommandSuspendRam           NodeQemuAgentCommand = "suspend-ram"
)

type NodeRrdCf string

const (
	NodeRrdCfAVERAGE NodeRrdCf = "AVERAGE"
	NodeRrdCfMAX     NodeRrdCf = "MAX"
)

type NodeRrdTimeframe string

const (
	NodeRrdTimeframeHour  NodeRrdTimeframe = "hour"
	NodeRrdTimeframeDay   NodeRrdTimeframe = "day"
	NodeRrdTimeframeWeek  NodeRrdTimeframe = "week"
	NodeRrdTimeframeMonth NodeRrdTimeframe = "month"
	NodeRrdTimeframeYear  NodeRrdTimeframe = "year"
)

type NodeQemuCloudinitDumpType string

const (
	NodeQemuCloudinitDumpTypeUser    NodeQemuCloudinitDumpType = "user"
	NodeQemuCloudinitDumpTypeNetwork NodeQemuCloudinitDumpType = "network"
	NodeQemuCloudinitDumpTypeMeta    NodeQemuCloudinitDumpType = "meta"
)

type NodeQemuTermproxySerial string

const (
	NodeQemuTermproxySerialSerial0 NodeQemuTermproxySerial = "serial0"
	NodeQemuTermproxySerialSerial1 NodeQemuTermproxySerial = "serial1"
	NodeQemuTermproxySerialSerial2 NodeQemuTermproxySerial = "serial2"
	NodeQemuTermproxySerialSerial3 NodeQemuTermproxySerial = "serial3"
)

type NodeQemuMigrationType string

const (
	NodeQemuMigrationTypeSecure   NodeQemuMigrationType = "secure"
	NodeQemuMigrationTypeInsecure NodeQemuMigrationType = "insecure"
)

type NodeFeatureFeature string

const (
	NodeFeatureFeatureSnapshot NodeFeatureFeature = "snapshot"
	NodeFeatureFeatureClone    NodeFeatureFeature = "clone"
	NodeFeatureFeatureCopy     NodeFeatureFeature = "copy"
)

type NodeQemuFormat string

const (
	NodeQemuFormatRaw   NodeQemuFormat = "raw"
	NodeQemuFormatQcow2 NodeQemuFormat = "qcow2"
	NodeQemuFormatVmdk  NodeQemuFormat = "vmdk"
)

type NodeQemuMoveDiskDisk string

const (
	NodeQemuMoveDiskDiskIde0      NodeQemuMoveDiskDisk = "ide0"
	NodeQemuMoveDiskDiskIde1      NodeQemuMoveDiskDisk = "ide1"
	NodeQemuMoveDiskDiskIde2      NodeQemuMoveDiskDisk = "ide2"
	NodeQemuMoveDiskDiskIde3      NodeQemuMoveDiskDisk = "ide3"
	NodeQemuMoveDiskDiskScsi0     NodeQemuMoveDiskDisk = "scsi0"
	NodeQemuMoveDiskDiskScsi1     NodeQemuMoveDiskDisk = "scsi1"
	NodeQemuMoveDiskDiskScsi2     NodeQemuMoveDiskDisk = "scsi2"
	NodeQemuMoveDiskDiskScsi3     NodeQemuMoveDiskDisk = "scsi3"
	NodeQemuMoveDiskDiskScsi4     NodeQemuMoveDiskDisk = "scsi4"
	NodeQemuMoveDiskDiskScsi5     NodeQemuMoveDiskDisk = "scsi5"
	NodeQemuMoveDiskDiskScsi6     NodeQemuMoveDiskDisk = "scsi6"
	NodeQemuMoveDiskDiskScsi7     NodeQemuMoveDiskDisk = "scsi7"
	NodeQemuMoveDiskDiskScsi8     NodeQemuMoveDiskDisk = "scsi8"
	NodeQemuMoveDiskDiskScsi9     NodeQemuMoveDiskDisk = "scsi9"
	NodeQemuMoveDiskDiskScsi10    NodeQemuMoveDiskDisk = "scsi10"
	NodeQemuMoveDiskDiskScsi11    NodeQemuMoveDiskDisk = "scsi11"
	NodeQemuMoveDiskDiskScsi12    NodeQemuMoveDiskDisk = "scsi12"
	NodeQemuMoveDiskDiskScsi13    NodeQemuMoveDiskDisk = "scsi13"
	NodeQemuMoveDiskDiskScsi14    NodeQemuMoveDiskDisk = "scsi14"
	NodeQemuMoveDiskDiskScsi15    NodeQemuMoveDiskDisk = "scsi15"
	NodeQemuMoveDiskDiskScsi16    NodeQemuMoveDiskDisk = "scsi16"
	NodeQemuMoveDiskDiskScsi17    NodeQemuMoveDiskDisk = "scsi17"
	NodeQemuMoveDiskDiskScsi18    NodeQemuMoveDiskDisk = "scsi18"
	NodeQemuMoveDiskDiskScsi19    NodeQemuMoveDiskDisk = "scsi19"
	NodeQemuMoveDiskDiskScsi20    NodeQemuMoveDiskDisk = "scsi20"
	NodeQemuMoveDiskDiskScsi21    NodeQemuMoveDiskDisk = "scsi21"
	NodeQemuMoveDiskDiskScsi22    NodeQemuMoveDiskDisk = "scsi22"
	NodeQemuMoveDiskDiskScsi23    NodeQemuMoveDiskDisk = "scsi23"
	NodeQemuMoveDiskDiskScsi24    NodeQemuMoveDiskDisk = "scsi24"
	NodeQemuMoveDiskDiskScsi25    NodeQemuMoveDiskDisk = "scsi25"
	NodeQemuMoveDiskDiskScsi26    NodeQemuMoveDiskDisk = "scsi26"
	NodeQemuMoveDiskDiskScsi27    NodeQemuMoveDiskDisk = "scsi27"
	NodeQemuMoveDiskDiskScsi28    NodeQemuMoveDiskDisk = "scsi28"
	NodeQemuMoveDiskDiskScsi29    NodeQemuMoveDiskDisk = "scsi29"
	NodeQemuMoveDiskDiskScsi30    NodeQemuMoveDiskDisk = "scsi30"
	NodeQemuMoveDiskDiskVirtio0   NodeQemuMoveDiskDisk = "virtio0"
	NodeQemuMoveDiskDiskVirtio1   NodeQemuMoveDiskDisk = "virtio1"
	NodeQemuMoveDiskDiskVirtio2   NodeQemuMoveDiskDisk = "virtio2"
	NodeQemuMoveDiskDiskVirtio3   NodeQemuMoveDiskDisk = "virtio3"
	NodeQemuMoveDiskDiskVirtio4   NodeQemuMoveDiskDisk = "virtio4"
	NodeQemuMoveDiskDiskVirtio5   NodeQemuMoveDiskDisk = "virtio5"
	NodeQemuMoveDiskDiskVirtio6   NodeQemuMoveDiskDisk = "virtio6"
	NodeQemuMoveDiskDiskVirtio7   NodeQemuMoveDiskDisk = "virtio7"
	NodeQemuMoveDiskDiskVirtio8   NodeQemuMoveDiskDisk = "virtio8"
	NodeQemuMoveDiskDiskVirtio9   NodeQemuMoveDiskDisk = "virtio9"
	NodeQemuMoveDiskDiskVirtio10  NodeQemuMoveDiskDisk = "virtio10"
	NodeQemuMoveDiskDiskVirtio11  NodeQemuMoveDiskDisk = "virtio11"
	NodeQemuMoveDiskDiskVirtio12  NodeQemuMoveDiskDisk = "virtio12"
	NodeQemuMoveDiskDiskVirtio13  NodeQemuMoveDiskDisk = "virtio13"
	NodeQemuMoveDiskDiskVirtio14  NodeQemuMoveDiskDisk = "virtio14"
	NodeQemuMoveDiskDiskVirtio15  NodeQemuMoveDiskDisk = "virtio15"
	NodeQemuMoveDiskDiskSata0     NodeQemuMoveDiskDisk = "sata0"
	NodeQemuMoveDiskDiskSata1     NodeQemuMoveDiskDisk = "sata1"
	NodeQemuMoveDiskDiskSata2     NodeQemuMoveDiskDisk = "sata2"
	NodeQemuMoveDiskDiskSata3     NodeQemuMoveDiskDisk = "sata3"
	NodeQemuMoveDiskDiskSata4     NodeQemuMoveDiskDisk = "sata4"
	NodeQemuMoveDiskDiskSata5     NodeQemuMoveDiskDisk = "sata5"
	NodeQemuMoveDiskDiskEfidisk0  NodeQemuMoveDiskDisk = "efidisk0"
	NodeQemuMoveDiskDiskTpmstate0 NodeQemuMoveDiskDisk = "tpmstate0"
	NodeQemuMoveDiskDiskUnused0   NodeQemuMoveDiskDisk = "unused0"
	NodeQemuMoveDiskDiskUnused1   NodeQemuMoveDiskDisk = "unused1"
	NodeQemuMoveDiskDiskUnused2   NodeQemuMoveDiskDisk = "unused2"
	NodeQemuMoveDiskDiskUnused3   NodeQemuMoveDiskDisk = "unused3"
	NodeQemuMoveDiskDiskUnused4   NodeQemuMoveDiskDisk = "unused4"
	NodeQemuMoveDiskDiskUnused5   NodeQemuMoveDiskDisk = "unused5"
	NodeQemuMoveDiskDiskUnused6   NodeQemuMoveDiskDisk = "unused6"
	NodeQemuMoveDiskDiskUnused7   NodeQemuMoveDiskDisk = "unused7"
	NodeQemuMoveDiskDiskUnused8   NodeQemuMoveDiskDisk = "unused8"
	NodeQemuMoveDiskDiskUnused9   NodeQemuMoveDiskDisk = "unused9"
	NodeQemuMoveDiskDiskUnused10  NodeQemuMoveDiskDisk = "unused10"
	NodeQemuMoveDiskDiskUnused11  NodeQemuMoveDiskDisk = "unused11"
	NodeQemuMoveDiskDiskUnused12  NodeQemuMoveDiskDisk = "unused12"
	NodeQemuMoveDiskDiskUnused13  NodeQemuMoveDiskDisk = "unused13"
	NodeQemuMoveDiskDiskUnused14  NodeQemuMoveDiskDisk = "unused14"
	NodeQemuMoveDiskDiskUnused15  NodeQemuMoveDiskDisk = "unused15"
	NodeQemuMoveDiskDiskUnused16  NodeQemuMoveDiskDisk = "unused16"
	NodeQemuMoveDiskDiskUnused17  NodeQemuMoveDiskDisk = "unused17"
	NodeQemuMoveDiskDiskUnused18  NodeQemuMoveDiskDisk = "unused18"
	NodeQemuMoveDiskDiskUnused19  NodeQemuMoveDiskDisk = "unused19"
	NodeQemuMoveDiskDiskUnused20  NodeQemuMoveDiskDisk = "unused20"
	NodeQemuMoveDiskDiskUnused21  NodeQemuMoveDiskDisk = "unused21"
	NodeQemuMoveDiskDiskUnused22  NodeQemuMoveDiskDisk = "unused22"
	NodeQemuMoveDiskDiskUnused23  NodeQemuMoveDiskDisk = "unused23"
	NodeQemuMoveDiskDiskUnused24  NodeQemuMoveDiskDisk = "unused24"
	NodeQemuMoveDiskDiskUnused25  NodeQemuMoveDiskDisk = "unused25"
	NodeQemuMoveDiskDiskUnused26  NodeQemuMoveDiskDisk = "unused26"
	NodeQemuMoveDiskDiskUnused27  NodeQemuMoveDiskDisk = "unused27"
	NodeQemuMoveDiskDiskUnused28  NodeQemuMoveDiskDisk = "unused28"
	NodeQemuMoveDiskDiskUnused29  NodeQemuMoveDiskDisk = "unused29"
	NodeQemuMoveDiskDiskUnused30  NodeQemuMoveDiskDisk = "unused30"
	NodeQemuMoveDiskDiskUnused31  NodeQemuMoveDiskDisk = "unused31"
	NodeQemuMoveDiskDiskUnused32  NodeQemuMoveDiskDisk = "unused32"
	NodeQemuMoveDiskDiskUnused33  NodeQemuMoveDiskDisk = "unused33"
	NodeQemuMoveDiskDiskUnused34  NodeQemuMoveDiskDisk = "unused34"
	NodeQemuMoveDiskDiskUnused35  NodeQemuMoveDiskDisk = "unused35"
	NodeQemuMoveDiskDiskUnused36  NodeQemuMoveDiskDisk = "unused36"
	NodeQemuMoveDiskDiskUnused37  NodeQemuMoveDiskDisk = "unused37"
	NodeQemuMoveDiskDiskUnused38  NodeQemuMoveDiskDisk = "unused38"
	NodeQemuMoveDiskDiskUnused39  NodeQemuMoveDiskDisk = "unused39"
	NodeQemuMoveDiskDiskUnused40  NodeQemuMoveDiskDisk = "unused40"
	NodeQemuMoveDiskDiskUnused41  NodeQemuMoveDiskDisk = "unused41"
	NodeQemuMoveDiskDiskUnused42  NodeQemuMoveDiskDisk = "unused42"
	NodeQemuMoveDiskDiskUnused43  NodeQemuMoveDiskDisk = "unused43"
	NodeQemuMoveDiskDiskUnused44  NodeQemuMoveDiskDisk = "unused44"
	NodeQemuMoveDiskDiskUnused45  NodeQemuMoveDiskDisk = "unused45"
	NodeQemuMoveDiskDiskUnused46  NodeQemuMoveDiskDisk = "unused46"
	NodeQemuMoveDiskDiskUnused47  NodeQemuMoveDiskDisk = "unused47"
	NodeQemuMoveDiskDiskUnused48  NodeQemuMoveDiskDisk = "unused48"
	NodeQemuMoveDiskDiskUnused49  NodeQemuMoveDiskDisk = "unused49"
	NodeQemuMoveDiskDiskUnused50  NodeQemuMoveDiskDisk = "unused50"
	NodeQemuMoveDiskDiskUnused51  NodeQemuMoveDiskDisk = "unused51"
	NodeQemuMoveDiskDiskUnused52  NodeQemuMoveDiskDisk = "unused52"
	NodeQemuMoveDiskDiskUnused53  NodeQemuMoveDiskDisk = "unused53"
	NodeQemuMoveDiskDiskUnused54  NodeQemuMoveDiskDisk = "unused54"
	NodeQemuMoveDiskDiskUnused55  NodeQemuMoveDiskDisk = "unused55"
	NodeQemuMoveDiskDiskUnused56  NodeQemuMoveDiskDisk = "unused56"
	NodeQemuMoveDiskDiskUnused57  NodeQemuMoveDiskDisk = "unused57"
	NodeQemuMoveDiskDiskUnused58  NodeQemuMoveDiskDisk = "unused58"
	NodeQemuMoveDiskDiskUnused59  NodeQemuMoveDiskDisk = "unused59"
	NodeQemuMoveDiskDiskUnused60  NodeQemuMoveDiskDisk = "unused60"
	NodeQemuMoveDiskDiskUnused61  NodeQemuMoveDiskDisk = "unused61"
	NodeQemuMoveDiskDiskUnused62  NodeQemuMoveDiskDisk = "unused62"
	NodeQemuMoveDiskDiskUnused63  NodeQemuMoveDiskDisk = "unused63"
	NodeQemuMoveDiskDiskUnused64  NodeQemuMoveDiskDisk = "unused64"
	NodeQemuMoveDiskDiskUnused65  NodeQemuMoveDiskDisk = "unused65"
	NodeQemuMoveDiskDiskUnused66  NodeQemuMoveDiskDisk = "unused66"
	NodeQemuMoveDiskDiskUnused67  NodeQemuMoveDiskDisk = "unused67"
	NodeQemuMoveDiskDiskUnused68  NodeQemuMoveDiskDisk = "unused68"
	NodeQemuMoveDiskDiskUnused69  NodeQemuMoveDiskDisk = "unused69"
	NodeQemuMoveDiskDiskUnused70  NodeQemuMoveDiskDisk = "unused70"
	NodeQemuMoveDiskDiskUnused71  NodeQemuMoveDiskDisk = "unused71"
	NodeQemuMoveDiskDiskUnused72  NodeQemuMoveDiskDisk = "unused72"
	NodeQemuMoveDiskDiskUnused73  NodeQemuMoveDiskDisk = "unused73"
	NodeQemuMoveDiskDiskUnused74  NodeQemuMoveDiskDisk = "unused74"
	NodeQemuMoveDiskDiskUnused75  NodeQemuMoveDiskDisk = "unused75"
	NodeQemuMoveDiskDiskUnused76  NodeQemuMoveDiskDisk = "unused76"
	NodeQemuMoveDiskDiskUnused77  NodeQemuMoveDiskDisk = "unused77"
	NodeQemuMoveDiskDiskUnused78  NodeQemuMoveDiskDisk = "unused78"
	NodeQemuMoveDiskDiskUnused79  NodeQemuMoveDiskDisk = "unused79"
	NodeQemuMoveDiskDiskUnused80  NodeQemuMoveDiskDisk = "unused80"
	NodeQemuMoveDiskDiskUnused81  NodeQemuMoveDiskDisk = "unused81"
	NodeQemuMoveDiskDiskUnused82  NodeQemuMoveDiskDisk = "unused82"
	NodeQemuMoveDiskDiskUnused83  NodeQemuMoveDiskDisk = "unused83"
	NodeQemuMoveDiskDiskUnused84  NodeQemuMoveDiskDisk = "unused84"
	NodeQemuMoveDiskDiskUnused85  NodeQemuMoveDiskDisk = "unused85"
	NodeQemuMoveDiskDiskUnused86  NodeQemuMoveDiskDisk = "unused86"
	NodeQemuMoveDiskDiskUnused87  NodeQemuMoveDiskDisk = "unused87"
	NodeQemuMoveDiskDiskUnused88  NodeQemuMoveDiskDisk = "unused88"
	NodeQemuMoveDiskDiskUnused89  NodeQemuMoveDiskDisk = "unused89"
	NodeQemuMoveDiskDiskUnused90  NodeQemuMoveDiskDisk = "unused90"
	NodeQemuMoveDiskDiskUnused91  NodeQemuMoveDiskDisk = "unused91"
	NodeQemuMoveDiskDiskUnused92  NodeQemuMoveDiskDisk = "unused92"
	NodeQemuMoveDiskDiskUnused93  NodeQemuMoveDiskDisk = "unused93"
	NodeQemuMoveDiskDiskUnused94  NodeQemuMoveDiskDisk = "unused94"
	NodeQemuMoveDiskDiskUnused95  NodeQemuMoveDiskDisk = "unused95"
	NodeQemuMoveDiskDiskUnused96  NodeQemuMoveDiskDisk = "unused96"
	NodeQemuMoveDiskDiskUnused97  NodeQemuMoveDiskDisk = "unused97"
	NodeQemuMoveDiskDiskUnused98  NodeQemuMoveDiskDisk = "unused98"
	NodeQemuMoveDiskDiskUnused99  NodeQemuMoveDiskDisk = "unused99"
	NodeQemuMoveDiskDiskUnused100 NodeQemuMoveDiskDisk = "unused100"
	NodeQemuMoveDiskDiskUnused101 NodeQemuMoveDiskDisk = "unused101"
	NodeQemuMoveDiskDiskUnused102 NodeQemuMoveDiskDisk = "unused102"
	NodeQemuMoveDiskDiskUnused103 NodeQemuMoveDiskDisk = "unused103"
	NodeQemuMoveDiskDiskUnused104 NodeQemuMoveDiskDisk = "unused104"
	NodeQemuMoveDiskDiskUnused105 NodeQemuMoveDiskDisk = "unused105"
	NodeQemuMoveDiskDiskUnused106 NodeQemuMoveDiskDisk = "unused106"
	NodeQemuMoveDiskDiskUnused107 NodeQemuMoveDiskDisk = "unused107"
	NodeQemuMoveDiskDiskUnused108 NodeQemuMoveDiskDisk = "unused108"
	NodeQemuMoveDiskDiskUnused109 NodeQemuMoveDiskDisk = "unused109"
	NodeQemuMoveDiskDiskUnused110 NodeQemuMoveDiskDisk = "unused110"
	NodeQemuMoveDiskDiskUnused111 NodeQemuMoveDiskDisk = "unused111"
	NodeQemuMoveDiskDiskUnused112 NodeQemuMoveDiskDisk = "unused112"
	NodeQemuMoveDiskDiskUnused113 NodeQemuMoveDiskDisk = "unused113"
	NodeQemuMoveDiskDiskUnused114 NodeQemuMoveDiskDisk = "unused114"
	NodeQemuMoveDiskDiskUnused115 NodeQemuMoveDiskDisk = "unused115"
	NodeQemuMoveDiskDiskUnused116 NodeQemuMoveDiskDisk = "unused116"
	NodeQemuMoveDiskDiskUnused117 NodeQemuMoveDiskDisk = "unused117"
	NodeQemuMoveDiskDiskUnused118 NodeQemuMoveDiskDisk = "unused118"
	NodeQemuMoveDiskDiskUnused119 NodeQemuMoveDiskDisk = "unused119"
	NodeQemuMoveDiskDiskUnused120 NodeQemuMoveDiskDisk = "unused120"
	NodeQemuMoveDiskDiskUnused121 NodeQemuMoveDiskDisk = "unused121"
	NodeQemuMoveDiskDiskUnused122 NodeQemuMoveDiskDisk = "unused122"
	NodeQemuMoveDiskDiskUnused123 NodeQemuMoveDiskDisk = "unused123"
	NodeQemuMoveDiskDiskUnused124 NodeQemuMoveDiskDisk = "unused124"
	NodeQemuMoveDiskDiskUnused125 NodeQemuMoveDiskDisk = "unused125"
	NodeQemuMoveDiskDiskUnused126 NodeQemuMoveDiskDisk = "unused126"
	NodeQemuMoveDiskDiskUnused127 NodeQemuMoveDiskDisk = "unused127"
	NodeQemuMoveDiskDiskUnused128 NodeQemuMoveDiskDisk = "unused128"
	NodeQemuMoveDiskDiskUnused129 NodeQemuMoveDiskDisk = "unused129"
	NodeQemuMoveDiskDiskUnused130 NodeQemuMoveDiskDisk = "unused130"
	NodeQemuMoveDiskDiskUnused131 NodeQemuMoveDiskDisk = "unused131"
	NodeQemuMoveDiskDiskUnused132 NodeQemuMoveDiskDisk = "unused132"
	NodeQemuMoveDiskDiskUnused133 NodeQemuMoveDiskDisk = "unused133"
	NodeQemuMoveDiskDiskUnused134 NodeQemuMoveDiskDisk = "unused134"
	NodeQemuMoveDiskDiskUnused135 NodeQemuMoveDiskDisk = "unused135"
	NodeQemuMoveDiskDiskUnused136 NodeQemuMoveDiskDisk = "unused136"
	NodeQemuMoveDiskDiskUnused137 NodeQemuMoveDiskDisk = "unused137"
	NodeQemuMoveDiskDiskUnused138 NodeQemuMoveDiskDisk = "unused138"
	NodeQemuMoveDiskDiskUnused139 NodeQemuMoveDiskDisk = "unused139"
	NodeQemuMoveDiskDiskUnused140 NodeQemuMoveDiskDisk = "unused140"
	NodeQemuMoveDiskDiskUnused141 NodeQemuMoveDiskDisk = "unused141"
	NodeQemuMoveDiskDiskUnused142 NodeQemuMoveDiskDisk = "unused142"
	NodeQemuMoveDiskDiskUnused143 NodeQemuMoveDiskDisk = "unused143"
	NodeQemuMoveDiskDiskUnused144 NodeQemuMoveDiskDisk = "unused144"
	NodeQemuMoveDiskDiskUnused145 NodeQemuMoveDiskDisk = "unused145"
	NodeQemuMoveDiskDiskUnused146 NodeQemuMoveDiskDisk = "unused146"
	NodeQemuMoveDiskDiskUnused147 NodeQemuMoveDiskDisk = "unused147"
	NodeQemuMoveDiskDiskUnused148 NodeQemuMoveDiskDisk = "unused148"
	NodeQemuMoveDiskDiskUnused149 NodeQemuMoveDiskDisk = "unused149"
	NodeQemuMoveDiskDiskUnused150 NodeQemuMoveDiskDisk = "unused150"
	NodeQemuMoveDiskDiskUnused151 NodeQemuMoveDiskDisk = "unused151"
	NodeQemuMoveDiskDiskUnused152 NodeQemuMoveDiskDisk = "unused152"
	NodeQemuMoveDiskDiskUnused153 NodeQemuMoveDiskDisk = "unused153"
	NodeQemuMoveDiskDiskUnused154 NodeQemuMoveDiskDisk = "unused154"
	NodeQemuMoveDiskDiskUnused155 NodeQemuMoveDiskDisk = "unused155"
	NodeQemuMoveDiskDiskUnused156 NodeQemuMoveDiskDisk = "unused156"
	NodeQemuMoveDiskDiskUnused157 NodeQemuMoveDiskDisk = "unused157"
	NodeQemuMoveDiskDiskUnused158 NodeQemuMoveDiskDisk = "unused158"
	NodeQemuMoveDiskDiskUnused159 NodeQemuMoveDiskDisk = "unused159"
	NodeQemuMoveDiskDiskUnused160 NodeQemuMoveDiskDisk = "unused160"
	NodeQemuMoveDiskDiskUnused161 NodeQemuMoveDiskDisk = "unused161"
	NodeQemuMoveDiskDiskUnused162 NodeQemuMoveDiskDisk = "unused162"
	NodeQemuMoveDiskDiskUnused163 NodeQemuMoveDiskDisk = "unused163"
	NodeQemuMoveDiskDiskUnused164 NodeQemuMoveDiskDisk = "unused164"
	NodeQemuMoveDiskDiskUnused165 NodeQemuMoveDiskDisk = "unused165"
	NodeQemuMoveDiskDiskUnused166 NodeQemuMoveDiskDisk = "unused166"
	NodeQemuMoveDiskDiskUnused167 NodeQemuMoveDiskDisk = "unused167"
	NodeQemuMoveDiskDiskUnused168 NodeQemuMoveDiskDisk = "unused168"
	NodeQemuMoveDiskDiskUnused169 NodeQemuMoveDiskDisk = "unused169"
	NodeQemuMoveDiskDiskUnused170 NodeQemuMoveDiskDisk = "unused170"
	NodeQemuMoveDiskDiskUnused171 NodeQemuMoveDiskDisk = "unused171"
	NodeQemuMoveDiskDiskUnused172 NodeQemuMoveDiskDisk = "unused172"
	NodeQemuMoveDiskDiskUnused173 NodeQemuMoveDiskDisk = "unused173"
	NodeQemuMoveDiskDiskUnused174 NodeQemuMoveDiskDisk = "unused174"
	NodeQemuMoveDiskDiskUnused175 NodeQemuMoveDiskDisk = "unused175"
	NodeQemuMoveDiskDiskUnused176 NodeQemuMoveDiskDisk = "unused176"
	NodeQemuMoveDiskDiskUnused177 NodeQemuMoveDiskDisk = "unused177"
	NodeQemuMoveDiskDiskUnused178 NodeQemuMoveDiskDisk = "unused178"
	NodeQemuMoveDiskDiskUnused179 NodeQemuMoveDiskDisk = "unused179"
	NodeQemuMoveDiskDiskUnused180 NodeQemuMoveDiskDisk = "unused180"
	NodeQemuMoveDiskDiskUnused181 NodeQemuMoveDiskDisk = "unused181"
	NodeQemuMoveDiskDiskUnused182 NodeQemuMoveDiskDisk = "unused182"
	NodeQemuMoveDiskDiskUnused183 NodeQemuMoveDiskDisk = "unused183"
	NodeQemuMoveDiskDiskUnused184 NodeQemuMoveDiskDisk = "unused184"
	NodeQemuMoveDiskDiskUnused185 NodeQemuMoveDiskDisk = "unused185"
	NodeQemuMoveDiskDiskUnused186 NodeQemuMoveDiskDisk = "unused186"
	NodeQemuMoveDiskDiskUnused187 NodeQemuMoveDiskDisk = "unused187"
	NodeQemuMoveDiskDiskUnused188 NodeQemuMoveDiskDisk = "unused188"
	NodeQemuMoveDiskDiskUnused189 NodeQemuMoveDiskDisk = "unused189"
	NodeQemuMoveDiskDiskUnused190 NodeQemuMoveDiskDisk = "unused190"
	NodeQemuMoveDiskDiskUnused191 NodeQemuMoveDiskDisk = "unused191"
	NodeQemuMoveDiskDiskUnused192 NodeQemuMoveDiskDisk = "unused192"
	NodeQemuMoveDiskDiskUnused193 NodeQemuMoveDiskDisk = "unused193"
	NodeQemuMoveDiskDiskUnused194 NodeQemuMoveDiskDisk = "unused194"
	NodeQemuMoveDiskDiskUnused195 NodeQemuMoveDiskDisk = "unused195"
	NodeQemuMoveDiskDiskUnused196 NodeQemuMoveDiskDisk = "unused196"
	NodeQemuMoveDiskDiskUnused197 NodeQemuMoveDiskDisk = "unused197"
	NodeQemuMoveDiskDiskUnused198 NodeQemuMoveDiskDisk = "unused198"
	NodeQemuMoveDiskDiskUnused199 NodeQemuMoveDiskDisk = "unused199"
	NodeQemuMoveDiskDiskUnused200 NodeQemuMoveDiskDisk = "unused200"
	NodeQemuMoveDiskDiskUnused201 NodeQemuMoveDiskDisk = "unused201"
	NodeQemuMoveDiskDiskUnused202 NodeQemuMoveDiskDisk = "unused202"
	NodeQemuMoveDiskDiskUnused203 NodeQemuMoveDiskDisk = "unused203"
	NodeQemuMoveDiskDiskUnused204 NodeQemuMoveDiskDisk = "unused204"
	NodeQemuMoveDiskDiskUnused205 NodeQemuMoveDiskDisk = "unused205"
	NodeQemuMoveDiskDiskUnused206 NodeQemuMoveDiskDisk = "unused206"
	NodeQemuMoveDiskDiskUnused207 NodeQemuMoveDiskDisk = "unused207"
	NodeQemuMoveDiskDiskUnused208 NodeQemuMoveDiskDisk = "unused208"
	NodeQemuMoveDiskDiskUnused209 NodeQemuMoveDiskDisk = "unused209"
	NodeQemuMoveDiskDiskUnused210 NodeQemuMoveDiskDisk = "unused210"
	NodeQemuMoveDiskDiskUnused211 NodeQemuMoveDiskDisk = "unused211"
	NodeQemuMoveDiskDiskUnused212 NodeQemuMoveDiskDisk = "unused212"
	NodeQemuMoveDiskDiskUnused213 NodeQemuMoveDiskDisk = "unused213"
	NodeQemuMoveDiskDiskUnused214 NodeQemuMoveDiskDisk = "unused214"
	NodeQemuMoveDiskDiskUnused215 NodeQemuMoveDiskDisk = "unused215"
	NodeQemuMoveDiskDiskUnused216 NodeQemuMoveDiskDisk = "unused216"
	NodeQemuMoveDiskDiskUnused217 NodeQemuMoveDiskDisk = "unused217"
	NodeQemuMoveDiskDiskUnused218 NodeQemuMoveDiskDisk = "unused218"
	NodeQemuMoveDiskDiskUnused219 NodeQemuMoveDiskDisk = "unused219"
	NodeQemuMoveDiskDiskUnused220 NodeQemuMoveDiskDisk = "unused220"
	NodeQemuMoveDiskDiskUnused221 NodeQemuMoveDiskDisk = "unused221"
	NodeQemuMoveDiskDiskUnused222 NodeQemuMoveDiskDisk = "unused222"
	NodeQemuMoveDiskDiskUnused223 NodeQemuMoveDiskDisk = "unused223"
	NodeQemuMoveDiskDiskUnused224 NodeQemuMoveDiskDisk = "unused224"
	NodeQemuMoveDiskDiskUnused225 NodeQemuMoveDiskDisk = "unused225"
	NodeQemuMoveDiskDiskUnused226 NodeQemuMoveDiskDisk = "unused226"
	NodeQemuMoveDiskDiskUnused227 NodeQemuMoveDiskDisk = "unused227"
	NodeQemuMoveDiskDiskUnused228 NodeQemuMoveDiskDisk = "unused228"
	NodeQemuMoveDiskDiskUnused229 NodeQemuMoveDiskDisk = "unused229"
	NodeQemuMoveDiskDiskUnused230 NodeQemuMoveDiskDisk = "unused230"
	NodeQemuMoveDiskDiskUnused231 NodeQemuMoveDiskDisk = "unused231"
	NodeQemuMoveDiskDiskUnused232 NodeQemuMoveDiskDisk = "unused232"
	NodeQemuMoveDiskDiskUnused233 NodeQemuMoveDiskDisk = "unused233"
	NodeQemuMoveDiskDiskUnused234 NodeQemuMoveDiskDisk = "unused234"
	NodeQemuMoveDiskDiskUnused235 NodeQemuMoveDiskDisk = "unused235"
	NodeQemuMoveDiskDiskUnused236 NodeQemuMoveDiskDisk = "unused236"
	NodeQemuMoveDiskDiskUnused237 NodeQemuMoveDiskDisk = "unused237"
	NodeQemuMoveDiskDiskUnused238 NodeQemuMoveDiskDisk = "unused238"
	NodeQemuMoveDiskDiskUnused239 NodeQemuMoveDiskDisk = "unused239"
	NodeQemuMoveDiskDiskUnused240 NodeQemuMoveDiskDisk = "unused240"
	NodeQemuMoveDiskDiskUnused241 NodeQemuMoveDiskDisk = "unused241"
	NodeQemuMoveDiskDiskUnused242 NodeQemuMoveDiskDisk = "unused242"
	NodeQemuMoveDiskDiskUnused243 NodeQemuMoveDiskDisk = "unused243"
	NodeQemuMoveDiskDiskUnused244 NodeQemuMoveDiskDisk = "unused244"
	NodeQemuMoveDiskDiskUnused245 NodeQemuMoveDiskDisk = "unused245"
	NodeQemuMoveDiskDiskUnused246 NodeQemuMoveDiskDisk = "unused246"
	NodeQemuMoveDiskDiskUnused247 NodeQemuMoveDiskDisk = "unused247"
	NodeQemuMoveDiskDiskUnused248 NodeQemuMoveDiskDisk = "unused248"
	NodeQemuMoveDiskDiskUnused249 NodeQemuMoveDiskDisk = "unused249"
	NodeQemuMoveDiskDiskUnused250 NodeQemuMoveDiskDisk = "unused250"
	NodeQemuMoveDiskDiskUnused251 NodeQemuMoveDiskDisk = "unused251"
	NodeQemuMoveDiskDiskUnused252 NodeQemuMoveDiskDisk = "unused252"
	NodeQemuMoveDiskDiskUnused253 NodeQemuMoveDiskDisk = "unused253"
	NodeQemuMoveDiskDiskUnused254 NodeQemuMoveDiskDisk = "unused254"
	NodeQemuMoveDiskDiskUnused255 NodeQemuMoveDiskDisk = "unused255"
)

type NodeQemuDisk string

const (
	NodeQemuDiskIde0      NodeQemuDisk = "ide0"
	NodeQemuDiskIde1      NodeQemuDisk = "ide1"
	NodeQemuDiskIde2      NodeQemuDisk = "ide2"
	NodeQemuDiskIde3      NodeQemuDisk = "ide3"
	NodeQemuDiskScsi0     NodeQemuDisk = "scsi0"
	NodeQemuDiskScsi1     NodeQemuDisk = "scsi1"
	NodeQemuDiskScsi2     NodeQemuDisk = "scsi2"
	NodeQemuDiskScsi3     NodeQemuDisk = "scsi3"
	NodeQemuDiskScsi4     NodeQemuDisk = "scsi4"
	NodeQemuDiskScsi5     NodeQemuDisk = "scsi5"
	NodeQemuDiskScsi6     NodeQemuDisk = "scsi6"
	NodeQemuDiskScsi7     NodeQemuDisk = "scsi7"
	NodeQemuDiskScsi8     NodeQemuDisk = "scsi8"
	NodeQemuDiskScsi9     NodeQemuDisk = "scsi9"
	NodeQemuDiskScsi10    NodeQemuDisk = "scsi10"
	NodeQemuDiskScsi11    NodeQemuDisk = "scsi11"
	NodeQemuDiskScsi12    NodeQemuDisk = "scsi12"
	NodeQemuDiskScsi13    NodeQemuDisk = "scsi13"
	NodeQemuDiskScsi14    NodeQemuDisk = "scsi14"
	NodeQemuDiskScsi15    NodeQemuDisk = "scsi15"
	NodeQemuDiskScsi16    NodeQemuDisk = "scsi16"
	NodeQemuDiskScsi17    NodeQemuDisk = "scsi17"
	NodeQemuDiskScsi18    NodeQemuDisk = "scsi18"
	NodeQemuDiskScsi19    NodeQemuDisk = "scsi19"
	NodeQemuDiskScsi20    NodeQemuDisk = "scsi20"
	NodeQemuDiskScsi21    NodeQemuDisk = "scsi21"
	NodeQemuDiskScsi22    NodeQemuDisk = "scsi22"
	NodeQemuDiskScsi23    NodeQemuDisk = "scsi23"
	NodeQemuDiskScsi24    NodeQemuDisk = "scsi24"
	NodeQemuDiskScsi25    NodeQemuDisk = "scsi25"
	NodeQemuDiskScsi26    NodeQemuDisk = "scsi26"
	NodeQemuDiskScsi27    NodeQemuDisk = "scsi27"
	NodeQemuDiskScsi28    NodeQemuDisk = "scsi28"
	NodeQemuDiskScsi29    NodeQemuDisk = "scsi29"
	NodeQemuDiskScsi30    NodeQemuDisk = "scsi30"
	NodeQemuDiskVirtio0   NodeQemuDisk = "virtio0"
	NodeQemuDiskVirtio1   NodeQemuDisk = "virtio1"
	NodeQemuDiskVirtio2   NodeQemuDisk = "virtio2"
	NodeQemuDiskVirtio3   NodeQemuDisk = "virtio3"
	NodeQemuDiskVirtio4   NodeQemuDisk = "virtio4"
	NodeQemuDiskVirtio5   NodeQemuDisk = "virtio5"
	NodeQemuDiskVirtio6   NodeQemuDisk = "virtio6"
	NodeQemuDiskVirtio7   NodeQemuDisk = "virtio7"
	NodeQemuDiskVirtio8   NodeQemuDisk = "virtio8"
	NodeQemuDiskVirtio9   NodeQemuDisk = "virtio9"
	NodeQemuDiskVirtio10  NodeQemuDisk = "virtio10"
	NodeQemuDiskVirtio11  NodeQemuDisk = "virtio11"
	NodeQemuDiskVirtio12  NodeQemuDisk = "virtio12"
	NodeQemuDiskVirtio13  NodeQemuDisk = "virtio13"
	NodeQemuDiskVirtio14  NodeQemuDisk = "virtio14"
	NodeQemuDiskVirtio15  NodeQemuDisk = "virtio15"
	NodeQemuDiskSata0     NodeQemuDisk = "sata0"
	NodeQemuDiskSata1     NodeQemuDisk = "sata1"
	NodeQemuDiskSata2     NodeQemuDisk = "sata2"
	NodeQemuDiskSata3     NodeQemuDisk = "sata3"
	NodeQemuDiskSata4     NodeQemuDisk = "sata4"
	NodeQemuDiskSata5     NodeQemuDisk = "sata5"
	NodeQemuDiskEfidisk0  NodeQemuDisk = "efidisk0"
	NodeQemuDiskTpmstate0 NodeQemuDisk = "tpmstate0"
)

type NodeLxcArch string

const (
	NodeLxcArchAmd64   NodeLxcArch = "amd64"
	NodeLxcArchI386    NodeLxcArch = "i386"
	NodeLxcArchArm64   NodeLxcArch = "arm64"
	NodeLxcArchArmhf   NodeLxcArch = "armhf"
	NodeLxcArchRiscv32 NodeLxcArch = "riscv32"
	NodeLxcArchRiscv64 NodeLxcArch = "riscv64"
)

type NodeLxcCmode string

const (
	NodeLxcCmodeShell   NodeLxcCmode = "shell"
	NodeLxcCmodeConsole NodeLxcCmode = "console"
	NodeLxcCmodeTty     NodeLxcCmode = "tty"
)

type NodeLxcLock string

const (
	NodeLxcLockBackup         NodeLxcLock = "backup"
	NodeLxcLockCreate         NodeLxcLock = "create"
	NodeLxcLockDestroyed      NodeLxcLock = "destroyed"
	NodeLxcLockDisk           NodeLxcLock = "disk"
	NodeLxcLockFstrim         NodeLxcLock = "fstrim"
	NodeLxcLockMigrate        NodeLxcLock = "migrate"
	NodeLxcLockMounted        NodeLxcLock = "mounted"
	NodeLxcLockRollback       NodeLxcLock = "rollback"
	NodeLxcLockSnapshot       NodeLxcLock = "snapshot"
	NodeLxcLockSnapshotDelete NodeLxcLock = "snapshot-delete"
)

type NodeLxcOstype string

const (
	NodeLxcOstypeDebian    NodeLxcOstype = "debian"
	NodeLxcOstypeDevuan    NodeLxcOstype = "devuan"
	NodeLxcOstypeUbuntu    NodeLxcOstype = "ubuntu"
	NodeLxcOstypeCentos    NodeLxcOstype = "centos"
	NodeLxcOstypeFedora    NodeLxcOstype = "fedora"
	NodeLxcOstypeOpensuse  NodeLxcOstype = "opensuse"
	NodeLxcOstypeArchlinux NodeLxcOstype = "archlinux"
	NodeLxcOstypeAlpine    NodeLxcOstype = "alpine"
	NodeLxcOstypeGentoo    NodeLxcOstype = "gentoo"
	NodeLxcOstypeNixos     NodeLxcOstype = "nixos"
	NodeLxcOstypeUnmanaged NodeLxcOstype = "unmanaged"
)

type NodeLxcResizeDisk string

const (
	NodeLxcResizeDiskRootfs NodeLxcResizeDisk = "rootfs"
	NodeLxcResizeDiskMp0    NodeLxcResizeDisk = "mp0"
	NodeLxcResizeDiskMp1    NodeLxcResizeDisk = "mp1"
	NodeLxcResizeDiskMp2    NodeLxcResizeDisk = "mp2"
	NodeLxcResizeDiskMp3    NodeLxcResizeDisk = "mp3"
	NodeLxcResizeDiskMp4    NodeLxcResizeDisk = "mp4"
	NodeLxcResizeDiskMp5    NodeLxcResizeDisk = "mp5"
	NodeLxcResizeDiskMp6    NodeLxcResizeDisk = "mp6"
	NodeLxcResizeDiskMp7    NodeLxcResizeDisk = "mp7"
	NodeLxcResizeDiskMp8    NodeLxcResizeDisk = "mp8"
	NodeLxcResizeDiskMp9    NodeLxcResizeDisk = "mp9"
	NodeLxcResizeDiskMp10   NodeLxcResizeDisk = "mp10"
	NodeLxcResizeDiskMp11   NodeLxcResizeDisk = "mp11"
	NodeLxcResizeDiskMp12   NodeLxcResizeDisk = "mp12"
	NodeLxcResizeDiskMp13   NodeLxcResizeDisk = "mp13"
	NodeLxcResizeDiskMp14   NodeLxcResizeDisk = "mp14"
	NodeLxcResizeDiskMp15   NodeLxcResizeDisk = "mp15"
	NodeLxcResizeDiskMp16   NodeLxcResizeDisk = "mp16"
	NodeLxcResizeDiskMp17   NodeLxcResizeDisk = "mp17"
	NodeLxcResizeDiskMp18   NodeLxcResizeDisk = "mp18"
	NodeLxcResizeDiskMp19   NodeLxcResizeDisk = "mp19"
	NodeLxcResizeDiskMp20   NodeLxcResizeDisk = "mp20"
	NodeLxcResizeDiskMp21   NodeLxcResizeDisk = "mp21"
	NodeLxcResizeDiskMp22   NodeLxcResizeDisk = "mp22"
	NodeLxcResizeDiskMp23   NodeLxcResizeDisk = "mp23"
	NodeLxcResizeDiskMp24   NodeLxcResizeDisk = "mp24"
	NodeLxcResizeDiskMp25   NodeLxcResizeDisk = "mp25"
	NodeLxcResizeDiskMp26   NodeLxcResizeDisk = "mp26"
	NodeLxcResizeDiskMp27   NodeLxcResizeDisk = "mp27"
	NodeLxcResizeDiskMp28   NodeLxcResizeDisk = "mp28"
	NodeLxcResizeDiskMp29   NodeLxcResizeDisk = "mp29"
	NodeLxcResizeDiskMp30   NodeLxcResizeDisk = "mp30"
	NodeLxcResizeDiskMp31   NodeLxcResizeDisk = "mp31"
	NodeLxcResizeDiskMp32   NodeLxcResizeDisk = "mp32"
	NodeLxcResizeDiskMp33   NodeLxcResizeDisk = "mp33"
	NodeLxcResizeDiskMp34   NodeLxcResizeDisk = "mp34"
	NodeLxcResizeDiskMp35   NodeLxcResizeDisk = "mp35"
	NodeLxcResizeDiskMp36   NodeLxcResizeDisk = "mp36"
	NodeLxcResizeDiskMp37   NodeLxcResizeDisk = "mp37"
	NodeLxcResizeDiskMp38   NodeLxcResizeDisk = "mp38"
	NodeLxcResizeDiskMp39   NodeLxcResizeDisk = "mp39"
	NodeLxcResizeDiskMp40   NodeLxcResizeDisk = "mp40"
	NodeLxcResizeDiskMp41   NodeLxcResizeDisk = "mp41"
	NodeLxcResizeDiskMp42   NodeLxcResizeDisk = "mp42"
	NodeLxcResizeDiskMp43   NodeLxcResizeDisk = "mp43"
	NodeLxcResizeDiskMp44   NodeLxcResizeDisk = "mp44"
	NodeLxcResizeDiskMp45   NodeLxcResizeDisk = "mp45"
	NodeLxcResizeDiskMp46   NodeLxcResizeDisk = "mp46"
	NodeLxcResizeDiskMp47   NodeLxcResizeDisk = "mp47"
	NodeLxcResizeDiskMp48   NodeLxcResizeDisk = "mp48"
	NodeLxcResizeDiskMp49   NodeLxcResizeDisk = "mp49"
	NodeLxcResizeDiskMp50   NodeLxcResizeDisk = "mp50"
	NodeLxcResizeDiskMp51   NodeLxcResizeDisk = "mp51"
	NodeLxcResizeDiskMp52   NodeLxcResizeDisk = "mp52"
	NodeLxcResizeDiskMp53   NodeLxcResizeDisk = "mp53"
	NodeLxcResizeDiskMp54   NodeLxcResizeDisk = "mp54"
	NodeLxcResizeDiskMp55   NodeLxcResizeDisk = "mp55"
	NodeLxcResizeDiskMp56   NodeLxcResizeDisk = "mp56"
	NodeLxcResizeDiskMp57   NodeLxcResizeDisk = "mp57"
	NodeLxcResizeDiskMp58   NodeLxcResizeDisk = "mp58"
	NodeLxcResizeDiskMp59   NodeLxcResizeDisk = "mp59"
	NodeLxcResizeDiskMp60   NodeLxcResizeDisk = "mp60"
	NodeLxcResizeDiskMp61   NodeLxcResizeDisk = "mp61"
	NodeLxcResizeDiskMp62   NodeLxcResizeDisk = "mp62"
	NodeLxcResizeDiskMp63   NodeLxcResizeDisk = "mp63"
	NodeLxcResizeDiskMp64   NodeLxcResizeDisk = "mp64"
	NodeLxcResizeDiskMp65   NodeLxcResizeDisk = "mp65"
	NodeLxcResizeDiskMp66   NodeLxcResizeDisk = "mp66"
	NodeLxcResizeDiskMp67   NodeLxcResizeDisk = "mp67"
	NodeLxcResizeDiskMp68   NodeLxcResizeDisk = "mp68"
	NodeLxcResizeDiskMp69   NodeLxcResizeDisk = "mp69"
	NodeLxcResizeDiskMp70   NodeLxcResizeDisk = "mp70"
	NodeLxcResizeDiskMp71   NodeLxcResizeDisk = "mp71"
	NodeLxcResizeDiskMp72   NodeLxcResizeDisk = "mp72"
	NodeLxcResizeDiskMp73   NodeLxcResizeDisk = "mp73"
	NodeLxcResizeDiskMp74   NodeLxcResizeDisk = "mp74"
	NodeLxcResizeDiskMp75   NodeLxcResizeDisk = "mp75"
	NodeLxcResizeDiskMp76   NodeLxcResizeDisk = "mp76"
	NodeLxcResizeDiskMp77   NodeLxcResizeDisk = "mp77"
	NodeLxcResizeDiskMp78   NodeLxcResizeDisk = "mp78"
	NodeLxcResizeDiskMp79   NodeLxcResizeDisk = "mp79"
	NodeLxcResizeDiskMp80   NodeLxcResizeDisk = "mp80"
	NodeLxcResizeDiskMp81   NodeLxcResizeDisk = "mp81"
	NodeLxcResizeDiskMp82   NodeLxcResizeDisk = "mp82"
	NodeLxcResizeDiskMp83   NodeLxcResizeDisk = "mp83"
	NodeLxcResizeDiskMp84   NodeLxcResizeDisk = "mp84"
	NodeLxcResizeDiskMp85   NodeLxcResizeDisk = "mp85"
	NodeLxcResizeDiskMp86   NodeLxcResizeDisk = "mp86"
	NodeLxcResizeDiskMp87   NodeLxcResizeDisk = "mp87"
	NodeLxcResizeDiskMp88   NodeLxcResizeDisk = "mp88"
	NodeLxcResizeDiskMp89   NodeLxcResizeDisk = "mp89"
	NodeLxcResizeDiskMp90   NodeLxcResizeDisk = "mp90"
	NodeLxcResizeDiskMp91   NodeLxcResizeDisk = "mp91"
	NodeLxcResizeDiskMp92   NodeLxcResizeDisk = "mp92"
	NodeLxcResizeDiskMp93   NodeLxcResizeDisk = "mp93"
	NodeLxcResizeDiskMp94   NodeLxcResizeDisk = "mp94"
	NodeLxcResizeDiskMp95   NodeLxcResizeDisk = "mp95"
	NodeLxcResizeDiskMp96   NodeLxcResizeDisk = "mp96"
	NodeLxcResizeDiskMp97   NodeLxcResizeDisk = "mp97"
	NodeLxcResizeDiskMp98   NodeLxcResizeDisk = "mp98"
	NodeLxcResizeDiskMp99   NodeLxcResizeDisk = "mp99"
	NodeLxcResizeDiskMp100  NodeLxcResizeDisk = "mp100"
	NodeLxcResizeDiskMp101  NodeLxcResizeDisk = "mp101"
	NodeLxcResizeDiskMp102  NodeLxcResizeDisk = "mp102"
	NodeLxcResizeDiskMp103  NodeLxcResizeDisk = "mp103"
	NodeLxcResizeDiskMp104  NodeLxcResizeDisk = "mp104"
	NodeLxcResizeDiskMp105  NodeLxcResizeDisk = "mp105"
	NodeLxcResizeDiskMp106  NodeLxcResizeDisk = "mp106"
	NodeLxcResizeDiskMp107  NodeLxcResizeDisk = "mp107"
	NodeLxcResizeDiskMp108  NodeLxcResizeDisk = "mp108"
	NodeLxcResizeDiskMp109  NodeLxcResizeDisk = "mp109"
	NodeLxcResizeDiskMp110  NodeLxcResizeDisk = "mp110"
	NodeLxcResizeDiskMp111  NodeLxcResizeDisk = "mp111"
	NodeLxcResizeDiskMp112  NodeLxcResizeDisk = "mp112"
	NodeLxcResizeDiskMp113  NodeLxcResizeDisk = "mp113"
	NodeLxcResizeDiskMp114  NodeLxcResizeDisk = "mp114"
	NodeLxcResizeDiskMp115  NodeLxcResizeDisk = "mp115"
	NodeLxcResizeDiskMp116  NodeLxcResizeDisk = "mp116"
	NodeLxcResizeDiskMp117  NodeLxcResizeDisk = "mp117"
	NodeLxcResizeDiskMp118  NodeLxcResizeDisk = "mp118"
	NodeLxcResizeDiskMp119  NodeLxcResizeDisk = "mp119"
	NodeLxcResizeDiskMp120  NodeLxcResizeDisk = "mp120"
	NodeLxcResizeDiskMp121  NodeLxcResizeDisk = "mp121"
	NodeLxcResizeDiskMp122  NodeLxcResizeDisk = "mp122"
	NodeLxcResizeDiskMp123  NodeLxcResizeDisk = "mp123"
	NodeLxcResizeDiskMp124  NodeLxcResizeDisk = "mp124"
	NodeLxcResizeDiskMp125  NodeLxcResizeDisk = "mp125"
	NodeLxcResizeDiskMp126  NodeLxcResizeDisk = "mp126"
	NodeLxcResizeDiskMp127  NodeLxcResizeDisk = "mp127"
	NodeLxcResizeDiskMp128  NodeLxcResizeDisk = "mp128"
	NodeLxcResizeDiskMp129  NodeLxcResizeDisk = "mp129"
	NodeLxcResizeDiskMp130  NodeLxcResizeDisk = "mp130"
	NodeLxcResizeDiskMp131  NodeLxcResizeDisk = "mp131"
	NodeLxcResizeDiskMp132  NodeLxcResizeDisk = "mp132"
	NodeLxcResizeDiskMp133  NodeLxcResizeDisk = "mp133"
	NodeLxcResizeDiskMp134  NodeLxcResizeDisk = "mp134"
	NodeLxcResizeDiskMp135  NodeLxcResizeDisk = "mp135"
	NodeLxcResizeDiskMp136  NodeLxcResizeDisk = "mp136"
	NodeLxcResizeDiskMp137  NodeLxcResizeDisk = "mp137"
	NodeLxcResizeDiskMp138  NodeLxcResizeDisk = "mp138"
	NodeLxcResizeDiskMp139  NodeLxcResizeDisk = "mp139"
	NodeLxcResizeDiskMp140  NodeLxcResizeDisk = "mp140"
	NodeLxcResizeDiskMp141  NodeLxcResizeDisk = "mp141"
	NodeLxcResizeDiskMp142  NodeLxcResizeDisk = "mp142"
	NodeLxcResizeDiskMp143  NodeLxcResizeDisk = "mp143"
	NodeLxcResizeDiskMp144  NodeLxcResizeDisk = "mp144"
	NodeLxcResizeDiskMp145  NodeLxcResizeDisk = "mp145"
	NodeLxcResizeDiskMp146  NodeLxcResizeDisk = "mp146"
	NodeLxcResizeDiskMp147  NodeLxcResizeDisk = "mp147"
	NodeLxcResizeDiskMp148  NodeLxcResizeDisk = "mp148"
	NodeLxcResizeDiskMp149  NodeLxcResizeDisk = "mp149"
	NodeLxcResizeDiskMp150  NodeLxcResizeDisk = "mp150"
	NodeLxcResizeDiskMp151  NodeLxcResizeDisk = "mp151"
	NodeLxcResizeDiskMp152  NodeLxcResizeDisk = "mp152"
	NodeLxcResizeDiskMp153  NodeLxcResizeDisk = "mp153"
	NodeLxcResizeDiskMp154  NodeLxcResizeDisk = "mp154"
	NodeLxcResizeDiskMp155  NodeLxcResizeDisk = "mp155"
	NodeLxcResizeDiskMp156  NodeLxcResizeDisk = "mp156"
	NodeLxcResizeDiskMp157  NodeLxcResizeDisk = "mp157"
	NodeLxcResizeDiskMp158  NodeLxcResizeDisk = "mp158"
	NodeLxcResizeDiskMp159  NodeLxcResizeDisk = "mp159"
	NodeLxcResizeDiskMp160  NodeLxcResizeDisk = "mp160"
	NodeLxcResizeDiskMp161  NodeLxcResizeDisk = "mp161"
	NodeLxcResizeDiskMp162  NodeLxcResizeDisk = "mp162"
	NodeLxcResizeDiskMp163  NodeLxcResizeDisk = "mp163"
	NodeLxcResizeDiskMp164  NodeLxcResizeDisk = "mp164"
	NodeLxcResizeDiskMp165  NodeLxcResizeDisk = "mp165"
	NodeLxcResizeDiskMp166  NodeLxcResizeDisk = "mp166"
	NodeLxcResizeDiskMp167  NodeLxcResizeDisk = "mp167"
	NodeLxcResizeDiskMp168  NodeLxcResizeDisk = "mp168"
	NodeLxcResizeDiskMp169  NodeLxcResizeDisk = "mp169"
	NodeLxcResizeDiskMp170  NodeLxcResizeDisk = "mp170"
	NodeLxcResizeDiskMp171  NodeLxcResizeDisk = "mp171"
	NodeLxcResizeDiskMp172  NodeLxcResizeDisk = "mp172"
	NodeLxcResizeDiskMp173  NodeLxcResizeDisk = "mp173"
	NodeLxcResizeDiskMp174  NodeLxcResizeDisk = "mp174"
	NodeLxcResizeDiskMp175  NodeLxcResizeDisk = "mp175"
	NodeLxcResizeDiskMp176  NodeLxcResizeDisk = "mp176"
	NodeLxcResizeDiskMp177  NodeLxcResizeDisk = "mp177"
	NodeLxcResizeDiskMp178  NodeLxcResizeDisk = "mp178"
	NodeLxcResizeDiskMp179  NodeLxcResizeDisk = "mp179"
	NodeLxcResizeDiskMp180  NodeLxcResizeDisk = "mp180"
	NodeLxcResizeDiskMp181  NodeLxcResizeDisk = "mp181"
	NodeLxcResizeDiskMp182  NodeLxcResizeDisk = "mp182"
	NodeLxcResizeDiskMp183  NodeLxcResizeDisk = "mp183"
	NodeLxcResizeDiskMp184  NodeLxcResizeDisk = "mp184"
	NodeLxcResizeDiskMp185  NodeLxcResizeDisk = "mp185"
	NodeLxcResizeDiskMp186  NodeLxcResizeDisk = "mp186"
	NodeLxcResizeDiskMp187  NodeLxcResizeDisk = "mp187"
	NodeLxcResizeDiskMp188  NodeLxcResizeDisk = "mp188"
	NodeLxcResizeDiskMp189  NodeLxcResizeDisk = "mp189"
	NodeLxcResizeDiskMp190  NodeLxcResizeDisk = "mp190"
	NodeLxcResizeDiskMp191  NodeLxcResizeDisk = "mp191"
	NodeLxcResizeDiskMp192  NodeLxcResizeDisk = "mp192"
	NodeLxcResizeDiskMp193  NodeLxcResizeDisk = "mp193"
	NodeLxcResizeDiskMp194  NodeLxcResizeDisk = "mp194"
	NodeLxcResizeDiskMp195  NodeLxcResizeDisk = "mp195"
	NodeLxcResizeDiskMp196  NodeLxcResizeDisk = "mp196"
	NodeLxcResizeDiskMp197  NodeLxcResizeDisk = "mp197"
	NodeLxcResizeDiskMp198  NodeLxcResizeDisk = "mp198"
	NodeLxcResizeDiskMp199  NodeLxcResizeDisk = "mp199"
	NodeLxcResizeDiskMp200  NodeLxcResizeDisk = "mp200"
	NodeLxcResizeDiskMp201  NodeLxcResizeDisk = "mp201"
	NodeLxcResizeDiskMp202  NodeLxcResizeDisk = "mp202"
	NodeLxcResizeDiskMp203  NodeLxcResizeDisk = "mp203"
	NodeLxcResizeDiskMp204  NodeLxcResizeDisk = "mp204"
	NodeLxcResizeDiskMp205  NodeLxcResizeDisk = "mp205"
	NodeLxcResizeDiskMp206  NodeLxcResizeDisk = "mp206"
	NodeLxcResizeDiskMp207  NodeLxcResizeDisk = "mp207"
	NodeLxcResizeDiskMp208  NodeLxcResizeDisk = "mp208"
	NodeLxcResizeDiskMp209  NodeLxcResizeDisk = "mp209"
	NodeLxcResizeDiskMp210  NodeLxcResizeDisk = "mp210"
	NodeLxcResizeDiskMp211  NodeLxcResizeDisk = "mp211"
	NodeLxcResizeDiskMp212  NodeLxcResizeDisk = "mp212"
	NodeLxcResizeDiskMp213  NodeLxcResizeDisk = "mp213"
	NodeLxcResizeDiskMp214  NodeLxcResizeDisk = "mp214"
	NodeLxcResizeDiskMp215  NodeLxcResizeDisk = "mp215"
	NodeLxcResizeDiskMp216  NodeLxcResizeDisk = "mp216"
	NodeLxcResizeDiskMp217  NodeLxcResizeDisk = "mp217"
	NodeLxcResizeDiskMp218  NodeLxcResizeDisk = "mp218"
	NodeLxcResizeDiskMp219  NodeLxcResizeDisk = "mp219"
	NodeLxcResizeDiskMp220  NodeLxcResizeDisk = "mp220"
	NodeLxcResizeDiskMp221  NodeLxcResizeDisk = "mp221"
	NodeLxcResizeDiskMp222  NodeLxcResizeDisk = "mp222"
	NodeLxcResizeDiskMp223  NodeLxcResizeDisk = "mp223"
	NodeLxcResizeDiskMp224  NodeLxcResizeDisk = "mp224"
	NodeLxcResizeDiskMp225  NodeLxcResizeDisk = "mp225"
	NodeLxcResizeDiskMp226  NodeLxcResizeDisk = "mp226"
	NodeLxcResizeDiskMp227  NodeLxcResizeDisk = "mp227"
	NodeLxcResizeDiskMp228  NodeLxcResizeDisk = "mp228"
	NodeLxcResizeDiskMp229  NodeLxcResizeDisk = "mp229"
	NodeLxcResizeDiskMp230  NodeLxcResizeDisk = "mp230"
	NodeLxcResizeDiskMp231  NodeLxcResizeDisk = "mp231"
	NodeLxcResizeDiskMp232  NodeLxcResizeDisk = "mp232"
	NodeLxcResizeDiskMp233  NodeLxcResizeDisk = "mp233"
	NodeLxcResizeDiskMp234  NodeLxcResizeDisk = "mp234"
	NodeLxcResizeDiskMp235  NodeLxcResizeDisk = "mp235"
	NodeLxcResizeDiskMp236  NodeLxcResizeDisk = "mp236"
	NodeLxcResizeDiskMp237  NodeLxcResizeDisk = "mp237"
	NodeLxcResizeDiskMp238  NodeLxcResizeDisk = "mp238"
	NodeLxcResizeDiskMp239  NodeLxcResizeDisk = "mp239"
	NodeLxcResizeDiskMp240  NodeLxcResizeDisk = "mp240"
	NodeLxcResizeDiskMp241  NodeLxcResizeDisk = "mp241"
	NodeLxcResizeDiskMp242  NodeLxcResizeDisk = "mp242"
	NodeLxcResizeDiskMp243  NodeLxcResizeDisk = "mp243"
	NodeLxcResizeDiskMp244  NodeLxcResizeDisk = "mp244"
	NodeLxcResizeDiskMp245  NodeLxcResizeDisk = "mp245"
	NodeLxcResizeDiskMp246  NodeLxcResizeDisk = "mp246"
	NodeLxcResizeDiskMp247  NodeLxcResizeDisk = "mp247"
	NodeLxcResizeDiskMp248  NodeLxcResizeDisk = "mp248"
	NodeLxcResizeDiskMp249  NodeLxcResizeDisk = "mp249"
	NodeLxcResizeDiskMp250  NodeLxcResizeDisk = "mp250"
	NodeLxcResizeDiskMp251  NodeLxcResizeDisk = "mp251"
	NodeLxcResizeDiskMp252  NodeLxcResizeDisk = "mp252"
	NodeLxcResizeDiskMp253  NodeLxcResizeDisk = "mp253"
	NodeLxcResizeDiskMp254  NodeLxcResizeDisk = "mp254"
	NodeLxcResizeDiskMp255  NodeLxcResizeDisk = "mp255"
)

type NodeLxcMoveVolumeVolume string

const (
	NodeLxcMoveVolumeVolumeRootfs    NodeLxcMoveVolumeVolume = "rootfs"
	NodeLxcMoveVolumeVolumeMp0       NodeLxcMoveVolumeVolume = "mp0"
	NodeLxcMoveVolumeVolumeMp1       NodeLxcMoveVolumeVolume = "mp1"
	NodeLxcMoveVolumeVolumeMp2       NodeLxcMoveVolumeVolume = "mp2"
	NodeLxcMoveVolumeVolumeMp3       NodeLxcMoveVolumeVolume = "mp3"
	NodeLxcMoveVolumeVolumeMp4       NodeLxcMoveVolumeVolume = "mp4"
	NodeLxcMoveVolumeVolumeMp5       NodeLxcMoveVolumeVolume = "mp5"
	NodeLxcMoveVolumeVolumeMp6       NodeLxcMoveVolumeVolume = "mp6"
	NodeLxcMoveVolumeVolumeMp7       NodeLxcMoveVolumeVolume = "mp7"
	NodeLxcMoveVolumeVolumeMp8       NodeLxcMoveVolumeVolume = "mp8"
	NodeLxcMoveVolumeVolumeMp9       NodeLxcMoveVolumeVolume = "mp9"
	NodeLxcMoveVolumeVolumeMp10      NodeLxcMoveVolumeVolume = "mp10"
	NodeLxcMoveVolumeVolumeMp11      NodeLxcMoveVolumeVolume = "mp11"
	NodeLxcMoveVolumeVolumeMp12      NodeLxcMoveVolumeVolume = "mp12"
	NodeLxcMoveVolumeVolumeMp13      NodeLxcMoveVolumeVolume = "mp13"
	NodeLxcMoveVolumeVolumeMp14      NodeLxcMoveVolumeVolume = "mp14"
	NodeLxcMoveVolumeVolumeMp15      NodeLxcMoveVolumeVolume = "mp15"
	NodeLxcMoveVolumeVolumeMp16      NodeLxcMoveVolumeVolume = "mp16"
	NodeLxcMoveVolumeVolumeMp17      NodeLxcMoveVolumeVolume = "mp17"
	NodeLxcMoveVolumeVolumeMp18      NodeLxcMoveVolumeVolume = "mp18"
	NodeLxcMoveVolumeVolumeMp19      NodeLxcMoveVolumeVolume = "mp19"
	NodeLxcMoveVolumeVolumeMp20      NodeLxcMoveVolumeVolume = "mp20"
	NodeLxcMoveVolumeVolumeMp21      NodeLxcMoveVolumeVolume = "mp21"
	NodeLxcMoveVolumeVolumeMp22      NodeLxcMoveVolumeVolume = "mp22"
	NodeLxcMoveVolumeVolumeMp23      NodeLxcMoveVolumeVolume = "mp23"
	NodeLxcMoveVolumeVolumeMp24      NodeLxcMoveVolumeVolume = "mp24"
	NodeLxcMoveVolumeVolumeMp25      NodeLxcMoveVolumeVolume = "mp25"
	NodeLxcMoveVolumeVolumeMp26      NodeLxcMoveVolumeVolume = "mp26"
	NodeLxcMoveVolumeVolumeMp27      NodeLxcMoveVolumeVolume = "mp27"
	NodeLxcMoveVolumeVolumeMp28      NodeLxcMoveVolumeVolume = "mp28"
	NodeLxcMoveVolumeVolumeMp29      NodeLxcMoveVolumeVolume = "mp29"
	NodeLxcMoveVolumeVolumeMp30      NodeLxcMoveVolumeVolume = "mp30"
	NodeLxcMoveVolumeVolumeMp31      NodeLxcMoveVolumeVolume = "mp31"
	NodeLxcMoveVolumeVolumeMp32      NodeLxcMoveVolumeVolume = "mp32"
	NodeLxcMoveVolumeVolumeMp33      NodeLxcMoveVolumeVolume = "mp33"
	NodeLxcMoveVolumeVolumeMp34      NodeLxcMoveVolumeVolume = "mp34"
	NodeLxcMoveVolumeVolumeMp35      NodeLxcMoveVolumeVolume = "mp35"
	NodeLxcMoveVolumeVolumeMp36      NodeLxcMoveVolumeVolume = "mp36"
	NodeLxcMoveVolumeVolumeMp37      NodeLxcMoveVolumeVolume = "mp37"
	NodeLxcMoveVolumeVolumeMp38      NodeLxcMoveVolumeVolume = "mp38"
	NodeLxcMoveVolumeVolumeMp39      NodeLxcMoveVolumeVolume = "mp39"
	NodeLxcMoveVolumeVolumeMp40      NodeLxcMoveVolumeVolume = "mp40"
	NodeLxcMoveVolumeVolumeMp41      NodeLxcMoveVolumeVolume = "mp41"
	NodeLxcMoveVolumeVolumeMp42      NodeLxcMoveVolumeVolume = "mp42"
	NodeLxcMoveVolumeVolumeMp43      NodeLxcMoveVolumeVolume = "mp43"
	NodeLxcMoveVolumeVolumeMp44      NodeLxcMoveVolumeVolume = "mp44"
	NodeLxcMoveVolumeVolumeMp45      NodeLxcMoveVolumeVolume = "mp45"
	NodeLxcMoveVolumeVolumeMp46      NodeLxcMoveVolumeVolume = "mp46"
	NodeLxcMoveVolumeVolumeMp47      NodeLxcMoveVolumeVolume = "mp47"
	NodeLxcMoveVolumeVolumeMp48      NodeLxcMoveVolumeVolume = "mp48"
	NodeLxcMoveVolumeVolumeMp49      NodeLxcMoveVolumeVolume = "mp49"
	NodeLxcMoveVolumeVolumeMp50      NodeLxcMoveVolumeVolume = "mp50"
	NodeLxcMoveVolumeVolumeMp51      NodeLxcMoveVolumeVolume = "mp51"
	NodeLxcMoveVolumeVolumeMp52      NodeLxcMoveVolumeVolume = "mp52"
	NodeLxcMoveVolumeVolumeMp53      NodeLxcMoveVolumeVolume = "mp53"
	NodeLxcMoveVolumeVolumeMp54      NodeLxcMoveVolumeVolume = "mp54"
	NodeLxcMoveVolumeVolumeMp55      NodeLxcMoveVolumeVolume = "mp55"
	NodeLxcMoveVolumeVolumeMp56      NodeLxcMoveVolumeVolume = "mp56"
	NodeLxcMoveVolumeVolumeMp57      NodeLxcMoveVolumeVolume = "mp57"
	NodeLxcMoveVolumeVolumeMp58      NodeLxcMoveVolumeVolume = "mp58"
	NodeLxcMoveVolumeVolumeMp59      NodeLxcMoveVolumeVolume = "mp59"
	NodeLxcMoveVolumeVolumeMp60      NodeLxcMoveVolumeVolume = "mp60"
	NodeLxcMoveVolumeVolumeMp61      NodeLxcMoveVolumeVolume = "mp61"
	NodeLxcMoveVolumeVolumeMp62      NodeLxcMoveVolumeVolume = "mp62"
	NodeLxcMoveVolumeVolumeMp63      NodeLxcMoveVolumeVolume = "mp63"
	NodeLxcMoveVolumeVolumeMp64      NodeLxcMoveVolumeVolume = "mp64"
	NodeLxcMoveVolumeVolumeMp65      NodeLxcMoveVolumeVolume = "mp65"
	NodeLxcMoveVolumeVolumeMp66      NodeLxcMoveVolumeVolume = "mp66"
	NodeLxcMoveVolumeVolumeMp67      NodeLxcMoveVolumeVolume = "mp67"
	NodeLxcMoveVolumeVolumeMp68      NodeLxcMoveVolumeVolume = "mp68"
	NodeLxcMoveVolumeVolumeMp69      NodeLxcMoveVolumeVolume = "mp69"
	NodeLxcMoveVolumeVolumeMp70      NodeLxcMoveVolumeVolume = "mp70"
	NodeLxcMoveVolumeVolumeMp71      NodeLxcMoveVolumeVolume = "mp71"
	NodeLxcMoveVolumeVolumeMp72      NodeLxcMoveVolumeVolume = "mp72"
	NodeLxcMoveVolumeVolumeMp73      NodeLxcMoveVolumeVolume = "mp73"
	NodeLxcMoveVolumeVolumeMp74      NodeLxcMoveVolumeVolume = "mp74"
	NodeLxcMoveVolumeVolumeMp75      NodeLxcMoveVolumeVolume = "mp75"
	NodeLxcMoveVolumeVolumeMp76      NodeLxcMoveVolumeVolume = "mp76"
	NodeLxcMoveVolumeVolumeMp77      NodeLxcMoveVolumeVolume = "mp77"
	NodeLxcMoveVolumeVolumeMp78      NodeLxcMoveVolumeVolume = "mp78"
	NodeLxcMoveVolumeVolumeMp79      NodeLxcMoveVolumeVolume = "mp79"
	NodeLxcMoveVolumeVolumeMp80      NodeLxcMoveVolumeVolume = "mp80"
	NodeLxcMoveVolumeVolumeMp81      NodeLxcMoveVolumeVolume = "mp81"
	NodeLxcMoveVolumeVolumeMp82      NodeLxcMoveVolumeVolume = "mp82"
	NodeLxcMoveVolumeVolumeMp83      NodeLxcMoveVolumeVolume = "mp83"
	NodeLxcMoveVolumeVolumeMp84      NodeLxcMoveVolumeVolume = "mp84"
	NodeLxcMoveVolumeVolumeMp85      NodeLxcMoveVolumeVolume = "mp85"
	NodeLxcMoveVolumeVolumeMp86      NodeLxcMoveVolumeVolume = "mp86"
	NodeLxcMoveVolumeVolumeMp87      NodeLxcMoveVolumeVolume = "mp87"
	NodeLxcMoveVolumeVolumeMp88      NodeLxcMoveVolumeVolume = "mp88"
	NodeLxcMoveVolumeVolumeMp89      NodeLxcMoveVolumeVolume = "mp89"
	NodeLxcMoveVolumeVolumeMp90      NodeLxcMoveVolumeVolume = "mp90"
	NodeLxcMoveVolumeVolumeMp91      NodeLxcMoveVolumeVolume = "mp91"
	NodeLxcMoveVolumeVolumeMp92      NodeLxcMoveVolumeVolume = "mp92"
	NodeLxcMoveVolumeVolumeMp93      NodeLxcMoveVolumeVolume = "mp93"
	NodeLxcMoveVolumeVolumeMp94      NodeLxcMoveVolumeVolume = "mp94"
	NodeLxcMoveVolumeVolumeMp95      NodeLxcMoveVolumeVolume = "mp95"
	NodeLxcMoveVolumeVolumeMp96      NodeLxcMoveVolumeVolume = "mp96"
	NodeLxcMoveVolumeVolumeMp97      NodeLxcMoveVolumeVolume = "mp97"
	NodeLxcMoveVolumeVolumeMp98      NodeLxcMoveVolumeVolume = "mp98"
	NodeLxcMoveVolumeVolumeMp99      NodeLxcMoveVolumeVolume = "mp99"
	NodeLxcMoveVolumeVolumeMp100     NodeLxcMoveVolumeVolume = "mp100"
	NodeLxcMoveVolumeVolumeMp101     NodeLxcMoveVolumeVolume = "mp101"
	NodeLxcMoveVolumeVolumeMp102     NodeLxcMoveVolumeVolume = "mp102"
	NodeLxcMoveVolumeVolumeMp103     NodeLxcMoveVolumeVolume = "mp103"
	NodeLxcMoveVolumeVolumeMp104     NodeLxcMoveVolumeVolume = "mp104"
	NodeLxcMoveVolumeVolumeMp105     NodeLxcMoveVolumeVolume = "mp105"
	NodeLxcMoveVolumeVolumeMp106     NodeLxcMoveVolumeVolume = "mp106"
	NodeLxcMoveVolumeVolumeMp107     NodeLxcMoveVolumeVolume = "mp107"
	NodeLxcMoveVolumeVolumeMp108     NodeLxcMoveVolumeVolume = "mp108"
	NodeLxcMoveVolumeVolumeMp109     NodeLxcMoveVolumeVolume = "mp109"
	NodeLxcMoveVolumeVolumeMp110     NodeLxcMoveVolumeVolume = "mp110"
	NodeLxcMoveVolumeVolumeMp111     NodeLxcMoveVolumeVolume = "mp111"
	NodeLxcMoveVolumeVolumeMp112     NodeLxcMoveVolumeVolume = "mp112"
	NodeLxcMoveVolumeVolumeMp113     NodeLxcMoveVolumeVolume = "mp113"
	NodeLxcMoveVolumeVolumeMp114     NodeLxcMoveVolumeVolume = "mp114"
	NodeLxcMoveVolumeVolumeMp115     NodeLxcMoveVolumeVolume = "mp115"
	NodeLxcMoveVolumeVolumeMp116     NodeLxcMoveVolumeVolume = "mp116"
	NodeLxcMoveVolumeVolumeMp117     NodeLxcMoveVolumeVolume = "mp117"
	NodeLxcMoveVolumeVolumeMp118     NodeLxcMoveVolumeVolume = "mp118"
	NodeLxcMoveVolumeVolumeMp119     NodeLxcMoveVolumeVolume = "mp119"
	NodeLxcMoveVolumeVolumeMp120     NodeLxcMoveVolumeVolume = "mp120"
	NodeLxcMoveVolumeVolumeMp121     NodeLxcMoveVolumeVolume = "mp121"
	NodeLxcMoveVolumeVolumeMp122     NodeLxcMoveVolumeVolume = "mp122"
	NodeLxcMoveVolumeVolumeMp123     NodeLxcMoveVolumeVolume = "mp123"
	NodeLxcMoveVolumeVolumeMp124     NodeLxcMoveVolumeVolume = "mp124"
	NodeLxcMoveVolumeVolumeMp125     NodeLxcMoveVolumeVolume = "mp125"
	NodeLxcMoveVolumeVolumeMp126     NodeLxcMoveVolumeVolume = "mp126"
	NodeLxcMoveVolumeVolumeMp127     NodeLxcMoveVolumeVolume = "mp127"
	NodeLxcMoveVolumeVolumeMp128     NodeLxcMoveVolumeVolume = "mp128"
	NodeLxcMoveVolumeVolumeMp129     NodeLxcMoveVolumeVolume = "mp129"
	NodeLxcMoveVolumeVolumeMp130     NodeLxcMoveVolumeVolume = "mp130"
	NodeLxcMoveVolumeVolumeMp131     NodeLxcMoveVolumeVolume = "mp131"
	NodeLxcMoveVolumeVolumeMp132     NodeLxcMoveVolumeVolume = "mp132"
	NodeLxcMoveVolumeVolumeMp133     NodeLxcMoveVolumeVolume = "mp133"
	NodeLxcMoveVolumeVolumeMp134     NodeLxcMoveVolumeVolume = "mp134"
	NodeLxcMoveVolumeVolumeMp135     NodeLxcMoveVolumeVolume = "mp135"
	NodeLxcMoveVolumeVolumeMp136     NodeLxcMoveVolumeVolume = "mp136"
	NodeLxcMoveVolumeVolumeMp137     NodeLxcMoveVolumeVolume = "mp137"
	NodeLxcMoveVolumeVolumeMp138     NodeLxcMoveVolumeVolume = "mp138"
	NodeLxcMoveVolumeVolumeMp139     NodeLxcMoveVolumeVolume = "mp139"
	NodeLxcMoveVolumeVolumeMp140     NodeLxcMoveVolumeVolume = "mp140"
	NodeLxcMoveVolumeVolumeMp141     NodeLxcMoveVolumeVolume = "mp141"
	NodeLxcMoveVolumeVolumeMp142     NodeLxcMoveVolumeVolume = "mp142"
	NodeLxcMoveVolumeVolumeMp143     NodeLxcMoveVolumeVolume = "mp143"
	NodeLxcMoveVolumeVolumeMp144     NodeLxcMoveVolumeVolume = "mp144"
	NodeLxcMoveVolumeVolumeMp145     NodeLxcMoveVolumeVolume = "mp145"
	NodeLxcMoveVolumeVolumeMp146     NodeLxcMoveVolumeVolume = "mp146"
	NodeLxcMoveVolumeVolumeMp147     NodeLxcMoveVolumeVolume = "mp147"
	NodeLxcMoveVolumeVolumeMp148     NodeLxcMoveVolumeVolume = "mp148"
	NodeLxcMoveVolumeVolumeMp149     NodeLxcMoveVolumeVolume = "mp149"
	NodeLxcMoveVolumeVolumeMp150     NodeLxcMoveVolumeVolume = "mp150"
	NodeLxcMoveVolumeVolumeMp151     NodeLxcMoveVolumeVolume = "mp151"
	NodeLxcMoveVolumeVolumeMp152     NodeLxcMoveVolumeVolume = "mp152"
	NodeLxcMoveVolumeVolumeMp153     NodeLxcMoveVolumeVolume = "mp153"
	NodeLxcMoveVolumeVolumeMp154     NodeLxcMoveVolumeVolume = "mp154"
	NodeLxcMoveVolumeVolumeMp155     NodeLxcMoveVolumeVolume = "mp155"
	NodeLxcMoveVolumeVolumeMp156     NodeLxcMoveVolumeVolume = "mp156"
	NodeLxcMoveVolumeVolumeMp157     NodeLxcMoveVolumeVolume = "mp157"
	NodeLxcMoveVolumeVolumeMp158     NodeLxcMoveVolumeVolume = "mp158"
	NodeLxcMoveVolumeVolumeMp159     NodeLxcMoveVolumeVolume = "mp159"
	NodeLxcMoveVolumeVolumeMp160     NodeLxcMoveVolumeVolume = "mp160"
	NodeLxcMoveVolumeVolumeMp161     NodeLxcMoveVolumeVolume = "mp161"
	NodeLxcMoveVolumeVolumeMp162     NodeLxcMoveVolumeVolume = "mp162"
	NodeLxcMoveVolumeVolumeMp163     NodeLxcMoveVolumeVolume = "mp163"
	NodeLxcMoveVolumeVolumeMp164     NodeLxcMoveVolumeVolume = "mp164"
	NodeLxcMoveVolumeVolumeMp165     NodeLxcMoveVolumeVolume = "mp165"
	NodeLxcMoveVolumeVolumeMp166     NodeLxcMoveVolumeVolume = "mp166"
	NodeLxcMoveVolumeVolumeMp167     NodeLxcMoveVolumeVolume = "mp167"
	NodeLxcMoveVolumeVolumeMp168     NodeLxcMoveVolumeVolume = "mp168"
	NodeLxcMoveVolumeVolumeMp169     NodeLxcMoveVolumeVolume = "mp169"
	NodeLxcMoveVolumeVolumeMp170     NodeLxcMoveVolumeVolume = "mp170"
	NodeLxcMoveVolumeVolumeMp171     NodeLxcMoveVolumeVolume = "mp171"
	NodeLxcMoveVolumeVolumeMp172     NodeLxcMoveVolumeVolume = "mp172"
	NodeLxcMoveVolumeVolumeMp173     NodeLxcMoveVolumeVolume = "mp173"
	NodeLxcMoveVolumeVolumeMp174     NodeLxcMoveVolumeVolume = "mp174"
	NodeLxcMoveVolumeVolumeMp175     NodeLxcMoveVolumeVolume = "mp175"
	NodeLxcMoveVolumeVolumeMp176     NodeLxcMoveVolumeVolume = "mp176"
	NodeLxcMoveVolumeVolumeMp177     NodeLxcMoveVolumeVolume = "mp177"
	NodeLxcMoveVolumeVolumeMp178     NodeLxcMoveVolumeVolume = "mp178"
	NodeLxcMoveVolumeVolumeMp179     NodeLxcMoveVolumeVolume = "mp179"
	NodeLxcMoveVolumeVolumeMp180     NodeLxcMoveVolumeVolume = "mp180"
	NodeLxcMoveVolumeVolumeMp181     NodeLxcMoveVolumeVolume = "mp181"
	NodeLxcMoveVolumeVolumeMp182     NodeLxcMoveVolumeVolume = "mp182"
	NodeLxcMoveVolumeVolumeMp183     NodeLxcMoveVolumeVolume = "mp183"
	NodeLxcMoveVolumeVolumeMp184     NodeLxcMoveVolumeVolume = "mp184"
	NodeLxcMoveVolumeVolumeMp185     NodeLxcMoveVolumeVolume = "mp185"
	NodeLxcMoveVolumeVolumeMp186     NodeLxcMoveVolumeVolume = "mp186"
	NodeLxcMoveVolumeVolumeMp187     NodeLxcMoveVolumeVolume = "mp187"
	NodeLxcMoveVolumeVolumeMp188     NodeLxcMoveVolumeVolume = "mp188"
	NodeLxcMoveVolumeVolumeMp189     NodeLxcMoveVolumeVolume = "mp189"
	NodeLxcMoveVolumeVolumeMp190     NodeLxcMoveVolumeVolume = "mp190"
	NodeLxcMoveVolumeVolumeMp191     NodeLxcMoveVolumeVolume = "mp191"
	NodeLxcMoveVolumeVolumeMp192     NodeLxcMoveVolumeVolume = "mp192"
	NodeLxcMoveVolumeVolumeMp193     NodeLxcMoveVolumeVolume = "mp193"
	NodeLxcMoveVolumeVolumeMp194     NodeLxcMoveVolumeVolume = "mp194"
	NodeLxcMoveVolumeVolumeMp195     NodeLxcMoveVolumeVolume = "mp195"
	NodeLxcMoveVolumeVolumeMp196     NodeLxcMoveVolumeVolume = "mp196"
	NodeLxcMoveVolumeVolumeMp197     NodeLxcMoveVolumeVolume = "mp197"
	NodeLxcMoveVolumeVolumeMp198     NodeLxcMoveVolumeVolume = "mp198"
	NodeLxcMoveVolumeVolumeMp199     NodeLxcMoveVolumeVolume = "mp199"
	NodeLxcMoveVolumeVolumeMp200     NodeLxcMoveVolumeVolume = "mp200"
	NodeLxcMoveVolumeVolumeMp201     NodeLxcMoveVolumeVolume = "mp201"
	NodeLxcMoveVolumeVolumeMp202     NodeLxcMoveVolumeVolume = "mp202"
	NodeLxcMoveVolumeVolumeMp203     NodeLxcMoveVolumeVolume = "mp203"
	NodeLxcMoveVolumeVolumeMp204     NodeLxcMoveVolumeVolume = "mp204"
	NodeLxcMoveVolumeVolumeMp205     NodeLxcMoveVolumeVolume = "mp205"
	NodeLxcMoveVolumeVolumeMp206     NodeLxcMoveVolumeVolume = "mp206"
	NodeLxcMoveVolumeVolumeMp207     NodeLxcMoveVolumeVolume = "mp207"
	NodeLxcMoveVolumeVolumeMp208     NodeLxcMoveVolumeVolume = "mp208"
	NodeLxcMoveVolumeVolumeMp209     NodeLxcMoveVolumeVolume = "mp209"
	NodeLxcMoveVolumeVolumeMp210     NodeLxcMoveVolumeVolume = "mp210"
	NodeLxcMoveVolumeVolumeMp211     NodeLxcMoveVolumeVolume = "mp211"
	NodeLxcMoveVolumeVolumeMp212     NodeLxcMoveVolumeVolume = "mp212"
	NodeLxcMoveVolumeVolumeMp213     NodeLxcMoveVolumeVolume = "mp213"
	NodeLxcMoveVolumeVolumeMp214     NodeLxcMoveVolumeVolume = "mp214"
	NodeLxcMoveVolumeVolumeMp215     NodeLxcMoveVolumeVolume = "mp215"
	NodeLxcMoveVolumeVolumeMp216     NodeLxcMoveVolumeVolume = "mp216"
	NodeLxcMoveVolumeVolumeMp217     NodeLxcMoveVolumeVolume = "mp217"
	NodeLxcMoveVolumeVolumeMp218     NodeLxcMoveVolumeVolume = "mp218"
	NodeLxcMoveVolumeVolumeMp219     NodeLxcMoveVolumeVolume = "mp219"
	NodeLxcMoveVolumeVolumeMp220     NodeLxcMoveVolumeVolume = "mp220"
	NodeLxcMoveVolumeVolumeMp221     NodeLxcMoveVolumeVolume = "mp221"
	NodeLxcMoveVolumeVolumeMp222     NodeLxcMoveVolumeVolume = "mp222"
	NodeLxcMoveVolumeVolumeMp223     NodeLxcMoveVolumeVolume = "mp223"
	NodeLxcMoveVolumeVolumeMp224     NodeLxcMoveVolumeVolume = "mp224"
	NodeLxcMoveVolumeVolumeMp225     NodeLxcMoveVolumeVolume = "mp225"
	NodeLxcMoveVolumeVolumeMp226     NodeLxcMoveVolumeVolume = "mp226"
	NodeLxcMoveVolumeVolumeMp227     NodeLxcMoveVolumeVolume = "mp227"
	NodeLxcMoveVolumeVolumeMp228     NodeLxcMoveVolumeVolume = "mp228"
	NodeLxcMoveVolumeVolumeMp229     NodeLxcMoveVolumeVolume = "mp229"
	NodeLxcMoveVolumeVolumeMp230     NodeLxcMoveVolumeVolume = "mp230"
	NodeLxcMoveVolumeVolumeMp231     NodeLxcMoveVolumeVolume = "mp231"
	NodeLxcMoveVolumeVolumeMp232     NodeLxcMoveVolumeVolume = "mp232"
	NodeLxcMoveVolumeVolumeMp233     NodeLxcMoveVolumeVolume = "mp233"
	NodeLxcMoveVolumeVolumeMp234     NodeLxcMoveVolumeVolume = "mp234"
	NodeLxcMoveVolumeVolumeMp235     NodeLxcMoveVolumeVolume = "mp235"
	NodeLxcMoveVolumeVolumeMp236     NodeLxcMoveVolumeVolume = "mp236"
	NodeLxcMoveVolumeVolumeMp237     NodeLxcMoveVolumeVolume = "mp237"
	NodeLxcMoveVolumeVolumeMp238     NodeLxcMoveVolumeVolume = "mp238"
	NodeLxcMoveVolumeVolumeMp239     NodeLxcMoveVolumeVolume = "mp239"
	NodeLxcMoveVolumeVolumeMp240     NodeLxcMoveVolumeVolume = "mp240"
	NodeLxcMoveVolumeVolumeMp241     NodeLxcMoveVolumeVolume = "mp241"
	NodeLxcMoveVolumeVolumeMp242     NodeLxcMoveVolumeVolume = "mp242"
	NodeLxcMoveVolumeVolumeMp243     NodeLxcMoveVolumeVolume = "mp243"
	NodeLxcMoveVolumeVolumeMp244     NodeLxcMoveVolumeVolume = "mp244"
	NodeLxcMoveVolumeVolumeMp245     NodeLxcMoveVolumeVolume = "mp245"
	NodeLxcMoveVolumeVolumeMp246     NodeLxcMoveVolumeVolume = "mp246"
	NodeLxcMoveVolumeVolumeMp247     NodeLxcMoveVolumeVolume = "mp247"
	NodeLxcMoveVolumeVolumeMp248     NodeLxcMoveVolumeVolume = "mp248"
	NodeLxcMoveVolumeVolumeMp249     NodeLxcMoveVolumeVolume = "mp249"
	NodeLxcMoveVolumeVolumeMp250     NodeLxcMoveVolumeVolume = "mp250"
	NodeLxcMoveVolumeVolumeMp251     NodeLxcMoveVolumeVolume = "mp251"
	NodeLxcMoveVolumeVolumeMp252     NodeLxcMoveVolumeVolume = "mp252"
	NodeLxcMoveVolumeVolumeMp253     NodeLxcMoveVolumeVolume = "mp253"
	NodeLxcMoveVolumeVolumeMp254     NodeLxcMoveVolumeVolume = "mp254"
	NodeLxcMoveVolumeVolumeMp255     NodeLxcMoveVolumeVolume = "mp255"
	NodeLxcMoveVolumeVolumeUnused0   NodeLxcMoveVolumeVolume = "unused0"
	NodeLxcMoveVolumeVolumeUnused1   NodeLxcMoveVolumeVolume = "unused1"
	NodeLxcMoveVolumeVolumeUnused2   NodeLxcMoveVolumeVolume = "unused2"
	NodeLxcMoveVolumeVolumeUnused3   NodeLxcMoveVolumeVolume = "unused3"
	NodeLxcMoveVolumeVolumeUnused4   NodeLxcMoveVolumeVolume = "unused4"
	NodeLxcMoveVolumeVolumeUnused5   NodeLxcMoveVolumeVolume = "unused5"
	NodeLxcMoveVolumeVolumeUnused6   NodeLxcMoveVolumeVolume = "unused6"
	NodeLxcMoveVolumeVolumeUnused7   NodeLxcMoveVolumeVolume = "unused7"
	NodeLxcMoveVolumeVolumeUnused8   NodeLxcMoveVolumeVolume = "unused8"
	NodeLxcMoveVolumeVolumeUnused9   NodeLxcMoveVolumeVolume = "unused9"
	NodeLxcMoveVolumeVolumeUnused10  NodeLxcMoveVolumeVolume = "unused10"
	NodeLxcMoveVolumeVolumeUnused11  NodeLxcMoveVolumeVolume = "unused11"
	NodeLxcMoveVolumeVolumeUnused12  NodeLxcMoveVolumeVolume = "unused12"
	NodeLxcMoveVolumeVolumeUnused13  NodeLxcMoveVolumeVolume = "unused13"
	NodeLxcMoveVolumeVolumeUnused14  NodeLxcMoveVolumeVolume = "unused14"
	NodeLxcMoveVolumeVolumeUnused15  NodeLxcMoveVolumeVolume = "unused15"
	NodeLxcMoveVolumeVolumeUnused16  NodeLxcMoveVolumeVolume = "unused16"
	NodeLxcMoveVolumeVolumeUnused17  NodeLxcMoveVolumeVolume = "unused17"
	NodeLxcMoveVolumeVolumeUnused18  NodeLxcMoveVolumeVolume = "unused18"
	NodeLxcMoveVolumeVolumeUnused19  NodeLxcMoveVolumeVolume = "unused19"
	NodeLxcMoveVolumeVolumeUnused20  NodeLxcMoveVolumeVolume = "unused20"
	NodeLxcMoveVolumeVolumeUnused21  NodeLxcMoveVolumeVolume = "unused21"
	NodeLxcMoveVolumeVolumeUnused22  NodeLxcMoveVolumeVolume = "unused22"
	NodeLxcMoveVolumeVolumeUnused23  NodeLxcMoveVolumeVolume = "unused23"
	NodeLxcMoveVolumeVolumeUnused24  NodeLxcMoveVolumeVolume = "unused24"
	NodeLxcMoveVolumeVolumeUnused25  NodeLxcMoveVolumeVolume = "unused25"
	NodeLxcMoveVolumeVolumeUnused26  NodeLxcMoveVolumeVolume = "unused26"
	NodeLxcMoveVolumeVolumeUnused27  NodeLxcMoveVolumeVolume = "unused27"
	NodeLxcMoveVolumeVolumeUnused28  NodeLxcMoveVolumeVolume = "unused28"
	NodeLxcMoveVolumeVolumeUnused29  NodeLxcMoveVolumeVolume = "unused29"
	NodeLxcMoveVolumeVolumeUnused30  NodeLxcMoveVolumeVolume = "unused30"
	NodeLxcMoveVolumeVolumeUnused31  NodeLxcMoveVolumeVolume = "unused31"
	NodeLxcMoveVolumeVolumeUnused32  NodeLxcMoveVolumeVolume = "unused32"
	NodeLxcMoveVolumeVolumeUnused33  NodeLxcMoveVolumeVolume = "unused33"
	NodeLxcMoveVolumeVolumeUnused34  NodeLxcMoveVolumeVolume = "unused34"
	NodeLxcMoveVolumeVolumeUnused35  NodeLxcMoveVolumeVolume = "unused35"
	NodeLxcMoveVolumeVolumeUnused36  NodeLxcMoveVolumeVolume = "unused36"
	NodeLxcMoveVolumeVolumeUnused37  NodeLxcMoveVolumeVolume = "unused37"
	NodeLxcMoveVolumeVolumeUnused38  NodeLxcMoveVolumeVolume = "unused38"
	NodeLxcMoveVolumeVolumeUnused39  NodeLxcMoveVolumeVolume = "unused39"
	NodeLxcMoveVolumeVolumeUnused40  NodeLxcMoveVolumeVolume = "unused40"
	NodeLxcMoveVolumeVolumeUnused41  NodeLxcMoveVolumeVolume = "unused41"
	NodeLxcMoveVolumeVolumeUnused42  NodeLxcMoveVolumeVolume = "unused42"
	NodeLxcMoveVolumeVolumeUnused43  NodeLxcMoveVolumeVolume = "unused43"
	NodeLxcMoveVolumeVolumeUnused44  NodeLxcMoveVolumeVolume = "unused44"
	NodeLxcMoveVolumeVolumeUnused45  NodeLxcMoveVolumeVolume = "unused45"
	NodeLxcMoveVolumeVolumeUnused46  NodeLxcMoveVolumeVolume = "unused46"
	NodeLxcMoveVolumeVolumeUnused47  NodeLxcMoveVolumeVolume = "unused47"
	NodeLxcMoveVolumeVolumeUnused48  NodeLxcMoveVolumeVolume = "unused48"
	NodeLxcMoveVolumeVolumeUnused49  NodeLxcMoveVolumeVolume = "unused49"
	NodeLxcMoveVolumeVolumeUnused50  NodeLxcMoveVolumeVolume = "unused50"
	NodeLxcMoveVolumeVolumeUnused51  NodeLxcMoveVolumeVolume = "unused51"
	NodeLxcMoveVolumeVolumeUnused52  NodeLxcMoveVolumeVolume = "unused52"
	NodeLxcMoveVolumeVolumeUnused53  NodeLxcMoveVolumeVolume = "unused53"
	NodeLxcMoveVolumeVolumeUnused54  NodeLxcMoveVolumeVolume = "unused54"
	NodeLxcMoveVolumeVolumeUnused55  NodeLxcMoveVolumeVolume = "unused55"
	NodeLxcMoveVolumeVolumeUnused56  NodeLxcMoveVolumeVolume = "unused56"
	NodeLxcMoveVolumeVolumeUnused57  NodeLxcMoveVolumeVolume = "unused57"
	NodeLxcMoveVolumeVolumeUnused58  NodeLxcMoveVolumeVolume = "unused58"
	NodeLxcMoveVolumeVolumeUnused59  NodeLxcMoveVolumeVolume = "unused59"
	NodeLxcMoveVolumeVolumeUnused60  NodeLxcMoveVolumeVolume = "unused60"
	NodeLxcMoveVolumeVolumeUnused61  NodeLxcMoveVolumeVolume = "unused61"
	NodeLxcMoveVolumeVolumeUnused62  NodeLxcMoveVolumeVolume = "unused62"
	NodeLxcMoveVolumeVolumeUnused63  NodeLxcMoveVolumeVolume = "unused63"
	NodeLxcMoveVolumeVolumeUnused64  NodeLxcMoveVolumeVolume = "unused64"
	NodeLxcMoveVolumeVolumeUnused65  NodeLxcMoveVolumeVolume = "unused65"
	NodeLxcMoveVolumeVolumeUnused66  NodeLxcMoveVolumeVolume = "unused66"
	NodeLxcMoveVolumeVolumeUnused67  NodeLxcMoveVolumeVolume = "unused67"
	NodeLxcMoveVolumeVolumeUnused68  NodeLxcMoveVolumeVolume = "unused68"
	NodeLxcMoveVolumeVolumeUnused69  NodeLxcMoveVolumeVolume = "unused69"
	NodeLxcMoveVolumeVolumeUnused70  NodeLxcMoveVolumeVolume = "unused70"
	NodeLxcMoveVolumeVolumeUnused71  NodeLxcMoveVolumeVolume = "unused71"
	NodeLxcMoveVolumeVolumeUnused72  NodeLxcMoveVolumeVolume = "unused72"
	NodeLxcMoveVolumeVolumeUnused73  NodeLxcMoveVolumeVolume = "unused73"
	NodeLxcMoveVolumeVolumeUnused74  NodeLxcMoveVolumeVolume = "unused74"
	NodeLxcMoveVolumeVolumeUnused75  NodeLxcMoveVolumeVolume = "unused75"
	NodeLxcMoveVolumeVolumeUnused76  NodeLxcMoveVolumeVolume = "unused76"
	NodeLxcMoveVolumeVolumeUnused77  NodeLxcMoveVolumeVolume = "unused77"
	NodeLxcMoveVolumeVolumeUnused78  NodeLxcMoveVolumeVolume = "unused78"
	NodeLxcMoveVolumeVolumeUnused79  NodeLxcMoveVolumeVolume = "unused79"
	NodeLxcMoveVolumeVolumeUnused80  NodeLxcMoveVolumeVolume = "unused80"
	NodeLxcMoveVolumeVolumeUnused81  NodeLxcMoveVolumeVolume = "unused81"
	NodeLxcMoveVolumeVolumeUnused82  NodeLxcMoveVolumeVolume = "unused82"
	NodeLxcMoveVolumeVolumeUnused83  NodeLxcMoveVolumeVolume = "unused83"
	NodeLxcMoveVolumeVolumeUnused84  NodeLxcMoveVolumeVolume = "unused84"
	NodeLxcMoveVolumeVolumeUnused85  NodeLxcMoveVolumeVolume = "unused85"
	NodeLxcMoveVolumeVolumeUnused86  NodeLxcMoveVolumeVolume = "unused86"
	NodeLxcMoveVolumeVolumeUnused87  NodeLxcMoveVolumeVolume = "unused87"
	NodeLxcMoveVolumeVolumeUnused88  NodeLxcMoveVolumeVolume = "unused88"
	NodeLxcMoveVolumeVolumeUnused89  NodeLxcMoveVolumeVolume = "unused89"
	NodeLxcMoveVolumeVolumeUnused90  NodeLxcMoveVolumeVolume = "unused90"
	NodeLxcMoveVolumeVolumeUnused91  NodeLxcMoveVolumeVolume = "unused91"
	NodeLxcMoveVolumeVolumeUnused92  NodeLxcMoveVolumeVolume = "unused92"
	NodeLxcMoveVolumeVolumeUnused93  NodeLxcMoveVolumeVolume = "unused93"
	NodeLxcMoveVolumeVolumeUnused94  NodeLxcMoveVolumeVolume = "unused94"
	NodeLxcMoveVolumeVolumeUnused95  NodeLxcMoveVolumeVolume = "unused95"
	NodeLxcMoveVolumeVolumeUnused96  NodeLxcMoveVolumeVolume = "unused96"
	NodeLxcMoveVolumeVolumeUnused97  NodeLxcMoveVolumeVolume = "unused97"
	NodeLxcMoveVolumeVolumeUnused98  NodeLxcMoveVolumeVolume = "unused98"
	NodeLxcMoveVolumeVolumeUnused99  NodeLxcMoveVolumeVolume = "unused99"
	NodeLxcMoveVolumeVolumeUnused100 NodeLxcMoveVolumeVolume = "unused100"
	NodeLxcMoveVolumeVolumeUnused101 NodeLxcMoveVolumeVolume = "unused101"
	NodeLxcMoveVolumeVolumeUnused102 NodeLxcMoveVolumeVolume = "unused102"
	NodeLxcMoveVolumeVolumeUnused103 NodeLxcMoveVolumeVolume = "unused103"
	NodeLxcMoveVolumeVolumeUnused104 NodeLxcMoveVolumeVolume = "unused104"
	NodeLxcMoveVolumeVolumeUnused105 NodeLxcMoveVolumeVolume = "unused105"
	NodeLxcMoveVolumeVolumeUnused106 NodeLxcMoveVolumeVolume = "unused106"
	NodeLxcMoveVolumeVolumeUnused107 NodeLxcMoveVolumeVolume = "unused107"
	NodeLxcMoveVolumeVolumeUnused108 NodeLxcMoveVolumeVolume = "unused108"
	NodeLxcMoveVolumeVolumeUnused109 NodeLxcMoveVolumeVolume = "unused109"
	NodeLxcMoveVolumeVolumeUnused110 NodeLxcMoveVolumeVolume = "unused110"
	NodeLxcMoveVolumeVolumeUnused111 NodeLxcMoveVolumeVolume = "unused111"
	NodeLxcMoveVolumeVolumeUnused112 NodeLxcMoveVolumeVolume = "unused112"
	NodeLxcMoveVolumeVolumeUnused113 NodeLxcMoveVolumeVolume = "unused113"
	NodeLxcMoveVolumeVolumeUnused114 NodeLxcMoveVolumeVolume = "unused114"
	NodeLxcMoveVolumeVolumeUnused115 NodeLxcMoveVolumeVolume = "unused115"
	NodeLxcMoveVolumeVolumeUnused116 NodeLxcMoveVolumeVolume = "unused116"
	NodeLxcMoveVolumeVolumeUnused117 NodeLxcMoveVolumeVolume = "unused117"
	NodeLxcMoveVolumeVolumeUnused118 NodeLxcMoveVolumeVolume = "unused118"
	NodeLxcMoveVolumeVolumeUnused119 NodeLxcMoveVolumeVolume = "unused119"
	NodeLxcMoveVolumeVolumeUnused120 NodeLxcMoveVolumeVolume = "unused120"
	NodeLxcMoveVolumeVolumeUnused121 NodeLxcMoveVolumeVolume = "unused121"
	NodeLxcMoveVolumeVolumeUnused122 NodeLxcMoveVolumeVolume = "unused122"
	NodeLxcMoveVolumeVolumeUnused123 NodeLxcMoveVolumeVolume = "unused123"
	NodeLxcMoveVolumeVolumeUnused124 NodeLxcMoveVolumeVolume = "unused124"
	NodeLxcMoveVolumeVolumeUnused125 NodeLxcMoveVolumeVolume = "unused125"
	NodeLxcMoveVolumeVolumeUnused126 NodeLxcMoveVolumeVolume = "unused126"
	NodeLxcMoveVolumeVolumeUnused127 NodeLxcMoveVolumeVolume = "unused127"
	NodeLxcMoveVolumeVolumeUnused128 NodeLxcMoveVolumeVolume = "unused128"
	NodeLxcMoveVolumeVolumeUnused129 NodeLxcMoveVolumeVolume = "unused129"
	NodeLxcMoveVolumeVolumeUnused130 NodeLxcMoveVolumeVolume = "unused130"
	NodeLxcMoveVolumeVolumeUnused131 NodeLxcMoveVolumeVolume = "unused131"
	NodeLxcMoveVolumeVolumeUnused132 NodeLxcMoveVolumeVolume = "unused132"
	NodeLxcMoveVolumeVolumeUnused133 NodeLxcMoveVolumeVolume = "unused133"
	NodeLxcMoveVolumeVolumeUnused134 NodeLxcMoveVolumeVolume = "unused134"
	NodeLxcMoveVolumeVolumeUnused135 NodeLxcMoveVolumeVolume = "unused135"
	NodeLxcMoveVolumeVolumeUnused136 NodeLxcMoveVolumeVolume = "unused136"
	NodeLxcMoveVolumeVolumeUnused137 NodeLxcMoveVolumeVolume = "unused137"
	NodeLxcMoveVolumeVolumeUnused138 NodeLxcMoveVolumeVolume = "unused138"
	NodeLxcMoveVolumeVolumeUnused139 NodeLxcMoveVolumeVolume = "unused139"
	NodeLxcMoveVolumeVolumeUnused140 NodeLxcMoveVolumeVolume = "unused140"
	NodeLxcMoveVolumeVolumeUnused141 NodeLxcMoveVolumeVolume = "unused141"
	NodeLxcMoveVolumeVolumeUnused142 NodeLxcMoveVolumeVolume = "unused142"
	NodeLxcMoveVolumeVolumeUnused143 NodeLxcMoveVolumeVolume = "unused143"
	NodeLxcMoveVolumeVolumeUnused144 NodeLxcMoveVolumeVolume = "unused144"
	NodeLxcMoveVolumeVolumeUnused145 NodeLxcMoveVolumeVolume = "unused145"
	NodeLxcMoveVolumeVolumeUnused146 NodeLxcMoveVolumeVolume = "unused146"
	NodeLxcMoveVolumeVolumeUnused147 NodeLxcMoveVolumeVolume = "unused147"
	NodeLxcMoveVolumeVolumeUnused148 NodeLxcMoveVolumeVolume = "unused148"
	NodeLxcMoveVolumeVolumeUnused149 NodeLxcMoveVolumeVolume = "unused149"
	NodeLxcMoveVolumeVolumeUnused150 NodeLxcMoveVolumeVolume = "unused150"
	NodeLxcMoveVolumeVolumeUnused151 NodeLxcMoveVolumeVolume = "unused151"
	NodeLxcMoveVolumeVolumeUnused152 NodeLxcMoveVolumeVolume = "unused152"
	NodeLxcMoveVolumeVolumeUnused153 NodeLxcMoveVolumeVolume = "unused153"
	NodeLxcMoveVolumeVolumeUnused154 NodeLxcMoveVolumeVolume = "unused154"
	NodeLxcMoveVolumeVolumeUnused155 NodeLxcMoveVolumeVolume = "unused155"
	NodeLxcMoveVolumeVolumeUnused156 NodeLxcMoveVolumeVolume = "unused156"
	NodeLxcMoveVolumeVolumeUnused157 NodeLxcMoveVolumeVolume = "unused157"
	NodeLxcMoveVolumeVolumeUnused158 NodeLxcMoveVolumeVolume = "unused158"
	NodeLxcMoveVolumeVolumeUnused159 NodeLxcMoveVolumeVolume = "unused159"
	NodeLxcMoveVolumeVolumeUnused160 NodeLxcMoveVolumeVolume = "unused160"
	NodeLxcMoveVolumeVolumeUnused161 NodeLxcMoveVolumeVolume = "unused161"
	NodeLxcMoveVolumeVolumeUnused162 NodeLxcMoveVolumeVolume = "unused162"
	NodeLxcMoveVolumeVolumeUnused163 NodeLxcMoveVolumeVolume = "unused163"
	NodeLxcMoveVolumeVolumeUnused164 NodeLxcMoveVolumeVolume = "unused164"
	NodeLxcMoveVolumeVolumeUnused165 NodeLxcMoveVolumeVolume = "unused165"
	NodeLxcMoveVolumeVolumeUnused166 NodeLxcMoveVolumeVolume = "unused166"
	NodeLxcMoveVolumeVolumeUnused167 NodeLxcMoveVolumeVolume = "unused167"
	NodeLxcMoveVolumeVolumeUnused168 NodeLxcMoveVolumeVolume = "unused168"
	NodeLxcMoveVolumeVolumeUnused169 NodeLxcMoveVolumeVolume = "unused169"
	NodeLxcMoveVolumeVolumeUnused170 NodeLxcMoveVolumeVolume = "unused170"
	NodeLxcMoveVolumeVolumeUnused171 NodeLxcMoveVolumeVolume = "unused171"
	NodeLxcMoveVolumeVolumeUnused172 NodeLxcMoveVolumeVolume = "unused172"
	NodeLxcMoveVolumeVolumeUnused173 NodeLxcMoveVolumeVolume = "unused173"
	NodeLxcMoveVolumeVolumeUnused174 NodeLxcMoveVolumeVolume = "unused174"
	NodeLxcMoveVolumeVolumeUnused175 NodeLxcMoveVolumeVolume = "unused175"
	NodeLxcMoveVolumeVolumeUnused176 NodeLxcMoveVolumeVolume = "unused176"
	NodeLxcMoveVolumeVolumeUnused177 NodeLxcMoveVolumeVolume = "unused177"
	NodeLxcMoveVolumeVolumeUnused178 NodeLxcMoveVolumeVolume = "unused178"
	NodeLxcMoveVolumeVolumeUnused179 NodeLxcMoveVolumeVolume = "unused179"
	NodeLxcMoveVolumeVolumeUnused180 NodeLxcMoveVolumeVolume = "unused180"
	NodeLxcMoveVolumeVolumeUnused181 NodeLxcMoveVolumeVolume = "unused181"
	NodeLxcMoveVolumeVolumeUnused182 NodeLxcMoveVolumeVolume = "unused182"
	NodeLxcMoveVolumeVolumeUnused183 NodeLxcMoveVolumeVolume = "unused183"
	NodeLxcMoveVolumeVolumeUnused184 NodeLxcMoveVolumeVolume = "unused184"
	NodeLxcMoveVolumeVolumeUnused185 NodeLxcMoveVolumeVolume = "unused185"
	NodeLxcMoveVolumeVolumeUnused186 NodeLxcMoveVolumeVolume = "unused186"
	NodeLxcMoveVolumeVolumeUnused187 NodeLxcMoveVolumeVolume = "unused187"
	NodeLxcMoveVolumeVolumeUnused188 NodeLxcMoveVolumeVolume = "unused188"
	NodeLxcMoveVolumeVolumeUnused189 NodeLxcMoveVolumeVolume = "unused189"
	NodeLxcMoveVolumeVolumeUnused190 NodeLxcMoveVolumeVolume = "unused190"
	NodeLxcMoveVolumeVolumeUnused191 NodeLxcMoveVolumeVolume = "unused191"
	NodeLxcMoveVolumeVolumeUnused192 NodeLxcMoveVolumeVolume = "unused192"
	NodeLxcMoveVolumeVolumeUnused193 NodeLxcMoveVolumeVolume = "unused193"
	NodeLxcMoveVolumeVolumeUnused194 NodeLxcMoveVolumeVolume = "unused194"
	NodeLxcMoveVolumeVolumeUnused195 NodeLxcMoveVolumeVolume = "unused195"
	NodeLxcMoveVolumeVolumeUnused196 NodeLxcMoveVolumeVolume = "unused196"
	NodeLxcMoveVolumeVolumeUnused197 NodeLxcMoveVolumeVolume = "unused197"
	NodeLxcMoveVolumeVolumeUnused198 NodeLxcMoveVolumeVolume = "unused198"
	NodeLxcMoveVolumeVolumeUnused199 NodeLxcMoveVolumeVolume = "unused199"
	NodeLxcMoveVolumeVolumeUnused200 NodeLxcMoveVolumeVolume = "unused200"
	NodeLxcMoveVolumeVolumeUnused201 NodeLxcMoveVolumeVolume = "unused201"
	NodeLxcMoveVolumeVolumeUnused202 NodeLxcMoveVolumeVolume = "unused202"
	NodeLxcMoveVolumeVolumeUnused203 NodeLxcMoveVolumeVolume = "unused203"
	NodeLxcMoveVolumeVolumeUnused204 NodeLxcMoveVolumeVolume = "unused204"
	NodeLxcMoveVolumeVolumeUnused205 NodeLxcMoveVolumeVolume = "unused205"
	NodeLxcMoveVolumeVolumeUnused206 NodeLxcMoveVolumeVolume = "unused206"
	NodeLxcMoveVolumeVolumeUnused207 NodeLxcMoveVolumeVolume = "unused207"
	NodeLxcMoveVolumeVolumeUnused208 NodeLxcMoveVolumeVolume = "unused208"
	NodeLxcMoveVolumeVolumeUnused209 NodeLxcMoveVolumeVolume = "unused209"
	NodeLxcMoveVolumeVolumeUnused210 NodeLxcMoveVolumeVolume = "unused210"
	NodeLxcMoveVolumeVolumeUnused211 NodeLxcMoveVolumeVolume = "unused211"
	NodeLxcMoveVolumeVolumeUnused212 NodeLxcMoveVolumeVolume = "unused212"
	NodeLxcMoveVolumeVolumeUnused213 NodeLxcMoveVolumeVolume = "unused213"
	NodeLxcMoveVolumeVolumeUnused214 NodeLxcMoveVolumeVolume = "unused214"
	NodeLxcMoveVolumeVolumeUnused215 NodeLxcMoveVolumeVolume = "unused215"
	NodeLxcMoveVolumeVolumeUnused216 NodeLxcMoveVolumeVolume = "unused216"
	NodeLxcMoveVolumeVolumeUnused217 NodeLxcMoveVolumeVolume = "unused217"
	NodeLxcMoveVolumeVolumeUnused218 NodeLxcMoveVolumeVolume = "unused218"
	NodeLxcMoveVolumeVolumeUnused219 NodeLxcMoveVolumeVolume = "unused219"
	NodeLxcMoveVolumeVolumeUnused220 NodeLxcMoveVolumeVolume = "unused220"
	NodeLxcMoveVolumeVolumeUnused221 NodeLxcMoveVolumeVolume = "unused221"
	NodeLxcMoveVolumeVolumeUnused222 NodeLxcMoveVolumeVolume = "unused222"
	NodeLxcMoveVolumeVolumeUnused223 NodeLxcMoveVolumeVolume = "unused223"
	NodeLxcMoveVolumeVolumeUnused224 NodeLxcMoveVolumeVolume = "unused224"
	NodeLxcMoveVolumeVolumeUnused225 NodeLxcMoveVolumeVolume = "unused225"
	NodeLxcMoveVolumeVolumeUnused226 NodeLxcMoveVolumeVolume = "unused226"
	NodeLxcMoveVolumeVolumeUnused227 NodeLxcMoveVolumeVolume = "unused227"
	NodeLxcMoveVolumeVolumeUnused228 NodeLxcMoveVolumeVolume = "unused228"
	NodeLxcMoveVolumeVolumeUnused229 NodeLxcMoveVolumeVolume = "unused229"
	NodeLxcMoveVolumeVolumeUnused230 NodeLxcMoveVolumeVolume = "unused230"
	NodeLxcMoveVolumeVolumeUnused231 NodeLxcMoveVolumeVolume = "unused231"
	NodeLxcMoveVolumeVolumeUnused232 NodeLxcMoveVolumeVolume = "unused232"
	NodeLxcMoveVolumeVolumeUnused233 NodeLxcMoveVolumeVolume = "unused233"
	NodeLxcMoveVolumeVolumeUnused234 NodeLxcMoveVolumeVolume = "unused234"
	NodeLxcMoveVolumeVolumeUnused235 NodeLxcMoveVolumeVolume = "unused235"
	NodeLxcMoveVolumeVolumeUnused236 NodeLxcMoveVolumeVolume = "unused236"
	NodeLxcMoveVolumeVolumeUnused237 NodeLxcMoveVolumeVolume = "unused237"
	NodeLxcMoveVolumeVolumeUnused238 NodeLxcMoveVolumeVolume = "unused238"
	NodeLxcMoveVolumeVolumeUnused239 NodeLxcMoveVolumeVolume = "unused239"
	NodeLxcMoveVolumeVolumeUnused240 NodeLxcMoveVolumeVolume = "unused240"
	NodeLxcMoveVolumeVolumeUnused241 NodeLxcMoveVolumeVolume = "unused241"
	NodeLxcMoveVolumeVolumeUnused242 NodeLxcMoveVolumeVolume = "unused242"
	NodeLxcMoveVolumeVolumeUnused243 NodeLxcMoveVolumeVolume = "unused243"
	NodeLxcMoveVolumeVolumeUnused244 NodeLxcMoveVolumeVolume = "unused244"
	NodeLxcMoveVolumeVolumeUnused245 NodeLxcMoveVolumeVolume = "unused245"
	NodeLxcMoveVolumeVolumeUnused246 NodeLxcMoveVolumeVolume = "unused246"
	NodeLxcMoveVolumeVolumeUnused247 NodeLxcMoveVolumeVolume = "unused247"
	NodeLxcMoveVolumeVolumeUnused248 NodeLxcMoveVolumeVolume = "unused248"
	NodeLxcMoveVolumeVolumeUnused249 NodeLxcMoveVolumeVolume = "unused249"
	NodeLxcMoveVolumeVolumeUnused250 NodeLxcMoveVolumeVolume = "unused250"
	NodeLxcMoveVolumeVolumeUnused251 NodeLxcMoveVolumeVolume = "unused251"
	NodeLxcMoveVolumeVolumeUnused252 NodeLxcMoveVolumeVolume = "unused252"
	NodeLxcMoveVolumeVolumeUnused253 NodeLxcMoveVolumeVolume = "unused253"
	NodeLxcMoveVolumeVolumeUnused254 NodeLxcMoveVolumeVolume = "unused254"
	NodeLxcMoveVolumeVolumeUnused255 NodeLxcMoveVolumeVolume = "unused255"
)

type NodeCephOsdLvInfoType string

const (
	NodeCephOsdLvInfoTypeBlock NodeCephOsdLvInfoType = "block"
	NodeCephOsdLvInfoTypeDb    NodeCephOsdLvInfoType = "db"
	NodeCephOsdLvInfoTypeWal   NodeCephOsdLvInfoType = "wal"
)

type NodeCephPoolApplication string

const (
	NodeCephPoolApplicationRbd    NodeCephPoolApplication = "rbd"
	NodeCephPoolApplicationCephfs NodeCephPoolApplication = "cephfs"
	NodeCephPoolApplicationRgw    NodeCephPoolApplication = "rgw"
)

type NodeCephPoolPgAutoscaleMode string

const (
	NodeCephPoolPgAutoscaleModeOn   NodeCephPoolPgAutoscaleMode = "on"
	NodeCephPoolPgAutoscaleModeOff  NodeCephPoolPgAutoscaleMode = "off"
	NodeCephPoolPgAutoscaleModeWarn NodeCephPoolPgAutoscaleMode = "warn"
)

type NodeCephCmdSafetyAction string

const (
	NodeCephCmdSafetyActionStop    NodeCephCmdSafetyAction = "stop"
	NodeCephCmdSafetyActionDestroy NodeCephCmdSafetyAction = "destroy"
)

type NodeCephCmdSafetyService string

const (
	NodeCephCmdSafetyServiceOsd NodeCephCmdSafetyService = "osd"
	NodeCephCmdSafetyServiceMon NodeCephCmdSafetyService = "mon"
	NodeCephCmdSafetyServiceMds NodeCephCmdSafetyService = "mds"
)

type NodeServiceService string

const (
	NodeServiceServiceChrony           NodeServiceService = "chrony"
	NodeServiceServiceCorosync         NodeServiceService = "corosync"
	NodeServiceServiceCron             NodeServiceService = "cron"
	NodeServiceServiceKsmtuned         NodeServiceService = "ksmtuned"
	NodeServiceServicePostfix          NodeServiceService = "postfix"
	NodeServiceServicePveCluster       NodeServiceService = "pve-cluster"
	NodeServiceServicePveFirewall      NodeServiceService = "pve-firewall"
	NodeServiceServicePveHaCrm         NodeServiceService = "pve-ha-crm"
	NodeServiceServicePveHaLrm         NodeServiceService = "pve-ha-lrm"
	NodeServiceServicePvedaemon        NodeServiceService = "pvedaemon"
	NodeServiceServicePvefwLogger      NodeServiceService = "pvefw-logger"
	NodeServiceServicePveproxy         NodeServiceService = "pveproxy"
	NodeServiceServicePvescheduler     NodeServiceService = "pvescheduler"
	NodeServiceServicePvestatd         NodeServiceService = "pvestatd"
	NodeServiceServiceSpiceproxy       NodeServiceService = "spiceproxy"
	NodeServiceServiceSshd             NodeServiceService = "sshd"
	NodeServiceServiceSyslog           NodeServiceService = "syslog"
	NodeServiceServiceSystemdJournald  NodeServiceService = "systemd-journald"
	NodeServiceServiceSystemdTimesyncd NodeServiceService = "systemd-timesyncd"
)

type NodeNetworkType string

const (
	NodeNetworkTypeBridge         NodeNetworkType = "bridge"
	NodeNetworkTypeBond           NodeNetworkType = "bond"
	NodeNetworkTypeEth            NodeNetworkType = "eth"
	NodeNetworkTypeAlias          NodeNetworkType = "alias"
	NodeNetworkTypeVlan           NodeNetworkType = "vlan"
	NodeNetworkTypeOVSBridge      NodeNetworkType = "OVSBridge"
	NodeNetworkTypeOVSBond        NodeNetworkType = "OVSBond"
	NodeNetworkTypeOVSPort        NodeNetworkType = "OVSPort"
	NodeNetworkTypeOVSIntPort     NodeNetworkType = "OVSIntPort"
	NodeNetworkTypeAnyBridge      NodeNetworkType = "any_bridge"
	NodeNetworkTypeAnyLocalBridge NodeNetworkType = "any_local_bridge"
)

type NodeNetworkBondMode string

const (
	NodeNetworkBondModeBalanceRr      NodeNetworkBondMode = "balance-rr"
	NodeNetworkBondModeActiveBackup   NodeNetworkBondMode = "active-backup"
	NodeNetworkBondModeBalanceXor     NodeNetworkBondMode = "balance-xor"
	NodeNetworkBondModeBroadcast      NodeNetworkBondMode = "broadcast"
	NodeNetworkBondMode8023ad         NodeNetworkBondMode = "802.3ad"
	NodeNetworkBondModeBalanceTlb     NodeNetworkBondMode = "balance-tlb"
	NodeNetworkBondModeBalanceAlb     NodeNetworkBondMode = "balance-alb"
	NodeNetworkBondModeBalanceSlb     NodeNetworkBondMode = "balance-slb"
	NodeNetworkBondModeLacpBalanceSlb NodeNetworkBondMode = "lacp-balance-slb"
	NodeNetworkBondModeLacpBalanceTcp NodeNetworkBondMode = "lacp-balance-tcp"
)

type NodeNetworkBondXmitHashPolicy string

const (
	NodeNetworkBondXmitHashPolicyLayer2  NodeNetworkBondXmitHashPolicy = "layer2"
	NodeNetworkBondXmitHashPolicyLayer23 NodeNetworkBondXmitHashPolicy = "layer2+3"
	NodeNetworkBondXmitHashPolicyLayer34 NodeNetworkBondXmitHashPolicy = "layer3+4"
)

type NodeNetworkParameterType string

const (
	NodeNetworkParameterTypeBridge     NodeNetworkParameterType = "bridge"
	NodeNetworkParameterTypeBond       NodeNetworkParameterType = "bond"
	NodeNetworkParameterTypeEth        NodeNetworkParameterType = "eth"
	NodeNetworkParameterTypeAlias      NodeNetworkParameterType = "alias"
	NodeNetworkParameterTypeVlan       NodeNetworkParameterType = "vlan"
	NodeNetworkParameterTypeOVSBridge  NodeNetworkParameterType = "OVSBridge"
	NodeNetworkParameterTypeOVSBond    NodeNetworkParameterType = "OVSBond"
	NodeNetworkParameterTypeOVSPort    NodeNetworkParameterType = "OVSPort"
	NodeNetworkParameterTypeOVSIntPort NodeNetworkParameterType = "OVSIntPort"
	NodeNetworkParameterTypeUnknown    NodeNetworkParameterType = "unknown"
)

type NodeTasksSource string

const (
	NodeTasksSourceArchive NodeTasksSource = "archive"
	NodeTasksSourceActive  NodeTasksSource = "active"
	NodeTasksSourceAll     NodeTasksSource = "all"
)

type NodeStoragePrunebackupsType string

const (
	NodeStoragePrunebackupsTypeQemu NodeStoragePrunebackupsType = "qemu"
	NodeStoragePrunebackupsTypeLxc  NodeStoragePrunebackupsType = "lxc"
)

type NodeStorageContentFormat string

const (
	NodeStorageContentFormatRaw    NodeStorageContentFormat = "raw"
	NodeStorageContentFormatQcow2  NodeStorageContentFormat = "qcow2"
	NodeStorageContentFormatSubvol NodeStorageContentFormat = "subvol"
)

type NodeStorageChecksumAlgorithm string

const (
	NodeStorageChecksumAlgorithmMd5    NodeStorageChecksumAlgorithm = "md5"
	NodeStorageChecksumAlgorithmSha1   NodeStorageChecksumAlgorithm = "sha1"
	NodeStorageChecksumAlgorithmSha224 NodeStorageChecksumAlgorithm = "sha224"
	NodeStorageChecksumAlgorithmSha256 NodeStorageChecksumAlgorithm = "sha256"
	NodeStorageChecksumAlgorithmSha384 NodeStorageChecksumAlgorithm = "sha384"
	NodeStorageChecksumAlgorithmSha512 NodeStorageChecksumAlgorithm = "sha512"
)

type NodeStorageContent string

const (
	NodeStorageContentIso    NodeStorageContent = "iso"
	NodeStorageContentVztmpl NodeStorageContent = "vztmpl"
)

type NodeDisksDirectoryFilesystem string

const (
	NodeDisksDirectoryFilesystemExt4 NodeDisksDirectoryFilesystem = "ext4"
	NodeDisksDirectoryFilesystemXfs  NodeDisksDirectoryFilesystem = "xfs"
)

type NodeDisksZfsCompression string

const (
	NodeDisksZfsCompressionOn   NodeDisksZfsCompression = "on"
	NodeDisksZfsCompressionOff  NodeDisksZfsCompression = "off"
	NodeDisksZfsCompressionGzip NodeDisksZfsCompression = "gzip"
	NodeDisksZfsCompressionLz4  NodeDisksZfsCompression = "lz4"
	NodeDisksZfsCompressionLzjb NodeDisksZfsCompression = "lzjb"
	NodeDisksZfsCompressionZle  NodeDisksZfsCompression = "zle"
	NodeDisksZfsCompressionZstd NodeDisksZfsCompression = "zstd"
)

type NodeDisksZfsRaidlevel string

const (
	NodeDisksZfsRaidlevelSingle NodeDisksZfsRaidlevel = "single"
	NodeDisksZfsRaidlevelMirror NodeDisksZfsRaidlevel = "mirror"
	NodeDisksZfsRaidlevelRaid10 NodeDisksZfsRaidlevel = "raid10"
	NodeDisksZfsRaidlevelRaidz  NodeDisksZfsRaidlevel = "raidz"
	NodeDisksZfsRaidlevelRaidz2 NodeDisksZfsRaidlevel = "raidz2"
	NodeDisksZfsRaidlevelRaidz3 NodeDisksZfsRaidlevel = "raidz3"
	NodeDisksZfsRaidlevelDraid  NodeDisksZfsRaidlevel = "draid"
	NodeDisksZfsRaidlevelDraid2 NodeDisksZfsRaidlevel = "draid2"
	NodeDisksZfsRaidlevelDraid3 NodeDisksZfsRaidlevel = "draid3"
)

type NodeDisksListType string

const (
	NodeDisksListTypeUnused       NodeDisksListType = "unused"
	NodeDisksListTypeJournalDisks NodeDisksListType = "journal_disks"
)

type NodeConfigProperty string

const (
	NodeConfigPropertyAcme                NodeConfigProperty = "acme"
	NodeConfigPropertyAcmedomain0         NodeConfigProperty = "acmedomain0"
	NodeConfigPropertyAcmedomain1         NodeConfigProperty = "acmedomain1"
	NodeConfigPropertyAcmedomain2         NodeConfigProperty = "acmedomain2"
	NodeConfigPropertyAcmedomain3         NodeConfigProperty = "acmedomain3"
	NodeConfigPropertyAcmedomain4         NodeConfigProperty = "acmedomain4"
	NodeConfigPropertyAcmedomain5         NodeConfigProperty = "acmedomain5"
	NodeConfigPropertyDescription         NodeConfigProperty = "description"
	NodeConfigPropertyStartallOnbootDelay NodeConfigProperty = "startall-onboot-delay"
	NodeConfigPropertyWakeonlan           NodeConfigProperty = "wakeonlan"
)

type NodeStatusCommand string

const (
	NodeStatusCommandReboot   NodeStatusCommand = "reboot"
	NodeStatusCommandShutdown NodeStatusCommand = "shutdown"
)

type NodeCmd string

const (
	NodeCmdUpgrade     NodeCmd = "upgrade"
	NodeCmdCephInstall NodeCmd = "ceph_install"
	NodeCmdLogin       NodeCmd = "login"
)

type StorageType string

const (
	StorageTypeBtrfs       StorageType = "btrfs"
	StorageTypeCephfs      StorageType = "cephfs"
	StorageTypeCifs        StorageType = "cifs"
	StorageTypeDir         StorageType = "dir"
	StorageTypeGlusterfs   StorageType = "glusterfs"
	StorageTypeIscsi       StorageType = "iscsi"
	StorageTypeIscsidirect StorageType = "iscsidirect"
	StorageTypeLvm         StorageType = "lvm"
	StorageTypeLvmthin     StorageType = "lvmthin"
	StorageTypeNfs         StorageType = "nfs"
	StorageTypePbs         StorageType = "pbs"
	StorageTypeRbd         StorageType = "rbd"
	StorageTypeZfs         StorageType = "zfs"
	StorageTypeZfspool     StorageType = "zfspool"
)

type StoragePreallocation string

const (
	StoragePreallocationOff      StoragePreallocation = "off"
	StoragePreallocationMetadata StoragePreallocation = "metadata"
	StoragePreallocationFalloc   StoragePreallocation = "falloc"
	StoragePreallocationFull     StoragePreallocation = "full"
)

type StorageSmbversion string

const (
	StorageSmbversionDefault StorageSmbversion = "default"
	StorageSmbversion20      StorageSmbversion = "2.0"
	StorageSmbversion21      StorageSmbversion = "2.1"
	StorageSmbversion3       StorageSmbversion = "3"
	StorageSmbversion30      StorageSmbversion = "3.0"
	StorageSmbversion311     StorageSmbversion = "3.11"
)

type StorageTransport string

const (
	StorageTransportTcp  StorageTransport = "tcp"
	StorageTransportRdma StorageTransport = "rdma"
	StorageTransportUnix StorageTransport = "unix"
)

type AccessDomainMode string

const (
	AccessDomainModeLdap         AccessDomainMode = "ldap"
	AccessDomainModeLdaps        AccessDomainMode = "ldaps"
	AccessDomainModeLdapStarttls AccessDomainMode = "ldap+starttls"
)

type AccessDomainSslversion string

const (
	AccessDomainSslversionTlsv1  AccessDomainSslversion = "tlsv1"
	AccessDomainSslversionTlsv11 AccessDomainSslversion = "tlsv1_1"
	AccessDomainSslversionTlsv12 AccessDomainSslversion = "tlsv1_2"
	AccessDomainSslversionTlsv13 AccessDomainSslversion = "tlsv1_3"
)

type AccessDomainsType string

const (
	AccessDomainsTypeAd     AccessDomainsType = "ad"
	AccessDomainsTypeLdap   AccessDomainsType = "ldap"
	AccessDomainsTypeOpenid AccessDomainsType = "openid"
	AccessDomainsTypePam    AccessDomainsType = "pam"
	AccessDomainsTypePve    AccessDomainsType = "pve"
)

type AccessTfaType string

const (
	AccessTfaTypeTotp     AccessTfaType = "totp"
	AccessTfaTypeU2f      AccessTfaType = "u2f"
	AccessTfaTypeWebauthn AccessTfaType = "webauthn"
	AccessTfaTypeRecovery AccessTfaType = "recovery"
	AccessTfaTypeYubico   AccessTfaType = "yubico"
)

type PoolType string

const (
	PoolTypeQemu    PoolType = "qemu"
	PoolTypeLxc     PoolType = "lxc"
	PoolTypeStorage PoolType = "storage"
)

type GetClusterResponse struct {
}

func (c *Client) GetCluster() (*GetClusterResponse, error) {

}

type GetClusterReplicationResponse struct {
}

func (c *Client) GetClusterReplication() (*GetClusterReplicationResponse, error) {

}

type PostClusterReplicationParameters struct {
	// Description.
	Comment *string `json:"comment"`
	// Flag to disable/deactivate the entry.
	Disable *bool `json:"disable"`
	// Replication Job ID. The ID is composed of a Guest ID and a job number, separated by a hyphen, i.e. '<GUEST>-<JOBNUM>'.
	Id string `json:"id"`
	// Rate limit in mbps (megabytes per second) as floating point number.
	Rate *float64 `json:"rate"`
	// Mark the replication job for removal. The job will remove all local replication snapshots. When set to 'full', it also tries to remove replicated volumes on the target. The job then removes itself from the configuration file.
	RemoveJob *ClusterReplicationRemoveJob `json:"remove_job"`
	// Storage replication schedule. The format is a subset of `systemd` calendar events.
	Schedule *string `json:"schedule"`
	// For internal use, to detect if the guest was stolen.
	Source *string `json:"source"`
	// Target node.
	Target string `json:"target"`
	// Section type.
	Type ClusterReplicationType `json:"type"`
}

type PostClusterReplicationResponse struct {
}

func (c *Client) PostClusterReplication(parameters PostClusterReplicationParameters) (*PostClusterReplicationResponse, error) {

}

type DeleteClusterReplicationByIdParameters struct {
	// Will remove the jobconfig entry, but will not cleanup.
	Force *bool `json:"force"`
	// Keep replicated data at target (do not remove).
	Keep *bool `json:"keep"`
}

type DeleteClusterReplicationByIdResponse struct {
}

func (c *Client) DeleteClusterReplicationById(Id string, parameters DeleteClusterReplicationByIdParameters) (*DeleteClusterReplicationByIdResponse, error) {

}

type GetClusterReplicationByIdResponse struct {
}

func (c *Client) GetClusterReplicationById(Id string) (*GetClusterReplicationByIdResponse, error) {

}

type PutClusterReplicationByIdParameters struct {
	// Description.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Flag to disable/deactivate the entry.
	Disable *bool `json:"disable"`
	// Rate limit in mbps (megabytes per second) as floating point number.
	Rate *float64 `json:"rate"`
	// Mark the replication job for removal. The job will remove all local replication snapshots. When set to 'full', it also tries to remove replicated volumes on the target. The job then removes itself from the configuration file.
	RemoveJob *ClusterReplicationRemoveJob `json:"remove_job"`
	// Storage replication schedule. The format is a subset of `systemd` calendar events.
	Schedule *string `json:"schedule"`
	// For internal use, to detect if the guest was stolen.
	Source *string `json:"source"`
}

type PutClusterReplicationByIdResponse struct {
}

func (c *Client) PutClusterReplicationById(Id string, parameters PutClusterReplicationByIdParameters) (*PutClusterReplicationByIdResponse, error) {

}

type GetClusterMetricsResponse struct {
}

func (c *Client) GetClusterMetrics() (*GetClusterMetricsResponse, error) {

}

type GetClusterMetricsServerResponse struct {
}

func (c *Client) GetClusterMetricsServer() (*GetClusterMetricsServerResponse, error) {

}

type DeleteClusterMetricsServerByIdResponse struct {
}

func (c *Client) DeleteClusterMetricsServerById(Id string) (*DeleteClusterMetricsServerByIdResponse, error) {

}

type GetClusterMetricsServerByIdResponse struct {
}

func (c *Client) GetClusterMetricsServerById(Id string) (*GetClusterMetricsServerByIdResponse, error) {

}

type PostClusterMetricsServerByIdParameters struct {
	// An API path prefix inserted between '<host>:<port>/' and '/api2/'. Can be useful if the InfluxDB service runs behind a reverse proxy.
	ApiPathPrefix *string `json:"api-path-prefix"`
	// The InfluxDB bucket/db. Only necessary when using the http v2 api.
	Bucket *string `json:"bucket"`
	// Flag to disable the plugin.
	Disable       *bool                              `json:"disable"`
	Influxdbproto *ClusterMetricsServerInfluxdbproto `json:"influxdbproto"`
	// InfluxDB max-body-size in bytes. Requests are batched up to this size.
	MaxBodySize *int `json:"max-body-size"`
	// MTU for metrics transmission over UDP
	Mtu *int `json:"mtu"`
	// The InfluxDB organization. Only necessary when using the http v2 api. Has no meaning when using v2 compatibility api.
	Organization *string `json:"organization"`
	// root graphite path (ex: proxmox.mycluster.mykey)
	Path *string `json:"path"`
	// server network port
	Port int `json:"port"`
	// Protocol to send graphite data. TCP or UDP (default)
	Proto *ClusterMetricsServerProto `json:"proto"`
	// server dns name or IP address
	Server string `json:"server"`
	// graphite TCP socket timeout (default=1)
	Timeout *int `json:"timeout"`
	// The InfluxDB access token. Only necessary when using the http v2 api. If the v2 compatibility api is used, use 'user:password' instead.
	Token *string `json:"token"`
	// Plugin type.
	Type ClusterMetricsServerType `json:"type"`
	// Set to 0 to disable certificate verification for https endpoints.
	VerifyCertificate *bool `json:"verify-certificate"`
}

type PostClusterMetricsServerByIdResponse struct {
}

func (c *Client) PostClusterMetricsServerById(Id string, parameters PostClusterMetricsServerByIdParameters) (*PostClusterMetricsServerByIdResponse, error) {

}

type PutClusterMetricsServerByIdParameters struct {
	// An API path prefix inserted between '<host>:<port>/' and '/api2/'. Can be useful if the InfluxDB service runs behind a reverse proxy.
	ApiPathPrefix *string `json:"api-path-prefix"`
	// The InfluxDB bucket/db. Only necessary when using the http v2 api.
	Bucket *string `json:"bucket"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Flag to disable the plugin.
	Disable       *bool                              `json:"disable"`
	Influxdbproto *ClusterMetricsServerInfluxdbproto `json:"influxdbproto"`
	// InfluxDB max-body-size in bytes. Requests are batched up to this size.
	MaxBodySize *int `json:"max-body-size"`
	// MTU for metrics transmission over UDP
	Mtu *int `json:"mtu"`
	// The InfluxDB organization. Only necessary when using the http v2 api. Has no meaning when using v2 compatibility api.
	Organization *string `json:"organization"`
	// root graphite path (ex: proxmox.mycluster.mykey)
	Path *string `json:"path"`
	// server network port
	Port int `json:"port"`
	// Protocol to send graphite data. TCP or UDP (default)
	Proto *ClusterMetricsServerProto `json:"proto"`
	// server dns name or IP address
	Server string `json:"server"`
	// graphite TCP socket timeout (default=1)
	Timeout *int `json:"timeout"`
	// The InfluxDB access token. Only necessary when using the http v2 api. If the v2 compatibility api is used, use 'user:password' instead.
	Token *string `json:"token"`
	// Set to 0 to disable certificate verification for https endpoints.
	VerifyCertificate *bool `json:"verify-certificate"`
}

type PutClusterMetricsServerByIdResponse struct {
}

func (c *Client) PutClusterMetricsServerById(Id string, parameters PutClusterMetricsServerByIdParameters) (*PutClusterMetricsServerByIdResponse, error) {

}

type GetClusterConfigResponse struct {
}

func (c *Client) GetClusterConfig() (*GetClusterConfigResponse, error) {

}

type PostClusterConfigParameters struct {
	// The name of the cluster.
	Clustername string `json:"clustername"`
	// Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	LinkN *string `json:"link[n]"`
	// Node id for this node.
	Nodeid *int `json:"nodeid"`
	// Number of votes for this node.
	Votes *int `json:"votes"`
}

type PostClusterConfigResponse struct {
}

func (c *Client) PostClusterConfig(parameters PostClusterConfigParameters) (*PostClusterConfigResponse, error) {

}

type GetClusterConfigApiversionResponse struct {
}

func (c *Client) GetClusterConfigApiversion() (*GetClusterConfigApiversionResponse, error) {

}

type GetClusterConfigNodesResponse struct {
}

func (c *Client) GetClusterConfigNodes() (*GetClusterConfigNodesResponse, error) {

}

type DeleteClusterConfigNodeResponse struct {
}

func (c *Client) DeleteClusterConfigNode(Node string) (*DeleteClusterConfigNodeResponse, error) {

}

type PostClusterConfigNodeParameters struct {
	// The JOIN_API_VERSION of the new node.
	Apiversion *int `json:"apiversion"`
	// Do not throw error if node already exists.
	Force *bool `json:"force"`
	// Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	LinkN *string `json:"link[n]"`
	// IP Address of node to add. Used as fallback if no links are given.
	NewNodeIp *string `json:"new_node_ip"`
	// Node id for this node.
	Nodeid *int `json:"nodeid"`
	// Number of votes for this node
	Votes *int `json:"votes"`
}

type PostClusterConfigNodeResponse struct {
}

func (c *Client) PostClusterConfigNode(Node string, parameters PostClusterConfigNodeParameters) (*PostClusterConfigNodeResponse, error) {

}

type GetClusterConfigJoinParameters struct {
	// The node for which the joinee gets the nodeinfo.
	Node *string `json:"node"`
}

type GetClusterConfigJoinResponse struct {
}

func (c *Client) GetClusterConfigJoin(parameters GetClusterConfigJoinParameters) (*GetClusterConfigJoinResponse, error) {

}

type PostClusterConfigJoinParameters struct {
	// Certificate SHA 256 fingerprint.
	Fingerprint string `json:"fingerprint"`
	// Do not throw error if node already exists.
	Force *bool `json:"force"`
	// Hostname (or IP) of an existing cluster member.
	Hostname string `json:"hostname"`
	// Address and priority information of a single corosync link. (up to 8 links supported; link0..link7)
	LinkN *string `json:"link[n]"`
	// Node id for this node.
	Nodeid *int `json:"nodeid"`
	// Superuser (root) password of peer node.
	Password string `json:"password"`
	// Number of votes for this node
	Votes *int `json:"votes"`
}

type PostClusterConfigJoinResponse struct {
}

func (c *Client) PostClusterConfigJoin(parameters PostClusterConfigJoinParameters) (*PostClusterConfigJoinResponse, error) {

}

type GetClusterConfigTotemResponse struct {
}

func (c *Client) GetClusterConfigTotem() (*GetClusterConfigTotemResponse, error) {

}

type GetClusterConfigQdeviceResponse struct {
}

func (c *Client) GetClusterConfigQdevice() (*GetClusterConfigQdeviceResponse, error) {

}

type GetClusterFirewallResponse struct {
}

func (c *Client) GetClusterFirewall() (*GetClusterFirewallResponse, error) {

}

type GetClusterFirewallGroupsResponse struct {
}

func (c *Client) GetClusterFirewallGroups() (*GetClusterFirewallGroupsResponse, error) {

}

type PostClusterFirewallGroupsParameters struct {
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Security Group name.
	Group string `json:"group"`
	// Rename/update an existing security group. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing group.
	Rename *string `json:"rename"`
}

type PostClusterFirewallGroupsResponse struct {
}

func (c *Client) PostClusterFirewallGroups(parameters PostClusterFirewallGroupsParameters) (*PostClusterFirewallGroupsResponse, error) {

}

type DeleteClusterFirewallGroupResponse struct {
}

func (c *Client) DeleteClusterFirewallGroup(Group string) (*DeleteClusterFirewallGroupResponse, error) {

}

type GetClusterFirewallGroupResponse struct {
}

func (c *Client) GetClusterFirewallGroup(Group string) (*GetClusterFirewallGroupResponse, error) {

}

type PostClusterFirewallGroupParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Update rule at position <pos>.
	Pos *int `json:"pos"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type FirewallType `json:"type"`
}

type PostClusterFirewallGroupResponse struct {
}

func (c *Client) PostClusterFirewallGroup(Group string, parameters PostClusterFirewallGroupParameters) (*PostClusterFirewallGroupResponse, error) {

}

type DeleteClusterFirewallGroupByPosParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteClusterFirewallGroupByPosResponse struct {
}

func (c *Client) DeleteClusterFirewallGroupByPos(Group string, Pos int, parameters DeleteClusterFirewallGroupByPosParameters) (*DeleteClusterFirewallGroupByPosResponse, error) {

}

type GetClusterFirewallGroupByPosResponse struct {
}

func (c *Client) GetClusterFirewallGroupByPos(Group string, Pos int) (*GetClusterFirewallGroupByPosResponse, error) {

}

type PutClusterFirewallGroupByPosParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action *string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Move rule to new position <moveto>. Other arguments are ignored.
	Moveto *int `json:"moveto"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type *FirewallType `json:"type"`
}

type PutClusterFirewallGroupByPosResponse struct {
}

func (c *Client) PutClusterFirewallGroupByPos(Group string, Pos int, parameters PutClusterFirewallGroupByPosParameters) (*PutClusterFirewallGroupByPosResponse, error) {

}

type GetClusterFirewallRulesResponse struct {
}

func (c *Client) GetClusterFirewallRules() (*GetClusterFirewallRulesResponse, error) {

}

type PostClusterFirewallRulesParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Update rule at position <pos>.
	Pos *int `json:"pos"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type FirewallType `json:"type"`
}

type PostClusterFirewallRulesResponse struct {
}

func (c *Client) PostClusterFirewallRules(parameters PostClusterFirewallRulesParameters) (*PostClusterFirewallRulesResponse, error) {

}

type DeleteClusterFirewallRuleParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteClusterFirewallRuleResponse struct {
}

func (c *Client) DeleteClusterFirewallRule(Pos int, parameters DeleteClusterFirewallRuleParameters) (*DeleteClusterFirewallRuleResponse, error) {

}

type GetClusterFirewallRuleResponse struct {
}

func (c *Client) GetClusterFirewallRule(Pos int) (*GetClusterFirewallRuleResponse, error) {

}

type PutClusterFirewallRuleParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action *string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Move rule to new position <moveto>. Other arguments are ignored.
	Moveto *int `json:"moveto"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type *FirewallType `json:"type"`
}

type PutClusterFirewallRuleResponse struct {
}

func (c *Client) PutClusterFirewallRule(Pos int, parameters PutClusterFirewallRuleParameters) (*PutClusterFirewallRuleResponse, error) {

}

type GetClusterFirewallIpsetResponse struct {
}

func (c *Client) GetClusterFirewallIpset() (*GetClusterFirewallIpsetResponse, error) {

}

type PostClusterFirewallIpsetParameters struct {
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// IP set name.
	Name string `json:"name"`
	// Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
	Rename *string `json:"rename"`
}

type PostClusterFirewallIpsetResponse struct {
}

func (c *Client) PostClusterFirewallIpset(parameters PostClusterFirewallIpsetParameters) (*PostClusterFirewallIpsetResponse, error) {

}

type DeleteClusterFirewallIpsetByNameParameters struct {
	// Delete all members of the IPSet, if there are any.
	Force *bool `json:"force"`
}

type DeleteClusterFirewallIpsetByNameResponse struct {
}

func (c *Client) DeleteClusterFirewallIpsetByName(Name string, parameters DeleteClusterFirewallIpsetByNameParameters) (*DeleteClusterFirewallIpsetByNameResponse, error) {

}

type GetClusterFirewallIpsetByNameResponse struct {
}

func (c *Client) GetClusterFirewallIpsetByName(Name string) (*GetClusterFirewallIpsetByNameResponse, error) {

}

type PostClusterFirewallIpsetByNameParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	Nomatch *bool   `json:"nomatch"`
}

type PostClusterFirewallIpsetByNameResponse struct {
}

func (c *Client) PostClusterFirewallIpsetByName(Name string, parameters PostClusterFirewallIpsetByNameParameters) (*PostClusterFirewallIpsetByNameResponse, error) {

}

type DeleteClusterFirewallIpsetByNameByCidrParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteClusterFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) DeleteClusterFirewallIpsetByNameByCidr(Name string, Cidr string, parameters DeleteClusterFirewallIpsetByNameByCidrParameters) (*DeleteClusterFirewallIpsetByNameByCidrResponse, error) {

}

type GetClusterFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) GetClusterFirewallIpsetByNameByCidr(Name string, Cidr string) (*GetClusterFirewallIpsetByNameByCidrResponse, error) {

}

type PutClusterFirewallIpsetByNameByCidrParameters struct {
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest  *string `json:"digest"`
	Nomatch *bool   `json:"nomatch"`
}

type PutClusterFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) PutClusterFirewallIpsetByNameByCidr(Name string, Cidr string, parameters PutClusterFirewallIpsetByNameByCidrParameters) (*PutClusterFirewallIpsetByNameByCidrResponse, error) {

}

type GetClusterFirewallAliasesResponse struct {
}

func (c *Client) GetClusterFirewallAliases() (*GetClusterFirewallAliasesResponse, error) {

}

type PostClusterFirewallAliasesParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	// Alias name.
	Name string `json:"name"`
}

type PostClusterFirewallAliasesResponse struct {
}

func (c *Client) PostClusterFirewallAliases(parameters PostClusterFirewallAliasesParameters) (*PostClusterFirewallAliasesResponse, error) {

}

type DeleteClusterFirewallAliasParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteClusterFirewallAliasResponse struct {
}

func (c *Client) DeleteClusterFirewallAlias(Name string, parameters DeleteClusterFirewallAliasParameters) (*DeleteClusterFirewallAliasResponse, error) {

}

type GetClusterFirewallAliasResponse struct {
}

func (c *Client) GetClusterFirewallAlias(Name string) (*GetClusterFirewallAliasResponse, error) {

}

type PutClusterFirewallAliasParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Rename an existing alias.
	Rename *string `json:"rename"`
}

type PutClusterFirewallAliasResponse struct {
}

func (c *Client) PutClusterFirewallAlias(Name string, parameters PutClusterFirewallAliasParameters) (*PutClusterFirewallAliasResponse, error) {

}

type GetClusterFirewallOptionsResponse struct {
}

func (c *Client) GetClusterFirewallOptions() (*GetClusterFirewallOptionsResponse, error) {

}

type PutClusterFirewallOptionsParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Enable ebtables rules cluster wide.
	Ebtables *bool `json:"ebtables"`
	// Enable or disable the firewall cluster wide.
	Enable *int `json:"enable"`
	// Log ratelimiting settings
	LogRatelimit *string `json:"log_ratelimit"`
	// Input policy.
	PolicyIn *FirewallOptionsPolicy `json:"policy_in"`
	// Output policy.
	PolicyOut *FirewallOptionsPolicy `json:"policy_out"`
}

type PutClusterFirewallOptionsResponse struct {
}

func (c *Client) PutClusterFirewallOptions(parameters PutClusterFirewallOptionsParameters) (*PutClusterFirewallOptionsResponse, error) {

}

type GetClusterFirewallMacrosResponse struct {
}

func (c *Client) GetClusterFirewallMacros() (*GetClusterFirewallMacrosResponse, error) {

}

type GetClusterFirewallRefsParameters struct {
	// Only list references of specified type.
	Type *FirewallRefsType `json:"type"`
}

type GetClusterFirewallRefsResponse struct {
}

func (c *Client) GetClusterFirewallRefs(parameters GetClusterFirewallRefsParameters) (*GetClusterFirewallRefsResponse, error) {

}

type GetClusterBackupResponse struct {
}

func (c *Client) GetClusterBackup() (*GetClusterBackupResponse, error) {

}

type PostClusterBackupParameters struct {
	// Backup all known guest systems on this host.
	All *bool `json:"all"`
	// Limit I/O bandwidth (in KiB/s).
	Bwlimit *int `json:"bwlimit"`
	// Description for the Job.
	Comment *string `json:"comment"`
	// Compress dump file.
	Compress *Compress `json:"compress"`
	// Day of week selection.
	Dow *string `json:"dow"`
	// Store resulting files to specified directory.
	Dumpdir *string `json:"dumpdir"`
	// Enable or disable the job.
	Enabled *bool `json:"enabled"`
	// Exclude specified guest systems (assumes --all)
	Exclude *string `json:"exclude"`
	// Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	ExcludePath *any `json:"exclude-path"`
	// Job ID (will be autogenerated).
	Id *string `json:"id"`
	// Set IO priority when using the BFQ scheduler. For snapshot and suspend mode backups of VMs, this only affects the compressor. A value of 8 means the idle priority is used, otherwise the best-effort priority is used with the specified value.
	Ionice *int `json:"ionice"`
	// Maximal time to wait for the global lock (minutes).
	Lockwait *int `json:"lockwait"`
	// Specify when to send an email
	Mailnotification *Mailnotification `json:"mailnotification"`
	// Comma-separated list of email addresses or users that should receive email notifications.
	Mailto *string `json:"mailto"`
	// Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Maxfiles *int `json:"maxfiles"`
	// Backup mode.
	Mode *Mode `json:"mode"`
	// Only run if executed on this node.
	Node *string `json:"node"`
	// Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	NotesTemplate *string `json:"notes-template"`
	// Other performance-related settings.
	Performance *string `json:"performance"`
	// Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pigz *int `json:"pigz"`
	// Backup all known guest systems included in the specified pool.
	Pool *string `json:"pool"`
	// If true, mark backup(s) as protected.
	Protected *bool `json:"protected"`
	// Use these retention options instead of those from the storage configuration.
	PruneBackups *string `json:"prune-backups"`
	// Be quiet.
	Quiet *bool `json:"quiet"`
	// Prune older backups according to 'prune-backups'.
	Remove *bool `json:"remove"`
	// If true, the job will be run as soon as possible if it was missed while the scheduler was not running.
	RepeatMissed *bool `json:"repeat-missed"`
	// Backup schedule. The format is a subset of `systemd` calendar events.
	Schedule *string `json:"schedule"`
	// Use specified hook script.
	Script *string `json:"script"`
	// Job Start time.
	Starttime *string `json:"starttime"`
	// Exclude temporary files and logs.
	Stdexcludes *bool `json:"stdexcludes"`
	// Stop running backup jobs on this host.
	Stop *bool `json:"stop"`
	// Maximal time to wait until a guest system is stopped (minutes).
	Stopwait *int `json:"stopwait"`
	// Store resulting file to this storage.
	Storage *string `json:"storage"`
	// Store temporary files to specified directory.
	Tmpdir *string `json:"tmpdir"`
	// The ID of the guest system you want to backup.
	Vmid *string `json:"vmid"`
	// Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
	Zstd *int `json:"zstd"`
}

type PostClusterBackupResponse struct {
}

func (c *Client) PostClusterBackup(parameters PostClusterBackupParameters) (*PostClusterBackupResponse, error) {

}

type DeleteClusterBackupByIdResponse struct {
}

func (c *Client) DeleteClusterBackupById(Id string) (*DeleteClusterBackupByIdResponse, error) {

}

type GetClusterBackupByIdResponse struct {
}

func (c *Client) GetClusterBackupById(Id string) (*GetClusterBackupByIdResponse, error) {

}

type PutClusterBackupByIdParameters struct {
	// Backup all known guest systems on this host.
	All *bool `json:"all"`
	// Limit I/O bandwidth (in KiB/s).
	Bwlimit *int `json:"bwlimit"`
	// Description for the Job.
	Comment *string `json:"comment"`
	// Compress dump file.
	Compress *Compress `json:"compress"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Day of week selection.
	Dow *string `json:"dow"`
	// Store resulting files to specified directory.
	Dumpdir *string `json:"dumpdir"`
	// Enable or disable the job.
	Enabled *bool `json:"enabled"`
	// Exclude specified guest systems (assumes --all)
	Exclude *string `json:"exclude"`
	// Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	ExcludePath *any `json:"exclude-path"`
	// Set IO priority when using the BFQ scheduler. For snapshot and suspend mode backups of VMs, this only affects the compressor. A value of 8 means the idle priority is used, otherwise the best-effort priority is used with the specified value.
	Ionice *int `json:"ionice"`
	// Maximal time to wait for the global lock (minutes).
	Lockwait *int `json:"lockwait"`
	// Specify when to send an email
	Mailnotification *Mailnotification `json:"mailnotification"`
	// Comma-separated list of email addresses or users that should receive email notifications.
	Mailto *string `json:"mailto"`
	// Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Maxfiles *int `json:"maxfiles"`
	// Backup mode.
	Mode *Mode `json:"mode"`
	// Only run if executed on this node.
	Node *string `json:"node"`
	// Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	NotesTemplate *string `json:"notes-template"`
	// Other performance-related settings.
	Performance *string `json:"performance"`
	// Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pigz *int `json:"pigz"`
	// Backup all known guest systems included in the specified pool.
	Pool *string `json:"pool"`
	// If true, mark backup(s) as protected.
	Protected *bool `json:"protected"`
	// Use these retention options instead of those from the storage configuration.
	PruneBackups *string `json:"prune-backups"`
	// Be quiet.
	Quiet *bool `json:"quiet"`
	// Prune older backups according to 'prune-backups'.
	Remove *bool `json:"remove"`
	// If true, the job will be run as soon as possible if it was missed while the scheduler was not running.
	RepeatMissed *bool `json:"repeat-missed"`
	// Backup schedule. The format is a subset of `systemd` calendar events.
	Schedule *string `json:"schedule"`
	// Use specified hook script.
	Script *string `json:"script"`
	// Job Start time.
	Starttime *string `json:"starttime"`
	// Exclude temporary files and logs.
	Stdexcludes *bool `json:"stdexcludes"`
	// Stop running backup jobs on this host.
	Stop *bool `json:"stop"`
	// Maximal time to wait until a guest system is stopped (minutes).
	Stopwait *int `json:"stopwait"`
	// Store resulting file to this storage.
	Storage *string `json:"storage"`
	// Store temporary files to specified directory.
	Tmpdir *string `json:"tmpdir"`
	// The ID of the guest system you want to backup.
	Vmid *string `json:"vmid"`
	// Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
	Zstd *int `json:"zstd"`
}

type PutClusterBackupByIdResponse struct {
}

func (c *Client) PutClusterBackupById(Id string, parameters PutClusterBackupByIdParameters) (*PutClusterBackupByIdResponse, error) {

}

type GetClusterBackupByIdIncludedVolumesResponse struct {
}

func (c *Client) GetClusterBackupByIdIncludedVolumes(Id string) (*GetClusterBackupByIdIncludedVolumesResponse, error) {

}

type GetClusterBackupInfoResponse struct {
}

func (c *Client) GetClusterBackupInfo() (*GetClusterBackupInfoResponse, error) {

}

type GetClusterBackupInfoNotBackedUpResponse struct {
}

func (c *Client) GetClusterBackupInfoNotBackedUp() (*GetClusterBackupInfoNotBackedUpResponse, error) {

}

type GetClusterHaResponse struct {
}

func (c *Client) GetClusterHa() (*GetClusterHaResponse, error) {

}

type GetClusterHaResourcesParameters struct {
	// Only list resources of specific type
	Type *ClusterHaResourcesType `json:"type"`
}

type GetClusterHaResourcesResponse struct {
}

func (c *Client) GetClusterHaResources(parameters GetClusterHaResourcesParameters) (*GetClusterHaResourcesResponse, error) {

}

type PostClusterHaResourcesParameters struct {
	// Description.
	Comment *string `json:"comment"`
	// The HA group identifier.
	Group *string `json:"group"`
	// Maximal number of service relocate tries when a service failes to start.
	MaxRelocate *int `json:"max_relocate"`
	// Maximal number of tries to restart the service on a node after its start failed.
	MaxRestart *int `json:"max_restart"`
	// HA resource ID. This consists of a resource type followed by a resource specific name, separated with colon (example: vm:100 / ct:100). For virtual machines and containers, you can simply use the VM or CT id as a shortcut (example: 100).
	Sid string `json:"sid"`
	// Requested resource state.
	State *ClusterHaResourceState `json:"state"`
	// Resource type.
	Type *ClusterHaResourcesType `json:"type"`
}

type PostClusterHaResourcesResponse struct {
}

func (c *Client) PostClusterHaResources(parameters PostClusterHaResourcesParameters) (*PostClusterHaResourcesResponse, error) {

}

type DeleteClusterHaResourceResponse struct {
}

func (c *Client) DeleteClusterHaResource(Sid string) (*DeleteClusterHaResourceResponse, error) {

}

type GetClusterHaResourceResponse struct {
}

func (c *Client) GetClusterHaResource(Sid string) (*GetClusterHaResourceResponse, error) {

}

type PutClusterHaResourceParameters struct {
	// Description.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// The HA group identifier.
	Group *string `json:"group"`
	// Maximal number of service relocate tries when a service failes to start.
	MaxRelocate *int `json:"max_relocate"`
	// Maximal number of tries to restart the service on a node after its start failed.
	MaxRestart *int `json:"max_restart"`
	// Requested resource state.
	State *ClusterHaResourceState `json:"state"`
}

type PutClusterHaResourceResponse struct {
}

func (c *Client) PutClusterHaResource(Sid string, parameters PutClusterHaResourceParameters) (*PutClusterHaResourceResponse, error) {

}

type PostClusterHaResourceMigrateParameters struct {
	// Target node.
	Node string `json:"node"`
}

type PostClusterHaResourceMigrateResponse struct {
}

func (c *Client) PostClusterHaResourceMigrate(Sid string, parameters PostClusterHaResourceMigrateParameters) (*PostClusterHaResourceMigrateResponse, error) {

}

type PostClusterHaResourceRelocateParameters struct {
	// Target node.
	Node string `json:"node"`
}

type PostClusterHaResourceRelocateResponse struct {
}

func (c *Client) PostClusterHaResourceRelocate(Sid string, parameters PostClusterHaResourceRelocateParameters) (*PostClusterHaResourceRelocateResponse, error) {

}

type GetClusterHaGroupsResponse struct {
}

func (c *Client) GetClusterHaGroups() (*GetClusterHaGroupsResponse, error) {

}

type PostClusterHaGroupsParameters struct {
	// Description.
	Comment *string `json:"comment"`
	// The HA group identifier.
	Group string `json:"group"`
	// List of cluster node names with optional priority.
	Nodes string `json:"nodes"`
	// The CRM tries to run services on the node with the highest priority. If a node with higher priority comes online, the CRM migrates the service to that node. Enabling nofailback prevents that behavior.
	Nofailback *bool `json:"nofailback"`
	// Resources bound to restricted groups may only run on nodes defined by the group.
	Restricted *bool `json:"restricted"`
	// Group type.
	Type *ClusterHaGroupsType `json:"type"`
}

type PostClusterHaGroupsResponse struct {
}

func (c *Client) PostClusterHaGroups(parameters PostClusterHaGroupsParameters) (*PostClusterHaGroupsResponse, error) {

}

type DeleteClusterHaGroupResponse struct {
}

func (c *Client) DeleteClusterHaGroup(Group string) (*DeleteClusterHaGroupResponse, error) {

}

type GetClusterHaGroupResponse struct {
}

func (c *Client) GetClusterHaGroup(Group string) (*GetClusterHaGroupResponse, error) {

}

type PutClusterHaGroupParameters struct {
	// Description.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// List of cluster node names with optional priority.
	Nodes *string `json:"nodes"`
	// The CRM tries to run services on the node with the highest priority. If a node with higher priority comes online, the CRM migrates the service to that node. Enabling nofailback prevents that behavior.
	Nofailback *bool `json:"nofailback"`
	// Resources bound to restricted groups may only run on nodes defined by the group.
	Restricted *bool `json:"restricted"`
}

type PutClusterHaGroupResponse struct {
}

func (c *Client) PutClusterHaGroup(Group string, parameters PutClusterHaGroupParameters) (*PutClusterHaGroupResponse, error) {

}

type GetClusterHaStatusResponse struct {
}

func (c *Client) GetClusterHaStatus() (*GetClusterHaStatusResponse, error) {

}

type GetClusterHaStatusCurrentResponse struct {
}

func (c *Client) GetClusterHaStatusCurrent() (*GetClusterHaStatusCurrentResponse, error) {

}

type GetClusterHaStatusManagerStatusResponse struct {
}

func (c *Client) GetClusterHaStatusManagerStatus() (*GetClusterHaStatusManagerStatusResponse, error) {

}

type GetClusterAcmeResponse struct {
}

func (c *Client) GetClusterAcme() (*GetClusterAcmeResponse, error) {

}

type GetClusterAcmePluginsParameters struct {
	// Only list ACME plugins of a specific type
	Type *ClusterAcmePluginsType `json:"type"`
}

type GetClusterAcmePluginsResponse struct {
}

func (c *Client) GetClusterAcmePlugins(parameters GetClusterAcmePluginsParameters) (*GetClusterAcmePluginsResponse, error) {

}

type PostClusterAcmePluginsParameters struct {
	// API plugin name
	Api *ClusterAcmePluginApi `json:"api"`
	// DNS plugin data. (base64 encoded)
	Data *string `json:"data"`
	// Flag to disable the config.
	Disable *bool `json:"disable"`
	// ACME Plugin ID name
	Id string `json:"id"`
	// List of cluster node names.
	Nodes *string `json:"nodes"`
	// ACME challenge type.
	Type ClusterAcmePluginsType `json:"type"`
	// Extra delay in seconds to wait before requesting validation. Allows to cope with a long TTL of DNS records.
	ValidationDelay *int `json:"validation-delay"`
}

type PostClusterAcmePluginsResponse struct {
}

func (c *Client) PostClusterAcmePlugins(parameters PostClusterAcmePluginsParameters) (*PostClusterAcmePluginsResponse, error) {

}

type DeleteClusterAcmePluginResponse struct {
}

func (c *Client) DeleteClusterAcmePlugin(Id string) (*DeleteClusterAcmePluginResponse, error) {

}

type GetClusterAcmePluginResponse struct {
}

func (c *Client) GetClusterAcmePlugin(Id string) (*GetClusterAcmePluginResponse, error) {

}

type PutClusterAcmePluginParameters struct {
	// API plugin name
	Api *ClusterAcmePluginApi `json:"api"`
	// DNS plugin data. (base64 encoded)
	Data *string `json:"data"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Flag to disable the config.
	Disable *bool `json:"disable"`
	// List of cluster node names.
	Nodes *string `json:"nodes"`
	// Extra delay in seconds to wait before requesting validation. Allows to cope with a long TTL of DNS records.
	ValidationDelay *int `json:"validation-delay"`
}

type PutClusterAcmePluginResponse struct {
}

func (c *Client) PutClusterAcmePlugin(Id string, parameters PutClusterAcmePluginParameters) (*PutClusterAcmePluginResponse, error) {

}

type GetClusterAcmeAccountResponse struct {
}

func (c *Client) GetClusterAcmeAccount() (*GetClusterAcmeAccountResponse, error) {

}

type PostClusterAcmeAccountParameters struct {
	// Contact email addresses.
	Contact string `json:"contact"`
	// URL of ACME CA directory endpoint.
	Directory *string `json:"directory"`
	// ACME account config file name.
	Name *string `json:"name"`
	// URL of CA TermsOfService - setting this indicates agreement.
	TosUrl *string `json:"tos_url"`
}

type PostClusterAcmeAccountResponse struct {
}

func (c *Client) PostClusterAcmeAccount(parameters PostClusterAcmeAccountParameters) (*PostClusterAcmeAccountResponse, error) {

}

type DeleteClusterAcmeAccountByNameResponse struct {
}

func (c *Client) DeleteClusterAcmeAccountByName(Name string) (*DeleteClusterAcmeAccountByNameResponse, error) {

}

type GetClusterAcmeAccountByNameResponse struct {
}

func (c *Client) GetClusterAcmeAccountByName(Name string) (*GetClusterAcmeAccountByNameResponse, error) {

}

type PutClusterAcmeAccountByNameParameters struct {
	// Contact email addresses.
	Contact *string `json:"contact"`
}

type PutClusterAcmeAccountByNameResponse struct {
}

func (c *Client) PutClusterAcmeAccountByName(Name string, parameters PutClusterAcmeAccountByNameParameters) (*PutClusterAcmeAccountByNameResponse, error) {

}

type GetClusterAcmeTosParameters struct {
	// URL of ACME CA directory endpoint.
	Directory *string `json:"directory"`
}

type GetClusterAcmeTosResponse struct {
}

func (c *Client) GetClusterAcmeTos(parameters GetClusterAcmeTosParameters) (*GetClusterAcmeTosResponse, error) {

}

type GetClusterAcmeDirectoriesResponse struct {
}

func (c *Client) GetClusterAcmeDirectories() (*GetClusterAcmeDirectoriesResponse, error) {

}

type GetClusterAcmeChallengeSchemaResponse struct {
}

func (c *Client) GetClusterAcmeChallengeSchema() (*GetClusterAcmeChallengeSchemaResponse, error) {

}

type GetClusterCephResponse struct {
}

func (c *Client) GetClusterCeph() (*GetClusterCephResponse, error) {

}

type GetClusterCephMetadataParameters struct {
	Scope *ClusterCephMetadataScope `json:"scope"`
}

type GetClusterCephMetadataResponse struct {
}

func (c *Client) GetClusterCephMetadata(parameters GetClusterCephMetadataParameters) (*GetClusterCephMetadataResponse, error) {

}

type GetClusterCephStatusResponse struct {
}

func (c *Client) GetClusterCephStatus() (*GetClusterCephStatusResponse, error) {

}

type GetClusterCephFlagsResponse struct {
}

func (c *Client) GetClusterCephFlags() (*GetClusterCephFlagsResponse, error) {

}

type PutClusterCephFlagsParameters struct {
	// Backfilling of PGs is suspended.
	Nobackfill *bool `json:"nobackfill"`
	// Deep Scrubbing is disabled.
	NodeepScrub *bool `json:"nodeep-scrub"`
	// OSD failure reports are being ignored, such that the monitors will not mark OSDs down.
	Nodown *bool `json:"nodown"`
	// OSDs that were previously marked out will not be marked back in when they start.
	Noin *bool `json:"noin"`
	// OSDs will not automatically be marked out after the configured interval.
	Noout *bool `json:"noout"`
	// Rebalancing of PGs is suspended.
	Norebalance *bool `json:"norebalance"`
	// Recovery of PGs is suspended.
	Norecover *bool `json:"norecover"`
	// Scrubbing is disabled.
	Noscrub *bool `json:"noscrub"`
	// Cache tiering activity is suspended.
	Notieragent *bool `json:"notieragent"`
	// OSDs are not allowed to start.
	Noup *bool `json:"noup"`
	// Pauses read and writes.
	Pause *bool `json:"pause"`
}

type PutClusterCephFlagsResponse struct {
}

func (c *Client) PutClusterCephFlags(parameters PutClusterCephFlagsParameters) (*PutClusterCephFlagsResponse, error) {

}

type GetClusterCephFlagResponse struct {
}

func (c *Client) GetClusterCephFlag(Flag ClusterCephFlagFlag) (*GetClusterCephFlagResponse, error) {

}

type PutClusterCephFlagParameters struct {
	// The new value of the flag
	Value bool `json:"value"`
}

type PutClusterCephFlagResponse struct {
}

func (c *Client) PutClusterCephFlag(Flag ClusterCephFlagFlag, parameters PutClusterCephFlagParameters) (*PutClusterCephFlagResponse, error) {

}

type GetClusterJobsResponse struct {
}

func (c *Client) GetClusterJobs() (*GetClusterJobsResponse, error) {

}

type GetClusterJobsRealmSyncResponse struct {
}

func (c *Client) GetClusterJobsRealmSync() (*GetClusterJobsRealmSyncResponse, error) {

}

type DeleteClusterJobsRealmSyncByIdResponse struct {
}

func (c *Client) DeleteClusterJobsRealmSyncById(Id string) (*DeleteClusterJobsRealmSyncByIdResponse, error) {

}

type GetClusterJobsRealmSyncByIdResponse struct {
}

func (c *Client) GetClusterJobsRealmSyncById(Id string) (*GetClusterJobsRealmSyncByIdResponse, error) {

}

type PostClusterJobsRealmSyncByIdParameters struct {
	// Description for the Job.
	Comment *string `json:"comment"`
	// Enable newly synced users immediately.
	EnableNew *bool `json:"enable-new"`
	// Determines if the job is enabled.
	Enabled *bool `json:"enabled"`
	// Authentication domain ID
	Realm *string `json:"realm"`
	// A semicolon-seperated list of things to remove when they or the user vanishes during a sync. The following values are possible: 'entry' removes the user/group when not returned from the sync. 'properties' removes the set properties on existing user/group that do not appear in the source (even custom ones). 'acl' removes acls when the user/group is not returned from the sync. Instead of a list it also can be 'none' (the default).
	RemoveVanished *string `json:"remove-vanished"`
	// Backup schedule. The format is a subset of `systemd` calendar events.
	Schedule string `json:"schedule"`
	// Select what to sync.
	Scope *SyncScope `json:"scope"`
}

type PostClusterJobsRealmSyncByIdResponse struct {
}

func (c *Client) PostClusterJobsRealmSyncById(Id string, parameters PostClusterJobsRealmSyncByIdParameters) (*PostClusterJobsRealmSyncByIdResponse, error) {

}

type PutClusterJobsRealmSyncByIdParameters struct {
	// Description for the Job.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Enable newly synced users immediately.
	EnableNew *bool `json:"enable-new"`
	// Determines if the job is enabled.
	Enabled *bool `json:"enabled"`
	// A semicolon-seperated list of things to remove when they or the user vanishes during a sync. The following values are possible: 'entry' removes the user/group when not returned from the sync. 'properties' removes the set properties on existing user/group that do not appear in the source (even custom ones). 'acl' removes acls when the user/group is not returned from the sync. Instead of a list it also can be 'none' (the default).
	RemoveVanished *string `json:"remove-vanished"`
	// Backup schedule. The format is a subset of `systemd` calendar events.
	Schedule string `json:"schedule"`
	// Select what to sync.
	Scope *SyncScope `json:"scope"`
}

type PutClusterJobsRealmSyncByIdResponse struct {
}

func (c *Client) PutClusterJobsRealmSyncById(Id string, parameters PutClusterJobsRealmSyncByIdParameters) (*PutClusterJobsRealmSyncByIdResponse, error) {

}

type GetClusterJobsScheduleAnalyzeParameters struct {
	// Number of event-iteration to simulate and return.
	Iterations *int `json:"iterations"`
	// Job schedule. The format is a subset of `systemd` calendar events.
	Schedule string `json:"schedule"`
	// UNIX timestamp to start the calculation from. Defaults to the current time.
	Starttime *int `json:"starttime"`
}

type GetClusterJobsScheduleAnalyzeResponse struct {
}

func (c *Client) GetClusterJobsScheduleAnalyze(parameters GetClusterJobsScheduleAnalyzeParameters) (*GetClusterJobsScheduleAnalyzeResponse, error) {

}

type GetClusterMappingResponse struct {
}

func (c *Client) GetClusterMapping() (*GetClusterMappingResponse, error) {

}

type GetClusterMappingPciParameters struct {
	// If given, checks the configurations on the given node for correctness, and adds relevant diagnostics for the devices to the response.
	CheckNode *string `json:"check-node"`
}

type GetClusterMappingPciResponse struct {
}

func (c *Client) GetClusterMappingPci(parameters GetClusterMappingPciParameters) (*GetClusterMappingPciResponse, error) {

}

type PostClusterMappingPciParameters struct {
	// Description of the logical PCI device.
	Description *string `json:"description"`
	// The ID of the logical PCI mapping.
	Id string `json:"id"`
	// A list of maps for the cluster nodes.
	Map  any   `json:"map"`
	Mdev *bool `json:"mdev"`
}

type PostClusterMappingPciResponse struct {
}

func (c *Client) PostClusterMappingPci(parameters PostClusterMappingPciParameters) (*PostClusterMappingPciResponse, error) {

}

type DeleteClusterMappingPciByIdResponse struct {
}

func (c *Client) DeleteClusterMappingPciById(Id string) (*DeleteClusterMappingPciByIdResponse, error) {

}

type GetClusterMappingPciByIdResponse struct {
}

func (c *Client) GetClusterMappingPciById(Id string) (*GetClusterMappingPciByIdResponse, error) {

}

type PutClusterMappingPciByIdParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Description of the logical PCI device.
	Description *string `json:"description"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// A list of maps for the cluster nodes.
	Map  *any  `json:"map"`
	Mdev *bool `json:"mdev"`
}

type PutClusterMappingPciByIdResponse struct {
}

func (c *Client) PutClusterMappingPciById(Id string, parameters PutClusterMappingPciByIdParameters) (*PutClusterMappingPciByIdResponse, error) {

}

type GetClusterMappingUsbParameters struct {
	// If given, checks the configurations on the given node for correctness, and adds relevant errors to the devices.
	CheckNode *string `json:"check-node"`
}

type GetClusterMappingUsbResponse struct {
}

func (c *Client) GetClusterMappingUsb(parameters GetClusterMappingUsbParameters) (*GetClusterMappingUsbResponse, error) {

}

type PostClusterMappingUsbParameters struct {
	// Description of the logical PCI device.
	Description *string `json:"description"`
	// The ID of the logical PCI mapping.
	Id string `json:"id"`
	// A list of maps for the cluster nodes.
	Map any `json:"map"`
}

type PostClusterMappingUsbResponse struct {
}

func (c *Client) PostClusterMappingUsb(parameters PostClusterMappingUsbParameters) (*PostClusterMappingUsbResponse, error) {

}

type DeleteClusterMappingUsbByIdResponse struct {
}

func (c *Client) DeleteClusterMappingUsbById(Id string) (*DeleteClusterMappingUsbByIdResponse, error) {

}

type GetClusterMappingUsbByIdResponse struct {
}

func (c *Client) GetClusterMappingUsbById(Id string) (*GetClusterMappingUsbByIdResponse, error) {

}

type PutClusterMappingUsbByIdParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Description of the logical PCI device.
	Description *string `json:"description"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// A list of maps for the cluster nodes.
	Map any `json:"map"`
}

type PutClusterMappingUsbByIdResponse struct {
}

func (c *Client) PutClusterMappingUsbById(Id string, parameters PutClusterMappingUsbByIdParameters) (*PutClusterMappingUsbByIdResponse, error) {

}

type GetClusterSdnResponse struct {
}

func (c *Client) GetClusterSdn() (*GetClusterSdnResponse, error) {

}

type PutClusterSdnResponse struct {
}

func (c *Client) PutClusterSdn() (*PutClusterSdnResponse, error) {

}

type GetClusterSdnVnetsParameters struct {
	// Display pending config.
	Pending *bool `json:"pending"`
	// Display running config.
	Running *bool `json:"running"`
}

type GetClusterSdnVnetsResponse struct {
}

func (c *Client) GetClusterSdnVnets(parameters GetClusterSdnVnetsParameters) (*GetClusterSdnVnetsResponse, error) {

}

type PostClusterSdnVnetsParameters struct {
	// alias name of the vnet
	Alias *string `json:"alias"`
	// vlan or vxlan id
	Tag *int `json:"tag"`
	// Type
	Type *ClusterSdnVnetsType `json:"type"`
	// Allow vm VLANs to pass through this vnet.
	Vlanaware *bool `json:"vlanaware"`
	// The SDN vnet object identifier.
	Vnet string `json:"vnet"`
	// zone id
	Zone string `json:"zone"`
}

type PostClusterSdnVnetsResponse struct {
}

func (c *Client) PostClusterSdnVnets(parameters PostClusterSdnVnetsParameters) (*PostClusterSdnVnetsResponse, error) {

}

type DeleteClusterSdnVnetResponse struct {
}

func (c *Client) DeleteClusterSdnVnet(Vnet string) (*DeleteClusterSdnVnetResponse, error) {

}

type GetClusterSdnVnetParameters struct {
	// Display pending config.
	Pending *bool `json:"pending"`
	// Display running config.
	Running *bool `json:"running"`
}

type GetClusterSdnVnetResponse struct {
}

func (c *Client) GetClusterSdnVnet(Vnet string, parameters GetClusterSdnVnetParameters) (*GetClusterSdnVnetResponse, error) {

}

type PutClusterSdnVnetParameters struct {
	// alias name of the vnet
	Alias *string `json:"alias"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// vlan or vxlan id
	Tag *int `json:"tag"`
	// Allow vm VLANs to pass through this vnet.
	Vlanaware *bool `json:"vlanaware"`
	// zone id
	Zone *string `json:"zone"`
}

type PutClusterSdnVnetResponse struct {
}

func (c *Client) PutClusterSdnVnet(Vnet string, parameters PutClusterSdnVnetParameters) (*PutClusterSdnVnetResponse, error) {

}

type GetClusterSdnVnetSubnetsParameters struct {
	// Display pending config.
	Pending *bool `json:"pending"`
	// Display running config.
	Running *bool `json:"running"`
}

type GetClusterSdnVnetSubnetsResponse struct {
}

func (c *Client) GetClusterSdnVnetSubnets(Vnet string, parameters GetClusterSdnVnetSubnetsParameters) (*GetClusterSdnVnetSubnetsResponse, error) {

}

type PostClusterSdnVnetSubnetsParameters struct {
	// dns domain zone prefix  ex: 'adm' -> <hostname>.adm.mydomain.com
	Dnszoneprefix *string `json:"dnszoneprefix"`
	// Subnet Gateway: Will be assign on vnet for layer3 zones
	Gateway *string `json:"gateway"`
	// enable masquerade for this subnet if pve-firewall
	Snat *bool `json:"snat"`
	// The SDN subnet object identifier.
	Subnet string                    `json:"subnet"`
	Type   ClusterSdnVnetSubnetsType `json:"type"`
}

type PostClusterSdnVnetSubnetsResponse struct {
}

func (c *Client) PostClusterSdnVnetSubnets(Vnet string, parameters PostClusterSdnVnetSubnetsParameters) (*PostClusterSdnVnetSubnetsResponse, error) {

}

type DeleteClusterSdnVnetSubnetResponse struct {
}

func (c *Client) DeleteClusterSdnVnetSubnet(Vnet string, Subnet string) (*DeleteClusterSdnVnetSubnetResponse, error) {

}

type GetClusterSdnVnetSubnetParameters struct {
	// Display pending config.
	Pending *bool `json:"pending"`
	// Display running config.
	Running *bool `json:"running"`
}

type GetClusterSdnVnetSubnetResponse struct {
}

func (c *Client) GetClusterSdnVnetSubnet(Vnet string, Subnet string, parameters GetClusterSdnVnetSubnetParameters) (*GetClusterSdnVnetSubnetResponse, error) {

}

type PutClusterSdnVnetSubnetParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// dns domain zone prefix  ex: 'adm' -> <hostname>.adm.mydomain.com
	Dnszoneprefix *string `json:"dnszoneprefix"`
	// Subnet Gateway: Will be assign on vnet for layer3 zones
	Gateway *string `json:"gateway"`
	// enable masquerade for this subnet if pve-firewall
	Snat *bool `json:"snat"`
}

type PutClusterSdnVnetSubnetResponse struct {
}

func (c *Client) PutClusterSdnVnetSubnet(Vnet string, Subnet string, parameters PutClusterSdnVnetSubnetParameters) (*PutClusterSdnVnetSubnetResponse, error) {

}

type GetClusterSdnZonesParameters struct {
	// Display pending config.
	Pending *bool `json:"pending"`
	// Display running config.
	Running *bool `json:"running"`
	// Only list SDN zones of specific type
	Type *ClusterSdnZonesType `json:"type"`
}

type GetClusterSdnZonesResponse struct {
}

func (c *Client) GetClusterSdnZones(parameters GetClusterSdnZonesParameters) (*GetClusterSdnZonesResponse, error) {

}

type PostClusterSdnZonesParameters struct {
	// Advertise evpn subnets if you have silent hosts
	AdvertiseSubnets *bool   `json:"advertise-subnets"`
	Bridge           *string `json:"bridge"`
	// Disable auto mac learning.
	BridgeDisableMacLearning *bool `json:"bridge-disable-mac-learning"`
	// Frr router name
	Controller *string `json:"controller"`
	// Disable ipv4 arp && ipv6 neighbour discovery suppression
	DisableArpNdSuppression *bool `json:"disable-arp-nd-suppression"`
	// dns api server
	Dns *string `json:"dns"`
	// dns domain zone  ex: mydomain.com
	Dnszone *string `json:"dnszone"`
	// Faucet dataplane id
	DpId *int `json:"dp-id"`
	// List of cluster node names.
	Exitnodes *string `json:"exitnodes"`
	// Allow exitnodes to connect to evpn guests
	ExitnodesLocalRouting *bool `json:"exitnodes-local-routing"`
	// Force traffic to this exitnode first.
	ExitnodesPrimary *string `json:"exitnodes-primary"`
	// use a specific ipam
	Ipam *string `json:"ipam"`
	// Anycast logical router mac address
	Mac *string `json:"mac"`
	// MTU
	Mtu *int `json:"mtu"`
	// List of cluster node names.
	Nodes *string `json:"nodes"`
	// peers address list.
	Peers *string `json:"peers"`
	// reverse dns api server
	Reversedns *string `json:"reversedns"`
	// Route-Target import
	RtImport *string `json:"rt-import"`
	// Service-VLAN Tag
	Tag *int `json:"tag"`
	// Plugin type.
	Type         ClusterSdnZonesType         `json:"type"`
	VlanProtocol *ClusterSdnZoneVlanProtocol `json:"vlan-protocol"`
	// l3vni.
	VrfVxlan *int `json:"vrf-vxlan"`
	// Vxlan tunnel udp port (default 4789).
	VxlanPort *int `json:"vxlan-port"`
	// The SDN zone object identifier.
	Zone string `json:"zone"`
}

type PostClusterSdnZonesResponse struct {
}

func (c *Client) PostClusterSdnZones(parameters PostClusterSdnZonesParameters) (*PostClusterSdnZonesResponse, error) {

}

type DeleteClusterSdnZoneResponse struct {
}

func (c *Client) DeleteClusterSdnZone(Zone string) (*DeleteClusterSdnZoneResponse, error) {

}

type GetClusterSdnZoneParameters struct {
	// Display pending config.
	Pending *bool `json:"pending"`
	// Display running config.
	Running *bool `json:"running"`
}

type GetClusterSdnZoneResponse struct {
}

func (c *Client) GetClusterSdnZone(Zone string, parameters GetClusterSdnZoneParameters) (*GetClusterSdnZoneResponse, error) {

}

type PutClusterSdnZoneParameters struct {
	// Advertise evpn subnets if you have silent hosts
	AdvertiseSubnets *bool   `json:"advertise-subnets"`
	Bridge           *string `json:"bridge"`
	// Disable auto mac learning.
	BridgeDisableMacLearning *bool `json:"bridge-disable-mac-learning"`
	// Frr router name
	Controller *string `json:"controller"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Disable ipv4 arp && ipv6 neighbour discovery suppression
	DisableArpNdSuppression *bool `json:"disable-arp-nd-suppression"`
	// dns api server
	Dns *string `json:"dns"`
	// dns domain zone  ex: mydomain.com
	Dnszone *string `json:"dnszone"`
	// Faucet dataplane id
	DpId *int `json:"dp-id"`
	// List of cluster node names.
	Exitnodes *string `json:"exitnodes"`
	// Allow exitnodes to connect to evpn guests
	ExitnodesLocalRouting *bool `json:"exitnodes-local-routing"`
	// Force traffic to this exitnode first.
	ExitnodesPrimary *string `json:"exitnodes-primary"`
	// use a specific ipam
	Ipam *string `json:"ipam"`
	// Anycast logical router mac address
	Mac *string `json:"mac"`
	// MTU
	Mtu *int `json:"mtu"`
	// List of cluster node names.
	Nodes *string `json:"nodes"`
	// peers address list.
	Peers *string `json:"peers"`
	// reverse dns api server
	Reversedns *string `json:"reversedns"`
	// Route-Target import
	RtImport *string `json:"rt-import"`
	// Service-VLAN Tag
	Tag          *int                        `json:"tag"`
	VlanProtocol *ClusterSdnZoneVlanProtocol `json:"vlan-protocol"`
	// l3vni.
	VrfVxlan *int `json:"vrf-vxlan"`
	// Vxlan tunnel udp port (default 4789).
	VxlanPort *int `json:"vxlan-port"`
}

type PutClusterSdnZoneResponse struct {
}

func (c *Client) PutClusterSdnZone(Zone string, parameters PutClusterSdnZoneParameters) (*PutClusterSdnZoneResponse, error) {

}

type GetClusterSdnControllersParameters struct {
	// Display pending config.
	Pending *bool `json:"pending"`
	// Display running config.
	Running *bool `json:"running"`
	// Only list sdn controllers of specific type
	Type *ClusterSdnControllersType `json:"type"`
}

type GetClusterSdnControllersResponse struct {
}

func (c *Client) GetClusterSdnControllers(parameters GetClusterSdnControllersParameters) (*GetClusterSdnControllersResponse, error) {

}

type PostClusterSdnControllersParameters struct {
	// autonomous system number
	Asn                     *int  `json:"asn"`
	BgpMultipathAsPathRelax *bool `json:"bgp-multipath-as-path-relax"`
	// The SDN controller object identifier.
	Controller string `json:"controller"`
	// Enable ebgp. (remote-as external)
	Ebgp         *bool `json:"ebgp"`
	EbgpMultihop *int  `json:"ebgp-multihop"`
	// source loopback interface.
	Loopback *string `json:"loopback"`
	// The cluster node name.
	Node *string `json:"node"`
	// peers address list.
	Peers *string `json:"peers"`
	// Plugin type.
	Type ClusterSdnControllersType `json:"type"`
}

type PostClusterSdnControllersResponse struct {
}

func (c *Client) PostClusterSdnControllers(parameters PostClusterSdnControllersParameters) (*PostClusterSdnControllersResponse, error) {

}

type DeleteClusterSdnControllerResponse struct {
}

func (c *Client) DeleteClusterSdnController(Controller string) (*DeleteClusterSdnControllerResponse, error) {

}

type GetClusterSdnControllerParameters struct {
	// Display pending config.
	Pending *bool `json:"pending"`
	// Display running config.
	Running *bool `json:"running"`
}

type GetClusterSdnControllerResponse struct {
}

func (c *Client) GetClusterSdnController(Controller string, parameters GetClusterSdnControllerParameters) (*GetClusterSdnControllerResponse, error) {

}

type PutClusterSdnControllerParameters struct {
	// autonomous system number
	Asn                     *int  `json:"asn"`
	BgpMultipathAsPathRelax *bool `json:"bgp-multipath-as-path-relax"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Enable ebgp. (remote-as external)
	Ebgp         *bool `json:"ebgp"`
	EbgpMultihop *int  `json:"ebgp-multihop"`
	// source loopback interface.
	Loopback *string `json:"loopback"`
	// The cluster node name.
	Node *string `json:"node"`
	// peers address list.
	Peers *string `json:"peers"`
}

type PutClusterSdnControllerResponse struct {
}

func (c *Client) PutClusterSdnController(Controller string, parameters PutClusterSdnControllerParameters) (*PutClusterSdnControllerResponse, error) {

}

type GetClusterSdnIpamsParameters struct {
	// Only list sdn ipams of specific type
	Type *ClusterSdnIpamsType `json:"type"`
}

type GetClusterSdnIpamsResponse struct {
}

func (c *Client) GetClusterSdnIpams(parameters GetClusterSdnIpamsParameters) (*GetClusterSdnIpamsResponse, error) {

}

type PostClusterSdnIpamsParameters struct {
	// The SDN ipam object identifier.
	Ipam    string  `json:"ipam"`
	Section *int    `json:"section"`
	Token   *string `json:"token"`
	// Plugin type.
	Type ClusterSdnIpamsType `json:"type"`
	Url  *string             `json:"url"`
}

type PostClusterSdnIpamsResponse struct {
}

func (c *Client) PostClusterSdnIpams(parameters PostClusterSdnIpamsParameters) (*PostClusterSdnIpamsResponse, error) {

}

type DeleteClusterSdnIpamResponse struct {
}

func (c *Client) DeleteClusterSdnIpam(Ipam string) (*DeleteClusterSdnIpamResponse, error) {

}

type GetClusterSdnIpamResponse struct {
}

func (c *Client) GetClusterSdnIpam(Ipam string) (*GetClusterSdnIpamResponse, error) {

}

type PutClusterSdnIpamParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest  *string `json:"digest"`
	Section *int    `json:"section"`
	Token   *string `json:"token"`
	Url     *string `json:"url"`
}

type PutClusterSdnIpamResponse struct {
}

func (c *Client) PutClusterSdnIpam(Ipam string, parameters PutClusterSdnIpamParameters) (*PutClusterSdnIpamResponse, error) {

}

type GetClusterSdnDnsParameters struct {
	// Only list sdn dns of specific type
	Type *ClusterSdnDnsType `json:"type"`
}

type GetClusterSdnDnsResponse struct {
}

func (c *Client) GetClusterSdnDns(parameters GetClusterSdnDnsParameters) (*GetClusterSdnDnsResponse, error) {

}

type PostClusterSdnDnsParameters struct {
	// The SDN dns object identifier.
	Dns           string `json:"dns"`
	Key           string `json:"key"`
	Reversemaskv6 *int   `json:"reversemaskv6"`
	Reversev6mask *int   `json:"reversev6mask"`
	Ttl           *int   `json:"ttl"`
	// Plugin type.
	Type ClusterSdnDnsType `json:"type"`
	Url  string            `json:"url"`
}

type PostClusterSdnDnsResponse struct {
}

func (c *Client) PostClusterSdnDns(parameters PostClusterSdnDnsParameters) (*PostClusterSdnDnsResponse, error) {

}

type DeleteClusterSdnDnsByDnsResponse struct {
}

func (c *Client) DeleteClusterSdnDnsByDns(Dns string) (*DeleteClusterSdnDnsByDnsResponse, error) {

}

type GetClusterSdnDnsByDnsResponse struct {
}

func (c *Client) GetClusterSdnDnsByDns(Dns string) (*GetClusterSdnDnsByDnsResponse, error) {

}

type PutClusterSdnDnsByDnsParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest        *string `json:"digest"`
	Key           *string `json:"key"`
	Reversemaskv6 *int    `json:"reversemaskv6"`
	Ttl           *int    `json:"ttl"`
	Url           *string `json:"url"`
}

type PutClusterSdnDnsByDnsResponse struct {
}

func (c *Client) PutClusterSdnDnsByDns(Dns string, parameters PutClusterSdnDnsByDnsParameters) (*PutClusterSdnDnsByDnsResponse, error) {

}

type GetClusterLogParameters struct {
	// Maximum number of entries.
	Max *int `json:"max"`
}

type GetClusterLogResponse struct {
}

func (c *Client) GetClusterLog(parameters GetClusterLogParameters) (*GetClusterLogResponse, error) {

}

type GetClusterResourcesParameters struct {
	Type *ClusterResourcesType `json:"type"`
}

type GetClusterResourcesResponse struct {
}

func (c *Client) GetClusterResources(parameters GetClusterResourcesParameters) (*GetClusterResourcesResponse, error) {

}

type GetClusterTasksResponse struct {
}

func (c *Client) GetClusterTasks() (*GetClusterTasksResponse, error) {

}

type GetClusterOptionsResponse struct {
}

func (c *Client) GetClusterOptions() (*GetClusterOptionsResponse, error) {

}

type PutClusterOptionsParameters struct {
	// Set I/O bandwidth limit for various operations (in KiB/s).
	Bwlimit *string `json:"bwlimit"`
	// Select the default Console viewer. You can either use the builtin java applet (VNC; deprecated and maps to html5), an external virt-viewer comtatible application (SPICE), an HTML5 based vnc viewer (noVNC), or an HTML5 based console client (xtermjs). If the selected viewer is not available (e.g. SPICE not activated for the VM), the fallback is noVNC.
	Console *ClusterOptionsConsole `json:"console"`
	// Cluster resource scheduling settings.
	Crs *string `json:"crs"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Datacenter description. Shown in the web-interface datacenter notes panel. This is saved as comment inside the configuration file.
	Description *string `json:"description"`
	// Specify email address to send notification from (default is root@$hostname)
	EmailFrom *string `json:"email_from"`
	// Set the fencing mode of the HA cluster. Hardware mode needs a valid configuration of fence devices in /etc/pve/ha/fence.cfg. With both all two modes are used.
	//
	// WARNING: 'hardware' and 'both' are EXPERIMENTAL & WIP
	Fencing *ClusterOptionsFencing `json:"fencing"`
	// Cluster wide HA settings.
	Ha *string `json:"ha"`
	// Specify external http proxy which is used for downloads (example: 'http://username:password@host:port/')
	HttpProxy *string `json:"http_proxy"`
	// Default keybord layout for vnc server.
	Keyboard *Keyboard `json:"keyboard"`
	// Default GUI language.
	Language *ClusterOptionsLanguage `json:"language"`
	// Prefix for autogenerated MAC addresses.
	MacPrefix *string `json:"mac_prefix"`
	// Defines how many workers (per node) are maximal started  on actions like 'stopall VMs' or task from the ha-manager.
	MaxWorkers *int `json:"max_workers"`
	// For cluster wide migration settings.
	Migration *string `json:"migration"`
	// Migration is secure using SSH tunnel by default. For secure private networks you can disable it to speed up migration. Deprecated, use the 'migration' property instead!
	MigrationUnsecure *bool `json:"migration_unsecure"`
	// Control the range for the free VMID auto-selection pool.
	NextId *string `json:"next-id"`
	// Cluster-wide notification settings.
	Notify *string `json:"notify"`
	// A list of tags that require a `Sys.Modify` on '/' to set and delete. Tags set here that are also in 'user-tag-access' also require `Sys.Modify`.
	RegisteredTags *string `json:"registered-tags"`
	// Tag style options.
	TagStyle *string `json:"tag-style"`
	// u2f
	U2f *string `json:"u2f"`
	// Privilege options for user-settable tags
	UserTagAccess *string `json:"user-tag-access"`
	// webauthn configuration
	Webauthn *string `json:"webauthn"`
}

type PutClusterOptionsResponse struct {
}

func (c *Client) PutClusterOptions(parameters PutClusterOptionsParameters) (*PutClusterOptionsResponse, error) {

}

type GetClusterStatusResponse struct {
}

func (c *Client) GetClusterStatus() (*GetClusterStatusResponse, error) {

}

type GetClusterNextidParameters struct {
	// The (unique) ID of the VM.
	Vmid *int `json:"vmid"`
}

type GetClusterNextidResponse struct {
}

func (c *Client) GetClusterNextid(parameters GetClusterNextidParameters) (*GetClusterNextidResponse, error) {

}

type GetNodesResponse struct {
}

func (c *Client) GetNodes() (*GetNodesResponse, error) {

}

type GetNodeResponse struct {
}

func (c *Client) GetNode(Node string) (*GetNodeResponse, error) {

}

type GetNodeQemuParameters struct {
	// Determine the full status of active VMs.
	Full *bool `json:"full"`
}

type GetNodeQemuResponse struct {
}

func (c *Client) GetNodeQemu(Node string, parameters GetNodeQemuParameters) (*GetNodeQemuResponse, error) {

}

type PostNodeQemuParameters struct {
	// Enable/disable ACPI.
	Acpi *bool `json:"acpi"`
	// List of host cores used to execute guest processes, for example: 0,5,8-11
	Affinity *string `json:"affinity"`
	// Enable/disable communication with the QEMU Guest Agent and its properties.
	Agent *string `json:"agent"`
	// Virtual processor architecture. Defaults to the host.
	Arch *NodeQemuArch `json:"arch"`
	// The backup archive. Either the file system path to a .tar or .vma file (use '-' to pipe data from stdin) or a proxmox storage backup volume identifier.
	Archive *string `json:"archive"`
	// Arbitrary arguments passed to kvm.
	Args *string `json:"args"`
	// Configure a audio device, useful in combination with QXL/Spice.
	Audio0 *string `json:"audio0"`
	// Automatic restart after crash (currently ignored).
	Autostart *bool `json:"autostart"`
	// Amount of target RAM for the VM in MiB. Using zero disables the ballon driver.
	Balloon *int `json:"balloon"`
	// Select BIOS implementation.
	Bios *NodeQemuBios `json:"bios"`
	// Specify guest boot order. Use the 'order=' sub-property as usage with no key or 'legacy=' is deprecated.
	Boot *string `json:"boot"`
	// Enable booting from specified disk. Deprecated: Use 'boot: order=foo;bar' instead.
	Bootdisk *string `json:"bootdisk"`
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *int `json:"bwlimit"`
	// This is an alias for option -ide2
	Cdrom *string `json:"cdrom"`
	// cloud-init: Specify custom files to replace the automatically generated ones at start.
	Cicustom *string `json:"cicustom"`
	// cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords.
	Cipassword *string `json:"cipassword"`
	// Specifies the cloud-init configuration format. The default depends on the configured operating system type (`ostype`. We use the `nocloud` format for Linux, and `configdrive2` for windows.
	Citype *NodeQemuCitype `json:"citype"`
	// cloud-init: do an automatic package upgrade after the first boot.
	Ciupgrade *bool `json:"ciupgrade"`
	// cloud-init: User name to change ssh keys and password for instead of the image's configured default user.
	Ciuser *string `json:"ciuser"`
	// The number of cores per socket.
	Cores *int `json:"cores"`
	// Emulated CPU type.
	Cpu *string `json:"cpu"`
	// Limit of CPU usage.
	Cpulimit *float64 `json:"cpulimit"`
	// CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2.
	Cpuunits *int `json:"cpuunits"`
	// Description for the VM. Shown in the web-interface VM's summary. This is saved as comment inside the configuration file.
	Description *string `json:"description"`
	// Configure a disk for storing EFI vars. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and that the default EFI vars are copied to the volume instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Efidisk0 *string `json:"efidisk0"`
	// Allow to overwrite existing VM.
	Force *bool `json:"force"`
	// Freeze CPU at startup (use 'c' monitor command to start execution).
	Freeze *bool `json:"freeze"`
	// Script that will be executed during various steps in the vms lifetime.
	Hookscript *string `json:"hookscript"`
	// Map host PCI devices into guest.
	HostpciN *string `json:"hostpci[n]"`
	// Selectively enable hotplug features. This is a comma separated list of hotplug features: 'network', 'disk', 'cpu', 'memory', 'usb' and 'cloudinit'. Use '0' to disable hotplug completely. Using '1' as value is an alias for the default `network,disk,usb`. USB hotplugging is possible for guests with machine version >= 7.1 and ostype l26 or windows > 7.
	Hotplug *string `json:"hotplug"`
	// Enable/disable hugepages memory.
	Hugepages *NodeQemuHugepages `json:"hugepages"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	IdeN *string `json:"ide[n]"`
	// cloud-init: Specify IP addresses and gateways for the corresponding interface.
	//
	// IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.
	//
	// The special string 'dhcp' can be used for IP addresses to use DHCP, in which case no explicit
	// gateway should be provided.
	// For IPv6 the special string 'auto' can be used to use stateless autoconfiguration. This requires
	// cloud-init 19.4 or newer.
	//
	// If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using
	// dhcp on IPv4.
	//
	IpconfigN *string `json:"ipconfig[n]"`
	// Inter-VM shared memory. Useful for direct communication between VMs, or to the host.
	Ivshmem *string `json:"ivshmem"`
	// Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts.
	Keephugepages *bool `json:"keephugepages"`
	// Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS.
	Keyboard *Keyboard `json:"keyboard"`
	// Enable/disable KVM hardware virtualization.
	Kvm *bool `json:"kvm"`
	// Start the VM immediately from the backup and restore in background. PBS only.
	LiveRestore *bool `json:"live-restore"`
	// Set the real time clock (RTC) to local time. This is enabled by default if the `ostype` indicates a Microsoft Windows OS.
	Localtime *bool `json:"localtime"`
	// Lock/unlock the VM.
	Lock *NodeQemuLock `json:"lock"`
	// Specifies the QEMU machine type.
	Machine *string `json:"machine"`
	// Amount of RAM for the VM in MiB. This is the maximum available memory when you use the balloon device.
	Memory *int `json:"memory"`
	// Set maximum tolerated downtime (in seconds) for migrations.
	MigrateDowntime *float64 `json:"migrate_downtime"`
	// Set maximum speed (in MB/s) for migrations. Value 0 is no limit.
	MigrateSpeed *int `json:"migrate_speed"`
	// Set a name for the VM. Only used on the configuration web interface.
	Name *string `json:"name"`
	// cloud-init: Sets DNS server IP address for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	Nameserver *string `json:"nameserver"`
	// Specify network devices.
	NetN *string `json:"net[n]"`
	// Enable/disable NUMA.
	Numa *bool `json:"numa"`
	// NUMA topology.
	NumaN *string `json:"numa[n]"`
	// Specifies whether a VM will be started during system bootup.
	Onboot *bool `json:"onboot"`
	// Specify guest operating system.
	Ostype *NodeQemuOstype `json:"ostype"`
	// Map host parallel devices (n is 0 to 2).
	ParallelN *string `json:"parallel[n]"`
	// Add the VM to the specified pool.
	Pool *string `json:"pool"`
	// Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.
	Protection *bool `json:"protection"`
	// Allow reboot. If set to '0' the VM exit on reboot.
	Reboot *bool `json:"reboot"`
	// Configure a VirtIO-based Random Number Generator.
	Rng0 *string `json:"rng0"`
	// Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	SataN *string `json:"sata[n]"`
	// Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	ScsiN *string `json:"scsi[n]"`
	// SCSI controller model
	Scsihw *NodeQemuScsihw `json:"scsihw"`
	// cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	Searchdomain *string `json:"searchdomain"`
	// Create a serial device inside the VM (n is 0 to 3)
	SerialN *string `json:"serial[n]"`
	// Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd.
	Shares *int `json:"shares"`
	// Specify SMBIOS type 1 fields.
	Smbios1 *string `json:"smbios1"`
	// The number of CPUs. Please use option -sockets instead.
	Smp *int `json:"smp"`
	// The number of CPU sockets.
	Sockets *int `json:"sockets"`
	// Configure additional enhancements for SPICE.
	SpiceEnhancements *string `json:"spice_enhancements"`
	// cloud-init: Setup public SSH keys (one key per line, OpenSSH format).
	Sshkeys *string `json:"sshkeys"`
	// Start VM after it was created successfully.
	Start *bool `json:"start"`
	// Set the initial date of the real time clock. Valid format for date are:'now' or '2006-06-17T16:01:21' or '2006-06-17'.
	Startdate *string `json:"startdate"`
	// Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Startup *string `json:"startup"`
	// Default storage.
	Storage *string `json:"storage"`
	// Enable/disable the USB tablet device.
	Tablet *bool `json:"tablet"`
	// Tags of the VM. This is only meta information.
	Tags *string `json:"tags"`
	// Enable/disable time drift fix.
	Tdf *bool `json:"tdf"`
	// Enable/disable Template.
	Template *bool `json:"template"`
	// Configure a Disk for storing TPM state. The format is fixed to 'raw'. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and 4 MiB will be used instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Tpmstate0 *string `json:"tpmstate0"`
	// Assign a unique random ethernet address.
	Unique *bool `json:"unique"`
	// Reference to unused volumes. This is used internally, and should not be modified manually.
	UnusedN *string `json:"unused[n]"`
	// Configure an USB device (n is 0 to 4, for machine version >= 7.1 and ostype l26 or windows > 7, n can be up to 14).
	UsbN *string `json:"usb[n]"`
	// Number of hotplugged vcpus.
	Vcpus *int `json:"vcpus"`
	// Configure the VGA hardware.
	Vga *string `json:"vga"`
	// Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	VirtioN *string `json:"virtio[n]"`
	// Set VM Generation ID. Use '1' to autogenerate on create or update, pass '0' to disable explicitly.
	Vmgenid *string `json:"vmgenid"`
	// The (unique) ID of the VM.
	Vmid int `json:"vmid"`
	// Default storage for VM state volumes/files.
	Vmstatestorage *string `json:"vmstatestorage"`
	// Create a virtual hardware watchdog device.
	Watchdog *string `json:"watchdog"`
}

type PostNodeQemuResponse struct {
}

func (c *Client) PostNodeQemu(Node string, parameters PostNodeQemuParameters) (*PostNodeQemuResponse, error) {

}

type DeleteNodeQemuByVmidParameters struct {
	// If set, destroy additionally all disks not referenced in the config but with a matching VMID from all enabled storages.
	DestroyUnreferencedDisks *bool `json:"destroy-unreferenced-disks"`
	// Remove VMID from configurations, like backup & replication jobs and HA.
	Purge *bool `json:"purge"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
}

type DeleteNodeQemuByVmidResponse struct {
}

func (c *Client) DeleteNodeQemuByVmid(Node string, Vmid int, parameters DeleteNodeQemuByVmidParameters) (*DeleteNodeQemuByVmidResponse, error) {

}

type GetNodeQemuByVmidResponse struct {
}

func (c *Client) GetNodeQemuByVmid(Node string, Vmid int) (*GetNodeQemuByVmidResponse, error) {

}

type GetNodeQemuByVmidFirewallResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewall(Node string, Vmid int) (*GetNodeQemuByVmidFirewallResponse, error) {

}

type GetNodeQemuByVmidFirewallRulesResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallRules(Node string, Vmid int) (*GetNodeQemuByVmidFirewallRulesResponse, error) {

}

type PostNodeQemuByVmidFirewallRulesParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Update rule at position <pos>.
	Pos *int `json:"pos"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type FirewallType `json:"type"`
}

type PostNodeQemuByVmidFirewallRulesResponse struct {
}

func (c *Client) PostNodeQemuByVmidFirewallRules(Node string, Vmid int, parameters PostNodeQemuByVmidFirewallRulesParameters) (*PostNodeQemuByVmidFirewallRulesResponse, error) {

}

type DeleteNodeQemuByVmidFirewallRuleParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteNodeQemuByVmidFirewallRuleResponse struct {
}

func (c *Client) DeleteNodeQemuByVmidFirewallRule(Node string, Vmid int, Pos int, parameters DeleteNodeQemuByVmidFirewallRuleParameters) (*DeleteNodeQemuByVmidFirewallRuleResponse, error) {

}

type GetNodeQemuByVmidFirewallRuleResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallRule(Node string, Vmid int, Pos int) (*GetNodeQemuByVmidFirewallRuleResponse, error) {

}

type PutNodeQemuByVmidFirewallRuleParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action *string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Move rule to new position <moveto>. Other arguments are ignored.
	Moveto *int `json:"moveto"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type *FirewallType `json:"type"`
}

type PutNodeQemuByVmidFirewallRuleResponse struct {
}

func (c *Client) PutNodeQemuByVmidFirewallRule(Node string, Vmid int, Pos int, parameters PutNodeQemuByVmidFirewallRuleParameters) (*PutNodeQemuByVmidFirewallRuleResponse, error) {

}

type GetNodeQemuByVmidFirewallAliasesResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallAliases(Node string, Vmid int) (*GetNodeQemuByVmidFirewallAliasesResponse, error) {

}

type PostNodeQemuByVmidFirewallAliasesParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	// Alias name.
	Name string `json:"name"`
}

type PostNodeQemuByVmidFirewallAliasesResponse struct {
}

func (c *Client) PostNodeQemuByVmidFirewallAliases(Node string, Vmid int, parameters PostNodeQemuByVmidFirewallAliasesParameters) (*PostNodeQemuByVmidFirewallAliasesResponse, error) {

}

type DeleteNodeQemuByVmidFirewallAliasParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteNodeQemuByVmidFirewallAliasResponse struct {
}

func (c *Client) DeleteNodeQemuByVmidFirewallAlias(Node string, Vmid int, Name string, parameters DeleteNodeQemuByVmidFirewallAliasParameters) (*DeleteNodeQemuByVmidFirewallAliasResponse, error) {

}

type GetNodeQemuByVmidFirewallAliasResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallAlias(Node string, Vmid int, Name string) (*GetNodeQemuByVmidFirewallAliasResponse, error) {

}

type PutNodeQemuByVmidFirewallAliasParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Rename an existing alias.
	Rename *string `json:"rename"`
}

type PutNodeQemuByVmidFirewallAliasResponse struct {
}

func (c *Client) PutNodeQemuByVmidFirewallAlias(Node string, Vmid int, Name string, parameters PutNodeQemuByVmidFirewallAliasParameters) (*PutNodeQemuByVmidFirewallAliasResponse, error) {

}

type GetNodeQemuByVmidFirewallIpsetResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallIpset(Node string, Vmid int) (*GetNodeQemuByVmidFirewallIpsetResponse, error) {

}

type PostNodeQemuByVmidFirewallIpsetParameters struct {
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// IP set name.
	Name string `json:"name"`
	// Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
	Rename *string `json:"rename"`
}

type PostNodeQemuByVmidFirewallIpsetResponse struct {
}

func (c *Client) PostNodeQemuByVmidFirewallIpset(Node string, Vmid int, parameters PostNodeQemuByVmidFirewallIpsetParameters) (*PostNodeQemuByVmidFirewallIpsetResponse, error) {

}

type DeleteNodeQemuByVmidFirewallIpsetByNameParameters struct {
	// Delete all members of the IPSet, if there are any.
	Force *bool `json:"force"`
}

type DeleteNodeQemuByVmidFirewallIpsetByNameResponse struct {
}

func (c *Client) DeleteNodeQemuByVmidFirewallIpsetByName(Node string, Vmid int, Name string, parameters DeleteNodeQemuByVmidFirewallIpsetByNameParameters) (*DeleteNodeQemuByVmidFirewallIpsetByNameResponse, error) {

}

type GetNodeQemuByVmidFirewallIpsetByNameResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallIpsetByName(Node string, Vmid int, Name string) (*GetNodeQemuByVmidFirewallIpsetByNameResponse, error) {

}

type PostNodeQemuByVmidFirewallIpsetByNameParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	Nomatch *bool   `json:"nomatch"`
}

type PostNodeQemuByVmidFirewallIpsetByNameResponse struct {
}

func (c *Client) PostNodeQemuByVmidFirewallIpsetByName(Node string, Vmid int, Name string, parameters PostNodeQemuByVmidFirewallIpsetByNameParameters) (*PostNodeQemuByVmidFirewallIpsetByNameResponse, error) {

}

type DeleteNodeQemuByVmidFirewallIpsetByNameByCidrParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteNodeQemuByVmidFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) DeleteNodeQemuByVmidFirewallIpsetByNameByCidr(Node string, Vmid int, Name string, Cidr string, parameters DeleteNodeQemuByVmidFirewallIpsetByNameByCidrParameters) (*DeleteNodeQemuByVmidFirewallIpsetByNameByCidrResponse, error) {

}

type GetNodeQemuByVmidFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallIpsetByNameByCidr(Node string, Vmid int, Name string, Cidr string) (*GetNodeQemuByVmidFirewallIpsetByNameByCidrResponse, error) {

}

type PutNodeQemuByVmidFirewallIpsetByNameByCidrParameters struct {
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest  *string `json:"digest"`
	Nomatch *bool   `json:"nomatch"`
}

type PutNodeQemuByVmidFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) PutNodeQemuByVmidFirewallIpsetByNameByCidr(Node string, Vmid int, Name string, Cidr string, parameters PutNodeQemuByVmidFirewallIpsetByNameByCidrParameters) (*PutNodeQemuByVmidFirewallIpsetByNameByCidrResponse, error) {

}

type GetNodeQemuByVmidFirewallOptionsResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallOptions(Node string, Vmid int) (*GetNodeQemuByVmidFirewallOptionsResponse, error) {

}

type PutNodeQemuByVmidFirewallOptionsParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Enable DHCP.
	Dhcp *bool `json:"dhcp"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Enable/disable firewall rules.
	Enable *bool `json:"enable"`
	// Enable default IP filters. This is equivalent to adding an empty ipfilter-net<id> ipset for every interface. Such ipsets implicitly contain sane default restrictions such as restricting IPv6 link local addresses to the one derived from the interface's MAC address. For containers the configured IP addresses will be implicitly added.
	Ipfilter *bool `json:"ipfilter"`
	// Log level for incoming traffic.
	LogLevelIn *FirewallLog `json:"log_level_in"`
	// Log level for outgoing traffic.
	LogLevelOut *FirewallLog `json:"log_level_out"`
	// Enable/disable MAC address filter.
	Macfilter *bool `json:"macfilter"`
	// Enable NDP (Neighbor Discovery Protocol).
	Ndp *bool `json:"ndp"`
	// Input policy.
	PolicyIn *FirewallOptionsPolicy `json:"policy_in"`
	// Output policy.
	PolicyOut *FirewallOptionsPolicy `json:"policy_out"`
	// Allow sending Router Advertisement.
	Radv *bool `json:"radv"`
}

type PutNodeQemuByVmidFirewallOptionsResponse struct {
}

func (c *Client) PutNodeQemuByVmidFirewallOptions(Node string, Vmid int, parameters PutNodeQemuByVmidFirewallOptionsParameters) (*PutNodeQemuByVmidFirewallOptionsResponse, error) {

}

type GetNodeQemuByVmidFirewallLogParameters struct {
	Limit *int `json:"limit"`
	// Display log since this UNIX epoch.
	Since *int `json:"since"`
	Start *int `json:"start"`
	// Display log until this UNIX epoch.
	Until *int `json:"until"`
}

type GetNodeQemuByVmidFirewallLogResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallLog(Node string, Vmid int, parameters GetNodeQemuByVmidFirewallLogParameters) (*GetNodeQemuByVmidFirewallLogResponse, error) {

}

type GetNodeQemuByVmidFirewallRefsParameters struct {
	// Only list references of specified type.
	Type *FirewallRefsType `json:"type"`
}

type GetNodeQemuByVmidFirewallRefsResponse struct {
}

func (c *Client) GetNodeQemuByVmidFirewallRefs(Node string, Vmid int, parameters GetNodeQemuByVmidFirewallRefsParameters) (*GetNodeQemuByVmidFirewallRefsResponse, error) {

}

type GetNodeQemuByVmidAgentResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgent(Node string, Vmid int) (*GetNodeQemuByVmidAgentResponse, error) {

}

type PostNodeQemuByVmidAgentParameters struct {
	// The QGA command.
	Command NodeQemuAgentCommand `json:"command"`
}

type PostNodeQemuByVmidAgentResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgent(Node string, Vmid int, parameters PostNodeQemuByVmidAgentParameters) (*PostNodeQemuByVmidAgentResponse, error) {

}

type PostNodeQemuByVmidAgentFsfreezeFreezeResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentFsfreezeFreeze(Node string, Vmid int) (*PostNodeQemuByVmidAgentFsfreezeFreezeResponse, error) {

}

type PostNodeQemuByVmidAgentFsfreezeStatusResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentFsfreezeStatus(Node string, Vmid int) (*PostNodeQemuByVmidAgentFsfreezeStatusResponse, error) {

}

type PostNodeQemuByVmidAgentFsfreezeThawResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentFsfreezeThaw(Node string, Vmid int) (*PostNodeQemuByVmidAgentFsfreezeThawResponse, error) {

}

type PostNodeQemuByVmidAgentFstrimResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentFstrim(Node string, Vmid int) (*PostNodeQemuByVmidAgentFstrimResponse, error) {

}

type GetNodeQemuByVmidAgentGetFsinfoResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetFsinfo(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetFsinfoResponse, error) {

}

type GetNodeQemuByVmidAgentGetHostNameResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetHostName(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetHostNameResponse, error) {

}

type GetNodeQemuByVmidAgentGetMemoryBlockInfoResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetMemoryBlockInfo(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetMemoryBlockInfoResponse, error) {

}

type GetNodeQemuByVmidAgentGetMemoryBlocksResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetMemoryBlocks(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetMemoryBlocksResponse, error) {

}

type GetNodeQemuByVmidAgentGetOsinfoResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetOsinfo(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetOsinfoResponse, error) {

}

type GetNodeQemuByVmidAgentGetTimeResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetTime(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetTimeResponse, error) {

}

type GetNodeQemuByVmidAgentGetTimezoneResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetTimezone(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetTimezoneResponse, error) {

}

type GetNodeQemuByVmidAgentGetUsersResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetUsers(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetUsersResponse, error) {

}

type GetNodeQemuByVmidAgentGetVcpusResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentGetVcpus(Node string, Vmid int) (*GetNodeQemuByVmidAgentGetVcpusResponse, error) {

}

type GetNodeQemuByVmidAgentInfoResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentInfo(Node string, Vmid int) (*GetNodeQemuByVmidAgentInfoResponse, error) {

}

type GetNodeQemuByVmidAgentNetworkGetInterfacesResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentNetworkGetInterfaces(Node string, Vmid int) (*GetNodeQemuByVmidAgentNetworkGetInterfacesResponse, error) {

}

type PostNodeQemuByVmidAgentPingResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentPing(Node string, Vmid int) (*PostNodeQemuByVmidAgentPingResponse, error) {

}

type PostNodeQemuByVmidAgentShutdownResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentShutdown(Node string, Vmid int) (*PostNodeQemuByVmidAgentShutdownResponse, error) {

}

type PostNodeQemuByVmidAgentSuspendDiskResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentSuspendDisk(Node string, Vmid int) (*PostNodeQemuByVmidAgentSuspendDiskResponse, error) {

}

type PostNodeQemuByVmidAgentSuspendHybridResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentSuspendHybrid(Node string, Vmid int) (*PostNodeQemuByVmidAgentSuspendHybridResponse, error) {

}

type PostNodeQemuByVmidAgentSuspendRamResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentSuspendRam(Node string, Vmid int) (*PostNodeQemuByVmidAgentSuspendRamResponse, error) {

}

type PostNodeQemuByVmidAgentSetUserPasswordParameters struct {
	// set to 1 if the password has already been passed through crypt()
	Crypted *bool `json:"crypted"`
	// The new password.
	Password string `json:"password"`
	// The user to set the password for.
	Username string `json:"username"`
}

type PostNodeQemuByVmidAgentSetUserPasswordResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentSetUserPassword(Node string, Vmid int, parameters PostNodeQemuByVmidAgentSetUserPasswordParameters) (*PostNodeQemuByVmidAgentSetUserPasswordResponse, error) {

}

type PostNodeQemuByVmidAgentExecParameters struct {
	// The command as a list of program + arguments.
	Command any `json:"command"`
	// Data to pass as 'input-data' to the guest. Usually treated as STDIN to 'command'.
	InputData *string `json:"input-data"`
}

type PostNodeQemuByVmidAgentExecResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentExec(Node string, Vmid int, parameters PostNodeQemuByVmidAgentExecParameters) (*PostNodeQemuByVmidAgentExecResponse, error) {

}

type GetNodeQemuByVmidAgentExecStatusParameters struct {
	// The PID to query
	Pid int `json:"pid"`
}

type GetNodeQemuByVmidAgentExecStatusResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentExecStatus(Node string, Vmid int, parameters GetNodeQemuByVmidAgentExecStatusParameters) (*GetNodeQemuByVmidAgentExecStatusResponse, error) {

}

type GetNodeQemuByVmidAgentFileReadParameters struct {
	// The path to the file
	File string `json:"file"`
}

type GetNodeQemuByVmidAgentFileReadResponse struct {
}

func (c *Client) GetNodeQemuByVmidAgentFileRead(Node string, Vmid int, parameters GetNodeQemuByVmidAgentFileReadParameters) (*GetNodeQemuByVmidAgentFileReadResponse, error) {

}

type PostNodeQemuByVmidAgentFileWriteParameters struct {
	// The content to write into the file.
	Content string `json:"content"`
	// If set, the content will be encoded as base64 (required by QEMU).Otherwise the content needs to be encoded beforehand - defaults to true.
	Encode *bool `json:"encode"`
	// The path to the file.
	File string `json:"file"`
}

type PostNodeQemuByVmidAgentFileWriteResponse struct {
}

func (c *Client) PostNodeQemuByVmidAgentFileWrite(Node string, Vmid int, parameters PostNodeQemuByVmidAgentFileWriteParameters) (*PostNodeQemuByVmidAgentFileWriteResponse, error) {

}

type GetNodeQemuByVmidRrdParameters struct {
	// The RRD consolidation function
	Cf *NodeRrdCf `json:"cf"`
	// The list of datasources you want to display.
	Ds string `json:"ds"`
	// Specify the time frame you are interested in.
	Timeframe NodeRrdTimeframe `json:"timeframe"`
}

type GetNodeQemuByVmidRrdResponse struct {
}

func (c *Client) GetNodeQemuByVmidRrd(Node string, Vmid int, parameters GetNodeQemuByVmidRrdParameters) (*GetNodeQemuByVmidRrdResponse, error) {

}

type GetNodeQemuByVmidRrddataParameters struct {
	// The RRD consolidation function
	Cf *NodeRrdCf `json:"cf"`
	// Specify the time frame you are interested in.
	Timeframe NodeRrdTimeframe `json:"timeframe"`
}

type GetNodeQemuByVmidRrddataResponse struct {
}

func (c *Client) GetNodeQemuByVmidRrddata(Node string, Vmid int, parameters GetNodeQemuByVmidRrddataParameters) (*GetNodeQemuByVmidRrddataResponse, error) {

}

type GetNodeQemuByVmidConfigParameters struct {
	// Get current values (instead of pending values).
	Current *bool `json:"current"`
	// Fetch config values from given snapshot.
	Snapshot *string `json:"snapshot"`
}

type GetNodeQemuByVmidConfigResponse struct {
}

func (c *Client) GetNodeQemuByVmidConfig(Node string, Vmid int, parameters GetNodeQemuByVmidConfigParameters) (*GetNodeQemuByVmidConfigResponse, error) {

}

type PostNodeQemuByVmidConfigParameters struct {
	// Enable/disable ACPI.
	Acpi *bool `json:"acpi"`
	// List of host cores used to execute guest processes, for example: 0,5,8-11
	Affinity *string `json:"affinity"`
	// Enable/disable communication with the QEMU Guest Agent and its properties.
	Agent *string `json:"agent"`
	// Virtual processor architecture. Defaults to the host.
	Arch *NodeQemuArch `json:"arch"`
	// Arbitrary arguments passed to kvm.
	Args *string `json:"args"`
	// Configure a audio device, useful in combination with QXL/Spice.
	Audio0 *string `json:"audio0"`
	// Automatic restart after crash (currently ignored).
	Autostart *bool `json:"autostart"`
	// Time to wait for the task to finish. We return 'null' if the task finish within that time.
	BackgroundDelay *int `json:"background_delay"`
	// Amount of target RAM for the VM in MiB. Using zero disables the ballon driver.
	Balloon *int `json:"balloon"`
	// Select BIOS implementation.
	Bios *NodeQemuBios `json:"bios"`
	// Specify guest boot order. Use the 'order=' sub-property as usage with no key or 'legacy=' is deprecated.
	Boot *string `json:"boot"`
	// Enable booting from specified disk. Deprecated: Use 'boot: order=foo;bar' instead.
	Bootdisk *string `json:"bootdisk"`
	// This is an alias for option -ide2
	Cdrom *string `json:"cdrom"`
	// cloud-init: Specify custom files to replace the automatically generated ones at start.
	Cicustom *string `json:"cicustom"`
	// cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords.
	Cipassword *string `json:"cipassword"`
	// Specifies the cloud-init configuration format. The default depends on the configured operating system type (`ostype`. We use the `nocloud` format for Linux, and `configdrive2` for windows.
	Citype *NodeQemuCitype `json:"citype"`
	// cloud-init: do an automatic package upgrade after the first boot.
	Ciupgrade *bool `json:"ciupgrade"`
	// cloud-init: User name to change ssh keys and password for instead of the image's configured default user.
	Ciuser *string `json:"ciuser"`
	// The number of cores per socket.
	Cores *int `json:"cores"`
	// Emulated CPU type.
	Cpu *string `json:"cpu"`
	// Limit of CPU usage.
	Cpulimit *float64 `json:"cpulimit"`
	// CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2.
	Cpuunits *int `json:"cpuunits"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Description for the VM. Shown in the web-interface VM's summary. This is saved as comment inside the configuration file.
	Description *string `json:"description"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Configure a disk for storing EFI vars. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and that the default EFI vars are copied to the volume instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Efidisk0 *string `json:"efidisk0"`
	// Force physical removal. Without this, we simple remove the disk from the config file and create an additional configuration entry called 'unused[n]', which contains the volume ID. Unlink of unused[n] always cause physical removal.
	Force *bool `json:"force"`
	// Freeze CPU at startup (use 'c' monitor command to start execution).
	Freeze *bool `json:"freeze"`
	// Script that will be executed during various steps in the vms lifetime.
	Hookscript *string `json:"hookscript"`
	// Map host PCI devices into guest.
	HostpciN *string `json:"hostpci[n]"`
	// Selectively enable hotplug features. This is a comma separated list of hotplug features: 'network', 'disk', 'cpu', 'memory', 'usb' and 'cloudinit'. Use '0' to disable hotplug completely. Using '1' as value is an alias for the default `network,disk,usb`. USB hotplugging is possible for guests with machine version >= 7.1 and ostype l26 or windows > 7.
	Hotplug *string `json:"hotplug"`
	// Enable/disable hugepages memory.
	Hugepages *NodeQemuHugepages `json:"hugepages"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	IdeN *string `json:"ide[n]"`
	// cloud-init: Specify IP addresses and gateways for the corresponding interface.
	//
	// IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.
	//
	// The special string 'dhcp' can be used for IP addresses to use DHCP, in which case no explicit
	// gateway should be provided.
	// For IPv6 the special string 'auto' can be used to use stateless autoconfiguration. This requires
	// cloud-init 19.4 or newer.
	//
	// If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using
	// dhcp on IPv4.
	//
	IpconfigN *string `json:"ipconfig[n]"`
	// Inter-VM shared memory. Useful for direct communication between VMs, or to the host.
	Ivshmem *string `json:"ivshmem"`
	// Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts.
	Keephugepages *bool `json:"keephugepages"`
	// Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS.
	Keyboard *Keyboard `json:"keyboard"`
	// Enable/disable KVM hardware virtualization.
	Kvm *bool `json:"kvm"`
	// Set the real time clock (RTC) to local time. This is enabled by default if the `ostype` indicates a Microsoft Windows OS.
	Localtime *bool `json:"localtime"`
	// Lock/unlock the VM.
	Lock *NodeQemuLock `json:"lock"`
	// Specifies the QEMU machine type.
	Machine *string `json:"machine"`
	// Amount of RAM for the VM in MiB. This is the maximum available memory when you use the balloon device.
	Memory *int `json:"memory"`
	// Set maximum tolerated downtime (in seconds) for migrations.
	MigrateDowntime *float64 `json:"migrate_downtime"`
	// Set maximum speed (in MB/s) for migrations. Value 0 is no limit.
	MigrateSpeed *int `json:"migrate_speed"`
	// Set a name for the VM. Only used on the configuration web interface.
	Name *string `json:"name"`
	// cloud-init: Sets DNS server IP address for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	Nameserver *string `json:"nameserver"`
	// Specify network devices.
	NetN *string `json:"net[n]"`
	// Enable/disable NUMA.
	Numa *bool `json:"numa"`
	// NUMA topology.
	NumaN *string `json:"numa[n]"`
	// Specifies whether a VM will be started during system bootup.
	Onboot *bool `json:"onboot"`
	// Specify guest operating system.
	Ostype *NodeQemuOstype `json:"ostype"`
	// Map host parallel devices (n is 0 to 2).
	ParallelN *string `json:"parallel[n]"`
	// Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.
	Protection *bool `json:"protection"`
	// Allow reboot. If set to '0' the VM exit on reboot.
	Reboot *bool `json:"reboot"`
	// Revert a pending change.
	Revert *string `json:"revert"`
	// Configure a VirtIO-based Random Number Generator.
	Rng0 *string `json:"rng0"`
	// Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	SataN *string `json:"sata[n]"`
	// Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	ScsiN *string `json:"scsi[n]"`
	// SCSI controller model
	Scsihw *NodeQemuScsihw `json:"scsihw"`
	// cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	Searchdomain *string `json:"searchdomain"`
	// Create a serial device inside the VM (n is 0 to 3)
	SerialN *string `json:"serial[n]"`
	// Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd.
	Shares *int `json:"shares"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
	// Specify SMBIOS type 1 fields.
	Smbios1 *string `json:"smbios1"`
	// The number of CPUs. Please use option -sockets instead.
	Smp *int `json:"smp"`
	// The number of CPU sockets.
	Sockets *int `json:"sockets"`
	// Configure additional enhancements for SPICE.
	SpiceEnhancements *string `json:"spice_enhancements"`
	// cloud-init: Setup public SSH keys (one key per line, OpenSSH format).
	Sshkeys *string `json:"sshkeys"`
	// Set the initial date of the real time clock. Valid format for date are:'now' or '2006-06-17T16:01:21' or '2006-06-17'.
	Startdate *string `json:"startdate"`
	// Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Startup *string `json:"startup"`
	// Enable/disable the USB tablet device.
	Tablet *bool `json:"tablet"`
	// Tags of the VM. This is only meta information.
	Tags *string `json:"tags"`
	// Enable/disable time drift fix.
	Tdf *bool `json:"tdf"`
	// Enable/disable Template.
	Template *bool `json:"template"`
	// Configure a Disk for storing TPM state. The format is fixed to 'raw'. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and 4 MiB will be used instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Tpmstate0 *string `json:"tpmstate0"`
	// Reference to unused volumes. This is used internally, and should not be modified manually.
	UnusedN *string `json:"unused[n]"`
	// Configure an USB device (n is 0 to 4, for machine version >= 7.1 and ostype l26 or windows > 7, n can be up to 14).
	UsbN *string `json:"usb[n]"`
	// Number of hotplugged vcpus.
	Vcpus *int `json:"vcpus"`
	// Configure the VGA hardware.
	Vga *string `json:"vga"`
	// Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	VirtioN *string `json:"virtio[n]"`
	// Set VM Generation ID. Use '1' to autogenerate on create or update, pass '0' to disable explicitly.
	Vmgenid *string `json:"vmgenid"`
	// Default storage for VM state volumes/files.
	Vmstatestorage *string `json:"vmstatestorage"`
	// Create a virtual hardware watchdog device.
	Watchdog *string `json:"watchdog"`
}

type PostNodeQemuByVmidConfigResponse struct {
}

func (c *Client) PostNodeQemuByVmidConfig(Node string, Vmid int, parameters PostNodeQemuByVmidConfigParameters) (*PostNodeQemuByVmidConfigResponse, error) {

}

type PutNodeQemuByVmidConfigParameters struct {
	// Enable/disable ACPI.
	Acpi *bool `json:"acpi"`
	// List of host cores used to execute guest processes, for example: 0,5,8-11
	Affinity *string `json:"affinity"`
	// Enable/disable communication with the QEMU Guest Agent and its properties.
	Agent *string `json:"agent"`
	// Virtual processor architecture. Defaults to the host.
	Arch *NodeQemuArch `json:"arch"`
	// Arbitrary arguments passed to kvm.
	Args *string `json:"args"`
	// Configure a audio device, useful in combination with QXL/Spice.
	Audio0 *string `json:"audio0"`
	// Automatic restart after crash (currently ignored).
	Autostart *bool `json:"autostart"`
	// Amount of target RAM for the VM in MiB. Using zero disables the ballon driver.
	Balloon *int `json:"balloon"`
	// Select BIOS implementation.
	Bios *NodeQemuBios `json:"bios"`
	// Specify guest boot order. Use the 'order=' sub-property as usage with no key or 'legacy=' is deprecated.
	Boot *string `json:"boot"`
	// Enable booting from specified disk. Deprecated: Use 'boot: order=foo;bar' instead.
	Bootdisk *string `json:"bootdisk"`
	// This is an alias for option -ide2
	Cdrom *string `json:"cdrom"`
	// cloud-init: Specify custom files to replace the automatically generated ones at start.
	Cicustom *string `json:"cicustom"`
	// cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords.
	Cipassword *string `json:"cipassword"`
	// Specifies the cloud-init configuration format. The default depends on the configured operating system type (`ostype`. We use the `nocloud` format for Linux, and `configdrive2` for windows.
	Citype *NodeQemuCitype `json:"citype"`
	// cloud-init: do an automatic package upgrade after the first boot.
	Ciupgrade *bool `json:"ciupgrade"`
	// cloud-init: User name to change ssh keys and password for instead of the image's configured default user.
	Ciuser *string `json:"ciuser"`
	// The number of cores per socket.
	Cores *int `json:"cores"`
	// Emulated CPU type.
	Cpu *string `json:"cpu"`
	// Limit of CPU usage.
	Cpulimit *float64 `json:"cpulimit"`
	// CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2.
	Cpuunits *int `json:"cpuunits"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Description for the VM. Shown in the web-interface VM's summary. This is saved as comment inside the configuration file.
	Description *string `json:"description"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Configure a disk for storing EFI vars. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and that the default EFI vars are copied to the volume instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Efidisk0 *string `json:"efidisk0"`
	// Force physical removal. Without this, we simple remove the disk from the config file and create an additional configuration entry called 'unused[n]', which contains the volume ID. Unlink of unused[n] always cause physical removal.
	Force *bool `json:"force"`
	// Freeze CPU at startup (use 'c' monitor command to start execution).
	Freeze *bool `json:"freeze"`
	// Script that will be executed during various steps in the vms lifetime.
	Hookscript *string `json:"hookscript"`
	// Map host PCI devices into guest.
	HostpciN *string `json:"hostpci[n]"`
	// Selectively enable hotplug features. This is a comma separated list of hotplug features: 'network', 'disk', 'cpu', 'memory', 'usb' and 'cloudinit'. Use '0' to disable hotplug completely. Using '1' as value is an alias for the default `network,disk,usb`. USB hotplugging is possible for guests with machine version >= 7.1 and ostype l26 or windows > 7.
	Hotplug *string `json:"hotplug"`
	// Enable/disable hugepages memory.
	Hugepages *NodeQemuHugepages `json:"hugepages"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	IdeN *string `json:"ide[n]"`
	// cloud-init: Specify IP addresses and gateways for the corresponding interface.
	//
	// IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.
	//
	// The special string 'dhcp' can be used for IP addresses to use DHCP, in which case no explicit
	// gateway should be provided.
	// For IPv6 the special string 'auto' can be used to use stateless autoconfiguration. This requires
	// cloud-init 19.4 or newer.
	//
	// If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using
	// dhcp on IPv4.
	//
	IpconfigN *string `json:"ipconfig[n]"`
	// Inter-VM shared memory. Useful for direct communication between VMs, or to the host.
	Ivshmem *string `json:"ivshmem"`
	// Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts.
	Keephugepages *bool `json:"keephugepages"`
	// Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS.
	Keyboard *Keyboard `json:"keyboard"`
	// Enable/disable KVM hardware virtualization.
	Kvm *bool `json:"kvm"`
	// Set the real time clock (RTC) to local time. This is enabled by default if the `ostype` indicates a Microsoft Windows OS.
	Localtime *bool `json:"localtime"`
	// Lock/unlock the VM.
	Lock *NodeQemuLock `json:"lock"`
	// Specifies the QEMU machine type.
	Machine *string `json:"machine"`
	// Amount of RAM for the VM in MiB. This is the maximum available memory when you use the balloon device.
	Memory *int `json:"memory"`
	// Set maximum tolerated downtime (in seconds) for migrations.
	MigrateDowntime *float64 `json:"migrate_downtime"`
	// Set maximum speed (in MB/s) for migrations. Value 0 is no limit.
	MigrateSpeed *int `json:"migrate_speed"`
	// Set a name for the VM. Only used on the configuration web interface.
	Name *string `json:"name"`
	// cloud-init: Sets DNS server IP address for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	Nameserver *string `json:"nameserver"`
	// Specify network devices.
	NetN *string `json:"net[n]"`
	// Enable/disable NUMA.
	Numa *bool `json:"numa"`
	// NUMA topology.
	NumaN *string `json:"numa[n]"`
	// Specifies whether a VM will be started during system bootup.
	Onboot *bool `json:"onboot"`
	// Specify guest operating system.
	Ostype *NodeQemuOstype `json:"ostype"`
	// Map host parallel devices (n is 0 to 2).
	ParallelN *string `json:"parallel[n]"`
	// Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.
	Protection *bool `json:"protection"`
	// Allow reboot. If set to '0' the VM exit on reboot.
	Reboot *bool `json:"reboot"`
	// Revert a pending change.
	Revert *string `json:"revert"`
	// Configure a VirtIO-based Random Number Generator.
	Rng0 *string `json:"rng0"`
	// Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	SataN *string `json:"sata[n]"`
	// Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	ScsiN *string `json:"scsi[n]"`
	// SCSI controller model
	Scsihw *NodeQemuScsihw `json:"scsihw"`
	// cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	Searchdomain *string `json:"searchdomain"`
	// Create a serial device inside the VM (n is 0 to 3)
	SerialN *string `json:"serial[n]"`
	// Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd.
	Shares *int `json:"shares"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
	// Specify SMBIOS type 1 fields.
	Smbios1 *string `json:"smbios1"`
	// The number of CPUs. Please use option -sockets instead.
	Smp *int `json:"smp"`
	// The number of CPU sockets.
	Sockets *int `json:"sockets"`
	// Configure additional enhancements for SPICE.
	SpiceEnhancements *string `json:"spice_enhancements"`
	// cloud-init: Setup public SSH keys (one key per line, OpenSSH format).
	Sshkeys *string `json:"sshkeys"`
	// Set the initial date of the real time clock. Valid format for date are:'now' or '2006-06-17T16:01:21' or '2006-06-17'.
	Startdate *string `json:"startdate"`
	// Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Startup *string `json:"startup"`
	// Enable/disable the USB tablet device.
	Tablet *bool `json:"tablet"`
	// Tags of the VM. This is only meta information.
	Tags *string `json:"tags"`
	// Enable/disable time drift fix.
	Tdf *bool `json:"tdf"`
	// Enable/disable Template.
	Template *bool `json:"template"`
	// Configure a Disk for storing TPM state. The format is fixed to 'raw'. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and 4 MiB will be used instead. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Tpmstate0 *string `json:"tpmstate0"`
	// Reference to unused volumes. This is used internally, and should not be modified manually.
	UnusedN *string `json:"unused[n]"`
	// Configure an USB device (n is 0 to 4, for machine version >= 7.1 and ostype l26 or windows > 7, n can be up to 14).
	UsbN *string `json:"usb[n]"`
	// Number of hotplugged vcpus.
	Vcpus *int `json:"vcpus"`
	// Configure the VGA hardware.
	Vga *string `json:"vga"`
	// Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	VirtioN *string `json:"virtio[n]"`
	// Set VM Generation ID. Use '1' to autogenerate on create or update, pass '0' to disable explicitly.
	Vmgenid *string `json:"vmgenid"`
	// Default storage for VM state volumes/files.
	Vmstatestorage *string `json:"vmstatestorage"`
	// Create a virtual hardware watchdog device.
	Watchdog *string `json:"watchdog"`
}

type PutNodeQemuByVmidConfigResponse struct {
}

func (c *Client) PutNodeQemuByVmidConfig(Node string, Vmid int, parameters PutNodeQemuByVmidConfigParameters) (*PutNodeQemuByVmidConfigResponse, error) {

}

type GetNodeQemuByVmidPendingResponse struct {
}

func (c *Client) GetNodeQemuByVmidPending(Node string, Vmid int) (*GetNodeQemuByVmidPendingResponse, error) {

}

type GetNodeQemuByVmidCloudinitResponse struct {
}

func (c *Client) GetNodeQemuByVmidCloudinit(Node string, Vmid int) (*GetNodeQemuByVmidCloudinitResponse, error) {

}

type PutNodeQemuByVmidCloudinitResponse struct {
}

func (c *Client) PutNodeQemuByVmidCloudinit(Node string, Vmid int) (*PutNodeQemuByVmidCloudinitResponse, error) {

}

type GetNodeQemuByVmidCloudinitDumpParameters struct {
	// Config type.
	Type NodeQemuCloudinitDumpType `json:"type"`
}

type GetNodeQemuByVmidCloudinitDumpResponse struct {
}

func (c *Client) GetNodeQemuByVmidCloudinitDump(Node string, Vmid int, parameters GetNodeQemuByVmidCloudinitDumpParameters) (*GetNodeQemuByVmidCloudinitDumpResponse, error) {

}

type PutNodeQemuByVmidUnlinkParameters struct {
	// Force physical removal. Without this, we simple remove the disk from the config file and create an additional configuration entry called 'unused[n]', which contains the volume ID. Unlink of unused[n] always cause physical removal.
	Force *bool `json:"force"`
	// A list of disk IDs you want to delete.
	Idlist string `json:"idlist"`
}

type PutNodeQemuByVmidUnlinkResponse struct {
}

func (c *Client) PutNodeQemuByVmidUnlink(Node string, Vmid int, parameters PutNodeQemuByVmidUnlinkParameters) (*PutNodeQemuByVmidUnlinkResponse, error) {

}

type PostNodeQemuByVmidVncproxyParameters struct {
	// Generates a random password to be used as ticket instead of the API ticket.
	GeneratePassword *bool `json:"generate-password"`
	// starts websockify instead of vncproxy
	Websocket *bool `json:"websocket"`
}

type PostNodeQemuByVmidVncproxyResponse struct {
}

func (c *Client) PostNodeQemuByVmidVncproxy(Node string, Vmid int, parameters PostNodeQemuByVmidVncproxyParameters) (*PostNodeQemuByVmidVncproxyResponse, error) {

}

type PostNodeQemuByVmidTermproxyParameters struct {
	// opens a serial terminal (defaults to display)
	Serial *NodeQemuTermproxySerial `json:"serial"`
}

type PostNodeQemuByVmidTermproxyResponse struct {
}

func (c *Client) PostNodeQemuByVmidTermproxy(Node string, Vmid int, parameters PostNodeQemuByVmidTermproxyParameters) (*PostNodeQemuByVmidTermproxyResponse, error) {

}

type GetNodeQemuByVmidVncwebsocketParameters struct {
	// Port number returned by previous vncproxy call.
	Port int `json:"port"`
	// Ticket from previous call to vncproxy.
	Vncticket string `json:"vncticket"`
}

type GetNodeQemuByVmidVncwebsocketResponse struct {
}

func (c *Client) GetNodeQemuByVmidVncwebsocket(Node string, Vmid int, parameters GetNodeQemuByVmidVncwebsocketParameters) (*GetNodeQemuByVmidVncwebsocketResponse, error) {

}

type PostNodeQemuByVmidSpiceproxyParameters struct {
	// SPICE proxy server. This can be used by the client to specify the proxy server. All nodes in a cluster runs 'spiceproxy', so it is up to the client to choose one. By default, we return the node where the VM is currently running. As reasonable setting is to use same node you use to connect to the API (This is window.location.hostname for the JS GUI).
	Proxy *string `json:"proxy"`
}

type PostNodeQemuByVmidSpiceproxyResponse struct {
}

func (c *Client) PostNodeQemuByVmidSpiceproxy(Node string, Vmid int, parameters PostNodeQemuByVmidSpiceproxyParameters) (*PostNodeQemuByVmidSpiceproxyResponse, error) {

}

type GetNodeQemuByVmidStatusResponse struct {
}

func (c *Client) GetNodeQemuByVmidStatus(Node string, Vmid int) (*GetNodeQemuByVmidStatusResponse, error) {

}

type GetNodeQemuByVmidStatusCurrentResponse struct {
}

func (c *Client) GetNodeQemuByVmidStatusCurrent(Node string, Vmid int) (*GetNodeQemuByVmidStatusCurrentResponse, error) {

}

type PostNodeQemuByVmidStatusStartParameters struct {
	// Override QEMU's -cpu argument with the given string.
	ForceCpu *string `json:"force-cpu"`
	// Specifies the QEMU machine type.
	Machine *string `json:"machine"`
	// The cluster node name.
	Migratedfrom *string `json:"migratedfrom"`
	// CIDR of the (sub) network that is used for migration.
	MigrationNetwork *string `json:"migration_network"`
	// Migration traffic is encrypted using an SSH tunnel by default. On secure, completely private networks this can be disabled to increase performance.
	MigrationType *NodeQemuMigrationType `json:"migration_type"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
	// Some command save/restore state from this location.
	Stateuri *string `json:"stateuri"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	Targetstorage *string `json:"targetstorage"`
	// Wait maximal timeout seconds.
	Timeout *int `json:"timeout"`
}

type PostNodeQemuByVmidStatusStartResponse struct {
}

func (c *Client) PostNodeQemuByVmidStatusStart(Node string, Vmid int, parameters PostNodeQemuByVmidStatusStartParameters) (*PostNodeQemuByVmidStatusStartResponse, error) {

}

type PostNodeQemuByVmidStatusStopParameters struct {
	// Do not deactivate storage volumes.
	KeepActive *bool `json:"keepActive"`
	// The cluster node name.
	Migratedfrom *string `json:"migratedfrom"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
	// Wait maximal timeout seconds.
	Timeout *int `json:"timeout"`
}

type PostNodeQemuByVmidStatusStopResponse struct {
}

func (c *Client) PostNodeQemuByVmidStatusStop(Node string, Vmid int, parameters PostNodeQemuByVmidStatusStopParameters) (*PostNodeQemuByVmidStatusStopResponse, error) {

}

type PostNodeQemuByVmidStatusResetParameters struct {
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
}

type PostNodeQemuByVmidStatusResetResponse struct {
}

func (c *Client) PostNodeQemuByVmidStatusReset(Node string, Vmid int, parameters PostNodeQemuByVmidStatusResetParameters) (*PostNodeQemuByVmidStatusResetResponse, error) {

}

type PostNodeQemuByVmidStatusShutdownParameters struct {
	// Make sure the VM stops.
	ForceStop *bool `json:"forceStop"`
	// Do not deactivate storage volumes.
	KeepActive *bool `json:"keepActive"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
	// Wait maximal timeout seconds.
	Timeout *int `json:"timeout"`
}

type PostNodeQemuByVmidStatusShutdownResponse struct {
}

func (c *Client) PostNodeQemuByVmidStatusShutdown(Node string, Vmid int, parameters PostNodeQemuByVmidStatusShutdownParameters) (*PostNodeQemuByVmidStatusShutdownResponse, error) {

}

type PostNodeQemuByVmidStatusRebootParameters struct {
	// Wait maximal timeout seconds for the shutdown.
	Timeout *int `json:"timeout"`
}

type PostNodeQemuByVmidStatusRebootResponse struct {
}

func (c *Client) PostNodeQemuByVmidStatusReboot(Node string, Vmid int, parameters PostNodeQemuByVmidStatusRebootParameters) (*PostNodeQemuByVmidStatusRebootResponse, error) {

}

type PostNodeQemuByVmidStatusSuspendParameters struct {
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
	// The storage for the VM state
	Statestorage *string `json:"statestorage"`
	// If set, suspends the VM to disk. Will be resumed on next VM start.
	Todisk *bool `json:"todisk"`
}

type PostNodeQemuByVmidStatusSuspendResponse struct {
}

func (c *Client) PostNodeQemuByVmidStatusSuspend(Node string, Vmid int, parameters PostNodeQemuByVmidStatusSuspendParameters) (*PostNodeQemuByVmidStatusSuspendResponse, error) {

}

type PostNodeQemuByVmidStatusResumeParameters struct {
	Nocheck *bool `json:"nocheck"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
}

type PostNodeQemuByVmidStatusResumeResponse struct {
}

func (c *Client) PostNodeQemuByVmidStatusResume(Node string, Vmid int, parameters PostNodeQemuByVmidStatusResumeParameters) (*PostNodeQemuByVmidStatusResumeResponse, error) {

}

type PutNodeQemuByVmidSendkeyParameters struct {
	// The key (qemu monitor encoding).
	Key string `json:"key"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
}

type PutNodeQemuByVmidSendkeyResponse struct {
}

func (c *Client) PutNodeQemuByVmidSendkey(Node string, Vmid int, parameters PutNodeQemuByVmidSendkeyParameters) (*PutNodeQemuByVmidSendkeyResponse, error) {

}

type GetNodeQemuByVmidFeatureParameters struct {
	// Feature to check.
	Feature NodeFeatureFeature `json:"feature"`
	// The name of the snapshot.
	Snapname *string `json:"snapname"`
}

type GetNodeQemuByVmidFeatureResponse struct {
}

func (c *Client) GetNodeQemuByVmidFeature(Node string, Vmid int, parameters GetNodeQemuByVmidFeatureParameters) (*GetNodeQemuByVmidFeatureResponse, error) {

}

type PostNodeQemuByVmidCloneParameters struct {
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *int `json:"bwlimit"`
	// Description for the new VM.
	Description *string `json:"description"`
	// Target format for file storage. Only valid for full clone.
	Format *NodeQemuFormat `json:"format"`
	// Create a full copy of all disks. This is always done when you clone a normal VM. For VM templates, we try to create a linked clone by default.
	Full *bool `json:"full"`
	// Set a name for the new VM.
	Name *string `json:"name"`
	// VMID for the clone.
	Newid int `json:"newid"`
	// Add the new VM to the specified pool.
	Pool *string `json:"pool"`
	// The name of the snapshot.
	Snapname *string `json:"snapname"`
	// Target storage for full clone.
	Storage *string `json:"storage"`
	// Target node. Only allowed if the original VM is on shared storage.
	Target *string `json:"target"`
}

type PostNodeQemuByVmidCloneResponse struct {
}

func (c *Client) PostNodeQemuByVmidClone(Node string, Vmid int, parameters PostNodeQemuByVmidCloneParameters) (*PostNodeQemuByVmidCloneResponse, error) {

}

type PostNodeQemuByVmidMoveDiskParameters struct {
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *int `json:"bwlimit"`
	// Delete the original disk after successful copy. By default the original disk is kept as unused disk.
	Delete *bool `json:"delete"`
	// Prevent changes if current configuration file has different SHA1"
	// 		    ." digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// The disk you want to move.
	Disk NodeQemuMoveDiskDisk `json:"disk"`
	// Target Format.
	Format *NodeQemuFormat `json:"format"`
	// Target storage.
	Storage *string `json:"storage"`
	// Prevent changes if the current config file of the target VM has a"
	// 		    ." different SHA1 digest. This can be used to detect concurrent modifications.
	TargetDigest *string `json:"target-digest"`
	// The config key the disk will be moved to on the target VM (for example, ide0 or scsi1). Default is the source disk key.
	TargetDisk *NodeQemuMoveDiskDisk `json:"target-disk"`
	// The (unique) ID of the VM.
	TargetVmid *int `json:"target-vmid"`
}

type PostNodeQemuByVmidMoveDiskResponse struct {
}

func (c *Client) PostNodeQemuByVmidMoveDisk(Node string, Vmid int, parameters PostNodeQemuByVmidMoveDiskParameters) (*PostNodeQemuByVmidMoveDiskResponse, error) {

}

type GetNodeQemuByVmidMigrateParameters struct {
	// Target node.
	Target *string `json:"target"`
}

type GetNodeQemuByVmidMigrateResponse struct {
}

func (c *Client) GetNodeQemuByVmidMigrate(Node string, Vmid int, parameters GetNodeQemuByVmidMigrateParameters) (*GetNodeQemuByVmidMigrateResponse, error) {

}

type PostNodeQemuByVmidMigrateParameters struct {
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *int `json:"bwlimit"`
	// Allow to migrate VMs which use local devices. Only root may use this option.
	Force *bool `json:"force"`
	// CIDR of the (sub) network that is used for migration.
	MigrationNetwork *string `json:"migration_network"`
	// Migration traffic is encrypted using an SSH tunnel by default. On secure, completely private networks this can be disabled to increase performance.
	MigrationType *NodeQemuMigrationType `json:"migration_type"`
	// Use online/live migration if VM is running. Ignored if VM is stopped.
	Online *bool `json:"online"`
	// Target node.
	Target string `json:"target"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	Targetstorage *string `json:"targetstorage"`
	// Enable live storage migration for local disk
	WithLocalDisks *bool `json:"with-local-disks"`
}

type PostNodeQemuByVmidMigrateResponse struct {
}

func (c *Client) PostNodeQemuByVmidMigrate(Node string, Vmid int, parameters PostNodeQemuByVmidMigrateParameters) (*PostNodeQemuByVmidMigrateResponse, error) {

}

type PostNodeQemuByVmidRemoteMigrateParameters struct {
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *int `json:"bwlimit"`
	// Delete the original VM and related data after successful migration. By default the original VM is kept on the source cluster in a stopped state.
	Delete *bool `json:"delete"`
	// Use online/live migration if VM is running. Ignored if VM is stopped.
	Online *bool `json:"online"`
	// Mapping from source to target bridges. Providing only a single bridge ID maps all source bridges to that bridge. Providing the special value '1' will map each source bridge to itself.
	TargetBridge string `json:"target-bridge"`
	// Remote target endpoint
	TargetEndpoint string `json:"target-endpoint"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	TargetStorage string `json:"target-storage"`
	// The (unique) ID of the VM.
	TargetVmid *int `json:"target-vmid"`
}

type PostNodeQemuByVmidRemoteMigrateResponse struct {
}

func (c *Client) PostNodeQemuByVmidRemoteMigrate(Node string, Vmid int, parameters PostNodeQemuByVmidRemoteMigrateParameters) (*PostNodeQemuByVmidRemoteMigrateResponse, error) {

}

type PostNodeQemuByVmidMonitorParameters struct {
	// The monitor command.
	Command string `json:"command"`
}

type PostNodeQemuByVmidMonitorResponse struct {
}

func (c *Client) PostNodeQemuByVmidMonitor(Node string, Vmid int, parameters PostNodeQemuByVmidMonitorParameters) (*PostNodeQemuByVmidMonitorResponse, error) {

}

type PutNodeQemuByVmidResizeParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// The disk you want to resize.
	Disk NodeQemuDisk `json:"disk"`
	// The new size. With the `+` sign the value is added to the actual size of the volume and without it, the value is taken as an absolute one. Shrinking disk size is not supported.
	Size string `json:"size"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
}

type PutNodeQemuByVmidResizeResponse struct {
}

func (c *Client) PutNodeQemuByVmidResize(Node string, Vmid int, parameters PutNodeQemuByVmidResizeParameters) (*PutNodeQemuByVmidResizeResponse, error) {

}

type GetNodeQemuByVmidSnapshotResponse struct {
}

func (c *Client) GetNodeQemuByVmidSnapshot(Node string, Vmid int) (*GetNodeQemuByVmidSnapshotResponse, error) {

}

type PostNodeQemuByVmidSnapshotParameters struct {
	// A textual description or comment.
	Description *string `json:"description"`
	// The name of the snapshot.
	Snapname string `json:"snapname"`
	// Save the vmstate
	Vmstate *bool `json:"vmstate"`
}

type PostNodeQemuByVmidSnapshotResponse struct {
}

func (c *Client) PostNodeQemuByVmidSnapshot(Node string, Vmid int, parameters PostNodeQemuByVmidSnapshotParameters) (*PostNodeQemuByVmidSnapshotResponse, error) {

}

type DeleteNodeQemuByVmidSnapshotBySnapnameParameters struct {
	// For removal from config file, even if removing disk snapshots fails.
	Force *bool `json:"force"`
}

type DeleteNodeQemuByVmidSnapshotBySnapnameResponse struct {
}

func (c *Client) DeleteNodeQemuByVmidSnapshotBySnapname(Node string, Vmid int, Snapname string, parameters DeleteNodeQemuByVmidSnapshotBySnapnameParameters) (*DeleteNodeQemuByVmidSnapshotBySnapnameResponse, error) {

}

type GetNodeQemuByVmidSnapshotBySnapnameResponse struct {
}

func (c *Client) GetNodeQemuByVmidSnapshotBySnapname(Node string, Vmid int, Snapname string) (*GetNodeQemuByVmidSnapshotBySnapnameResponse, error) {

}

type GetNodeQemuByVmidSnapshotBySnapnameConfigResponse struct {
}

func (c *Client) GetNodeQemuByVmidSnapshotBySnapnameConfig(Node string, Vmid int, Snapname string) (*GetNodeQemuByVmidSnapshotBySnapnameConfigResponse, error) {

}

type PutNodeQemuByVmidSnapshotBySnapnameConfigParameters struct {
	// A textual description or comment.
	Description *string `json:"description"`
}

type PutNodeQemuByVmidSnapshotBySnapnameConfigResponse struct {
}

func (c *Client) PutNodeQemuByVmidSnapshotBySnapnameConfig(Node string, Vmid int, Snapname string, parameters PutNodeQemuByVmidSnapshotBySnapnameConfigParameters) (*PutNodeQemuByVmidSnapshotBySnapnameConfigResponse, error) {

}

type PostNodeQemuByVmidSnapshotBySnapnameRollbackParameters struct {
	// Whether the VM should get started after rolling back successfully. (Note: VMs will be automatically started if the snapshot includes RAM.)
	Start *bool `json:"start"`
}

type PostNodeQemuByVmidSnapshotBySnapnameRollbackResponse struct {
}

func (c *Client) PostNodeQemuByVmidSnapshotBySnapnameRollback(Node string, Vmid int, Snapname string, parameters PostNodeQemuByVmidSnapshotBySnapnameRollbackParameters) (*PostNodeQemuByVmidSnapshotBySnapnameRollbackResponse, error) {

}

type PostNodeQemuByVmidTemplateParameters struct {
	// If you want to convert only 1 disk to base image.
	Disk *NodeQemuDisk `json:"disk"`
}

type PostNodeQemuByVmidTemplateResponse struct {
}

func (c *Client) PostNodeQemuByVmidTemplate(Node string, Vmid int, parameters PostNodeQemuByVmidTemplateParameters) (*PostNodeQemuByVmidTemplateResponse, error) {

}

type PostNodeQemuByVmidMtunnelParameters struct {
	// List of network bridges to check availability. Will be checked again for actually used bridges during migration.
	Bridges *string `json:"bridges"`
	// List of storages to check permission and availability. Will be checked again for all actually used storages during migration.
	Storages *string `json:"storages"`
}

type PostNodeQemuByVmidMtunnelResponse struct {
}

func (c *Client) PostNodeQemuByVmidMtunnel(Node string, Vmid int, parameters PostNodeQemuByVmidMtunnelParameters) (*PostNodeQemuByVmidMtunnelResponse, error) {

}

type GetNodeQemuByVmidMtunnelwebsocketParameters struct {
	// unix socket to forward to
	Socket string `json:"socket"`
	// ticket return by initial 'mtunnel' API call, or retrieved via 'ticket' tunnel command
	Ticket string `json:"ticket"`
}

type GetNodeQemuByVmidMtunnelwebsocketResponse struct {
}

func (c *Client) GetNodeQemuByVmidMtunnelwebsocket(Node string, Vmid int, parameters GetNodeQemuByVmidMtunnelwebsocketParameters) (*GetNodeQemuByVmidMtunnelwebsocketResponse, error) {

}

type GetNodeLxcResponse struct {
}

func (c *Client) GetNodeLxc(Node string) (*GetNodeLxcResponse, error) {

}

type PostNodeLxcParameters struct {
	// OS architecture type.
	Arch *NodeLxcArch `json:"arch"`
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *float64 `json:"bwlimit"`
	// Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
	Cmode *NodeLxcCmode `json:"cmode"`
	// Attach a console device (/dev/console) to the container.
	Console *bool `json:"console"`
	// The number of cores assigned to the container. A container can use all available cores by default.
	Cores *int `json:"cores"`
	// Limit of CPU usage.
	//
	// NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
	Cpulimit *float64 `json:"cpulimit"`
	// CPU weight for a container, will be clamped to [1, 10000] in cgroup v2.
	Cpuunits *int `json:"cpuunits"`
	// Try to be more verbose. For now this only enables debug log-level on start.
	Debug *bool `json:"debug"`
	// Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
	Description *string `json:"description"`
	// Allow containers access to advanced features.
	Features *string `json:"features"`
	// Allow to overwrite existing container.
	Force *bool `json:"force"`
	// Script that will be exectued during various steps in the containers lifetime.
	Hookscript *string `json:"hookscript"`
	// Set a host name for the container.
	Hostname *string `json:"hostname"`
	// Ignore errors when extracting the template.
	IgnoreUnpackErrors *bool `json:"ignore-unpack-errors"`
	// Lock/unlock the container.
	Lock *NodeLxcLock `json:"lock"`
	// Amount of RAM for the container in MB.
	Memory *int `json:"memory"`
	// Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	MpN *string `json:"mp[n]"`
	// Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Nameserver *string `json:"nameserver"`
	// Specifies network interfaces for the container.
	NetN *string `json:"net[n]"`
	// Specifies whether a container will be started during system bootup.
	Onboot *bool `json:"onboot"`
	// The OS template or backup file.
	Ostemplate string `json:"ostemplate"`
	// OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup.
	Ostype *NodeLxcOstype `json:"ostype"`
	// Sets root password inside container.
	Password *string `json:"password"`
	// Add the VM to the specified pool.
	Pool *string `json:"pool"`
	// Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
	Protection *bool `json:"protection"`
	// Mark this as restore task.
	Restore *bool `json:"restore"`
	// Use volume as container root.
	Rootfs *string `json:"rootfs"`
	// Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Searchdomain *string `json:"searchdomain"`
	// Setup public SSH keys (one key per line, OpenSSH format).
	SshPublicKeys *string `json:"ssh-public-keys"`
	// Start the CT after its creation finished successfully.
	Start *bool `json:"start"`
	// Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Startup *string `json:"startup"`
	// Default Storage.
	Storage *string `json:"storage"`
	// Amount of SWAP for the container in MB.
	Swap *int `json:"swap"`
	// Tags of the Container. This is only meta information.
	Tags *string `json:"tags"`
	// Enable/disable Template.
	Template *bool `json:"template"`
	// Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
	Timezone *string `json:"timezone"`
	// Specify the number of tty available to the container
	Tty *int `json:"tty"`
	// Assign a unique random ethernet address.
	Unique *bool `json:"unique"`
	// Makes the container run as unprivileged user. (Should not be modified manually.)
	Unprivileged *bool `json:"unprivileged"`
	// Reference to unused volumes. This is used internally, and should not be modified manually.
	UnusedN *string `json:"unused[n]"`
	// The (unique) ID of the VM.
	Vmid int `json:"vmid"`
}

type PostNodeLxcResponse struct {
}

func (c *Client) PostNodeLxc(Node string, parameters PostNodeLxcParameters) (*PostNodeLxcResponse, error) {

}

type DeleteNodeLxcByVmidParameters struct {
	// If set, destroy additionally all disks with the VMID from all enabled storages which are not referenced in the config.
	DestroyUnreferencedDisks *bool `json:"destroy-unreferenced-disks"`
	// Force destroy, even if running.
	Force *bool `json:"force"`
	// Remove container from all related configurations. For example, backup jobs, replication jobs or HA. Related ACLs and Firewall entries will *always* be removed.
	Purge *bool `json:"purge"`
}

type DeleteNodeLxcByVmidResponse struct {
}

func (c *Client) DeleteNodeLxcByVmid(Node string, Vmid int, parameters DeleteNodeLxcByVmidParameters) (*DeleteNodeLxcByVmidResponse, error) {

}

type GetNodeLxcByVmidResponse struct {
}

func (c *Client) GetNodeLxcByVmid(Node string, Vmid int) (*GetNodeLxcByVmidResponse, error) {

}

type GetNodeLxcByVmidConfigParameters struct {
	// Get current values (instead of pending values).
	Current *bool `json:"current"`
	// Fetch config values from given snapshot.
	Snapshot *string `json:"snapshot"`
}

type GetNodeLxcByVmidConfigResponse struct {
}

func (c *Client) GetNodeLxcByVmidConfig(Node string, Vmid int, parameters GetNodeLxcByVmidConfigParameters) (*GetNodeLxcByVmidConfigResponse, error) {

}

type PutNodeLxcByVmidConfigParameters struct {
	// OS architecture type.
	Arch *NodeLxcArch `json:"arch"`
	// Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to 'console' it tries to attach to /dev/console instead. If you set cmode to 'shell', it simply invokes a shell inside the container (no login).
	Cmode *NodeLxcCmode `json:"cmode"`
	// Attach a console device (/dev/console) to the container.
	Console *bool `json:"console"`
	// The number of cores assigned to the container. A container can use all available cores by default.
	Cores *int `json:"cores"`
	// Limit of CPU usage.
	//
	// NOTE: If the computer has 2 CPUs, it has a total of '2' CPU time. Value '0' indicates no CPU limit.
	Cpulimit *float64 `json:"cpulimit"`
	// CPU weight for a container, will be clamped to [1, 10000] in cgroup v2.
	Cpuunits *int `json:"cpuunits"`
	// Try to be more verbose. For now this only enables debug log-level on start.
	Debug *bool `json:"debug"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Description for the Container. Shown in the web-interface CT's summary. This is saved as comment inside the configuration file.
	Description *string `json:"description"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Allow containers access to advanced features.
	Features *string `json:"features"`
	// Script that will be exectued during various steps in the containers lifetime.
	Hookscript *string `json:"hookscript"`
	// Set a host name for the container.
	Hostname *string `json:"hostname"`
	// Lock/unlock the container.
	Lock *NodeLxcLock `json:"lock"`
	// Amount of RAM for the container in MB.
	Memory *int `json:"memory"`
	// Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	MpN *string `json:"mp[n]"`
	// Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Nameserver *string `json:"nameserver"`
	// Specifies network interfaces for the container.
	NetN *string `json:"net[n]"`
	// Specifies whether a container will be started during system bootup.
	Onboot *bool `json:"onboot"`
	// OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/<ostype>.common.conf. Value 'unmanaged' can be used to skip and OS specific setup.
	Ostype *NodeLxcOstype `json:"ostype"`
	// Sets the protection flag of the container. This will prevent the CT or CT's disk remove/update operation.
	Protection *bool `json:"protection"`
	// Revert a pending change.
	Revert *string `json:"revert"`
	// Use volume as container root.
	Rootfs *string `json:"rootfs"`
	// Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver.
	Searchdomain *string `json:"searchdomain"`
	// Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the 'up' or 'down' delay in seconds, which specifies a delay to wait before the next VM is started or stopped.
	Startup *string `json:"startup"`
	// Amount of SWAP for the container in MB.
	Swap *int `json:"swap"`
	// Tags of the Container. This is only meta information.
	Tags *string `json:"tags"`
	// Enable/disable Template.
	Template *bool `json:"template"`
	// Time zone to use in the container. If option isn't set, then nothing will be done. Can be set to 'host' to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab
	Timezone *string `json:"timezone"`
	// Specify the number of tty available to the container
	Tty *int `json:"tty"`
	// Makes the container run as unprivileged user. (Should not be modified manually.)
	Unprivileged *bool `json:"unprivileged"`
	// Reference to unused volumes. This is used internally, and should not be modified manually.
	UnusedN *string `json:"unused[n]"`
}

type PutNodeLxcByVmidConfigResponse struct {
}

func (c *Client) PutNodeLxcByVmidConfig(Node string, Vmid int, parameters PutNodeLxcByVmidConfigParameters) (*PutNodeLxcByVmidConfigResponse, error) {

}

type GetNodeLxcByVmidStatusResponse struct {
}

func (c *Client) GetNodeLxcByVmidStatus(Node string, Vmid int) (*GetNodeLxcByVmidStatusResponse, error) {

}

type GetNodeLxcByVmidStatusCurrentResponse struct {
}

func (c *Client) GetNodeLxcByVmidStatusCurrent(Node string, Vmid int) (*GetNodeLxcByVmidStatusCurrentResponse, error) {

}

type PostNodeLxcByVmidStatusStartParameters struct {
	// If set, enables very verbose debug log-level on start.
	Debug *bool `json:"debug"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
}

type PostNodeLxcByVmidStatusStartResponse struct {
}

func (c *Client) PostNodeLxcByVmidStatusStart(Node string, Vmid int, parameters PostNodeLxcByVmidStatusStartParameters) (*PostNodeLxcByVmidStatusStartResponse, error) {

}

type PostNodeLxcByVmidStatusStopParameters struct {
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock"`
}

type PostNodeLxcByVmidStatusStopResponse struct {
}

func (c *Client) PostNodeLxcByVmidStatusStop(Node string, Vmid int, parameters PostNodeLxcByVmidStatusStopParameters) (*PostNodeLxcByVmidStatusStopResponse, error) {

}

type PostNodeLxcByVmidStatusShutdownParameters struct {
	// Make sure the Container stops.
	ForceStop *bool `json:"forceStop"`
	// Wait maximal timeout seconds.
	Timeout *int `json:"timeout"`
}

type PostNodeLxcByVmidStatusShutdownResponse struct {
}

func (c *Client) PostNodeLxcByVmidStatusShutdown(Node string, Vmid int, parameters PostNodeLxcByVmidStatusShutdownParameters) (*PostNodeLxcByVmidStatusShutdownResponse, error) {

}

type PostNodeLxcByVmidStatusSuspendResponse struct {
}

func (c *Client) PostNodeLxcByVmidStatusSuspend(Node string, Vmid int) (*PostNodeLxcByVmidStatusSuspendResponse, error) {

}

type PostNodeLxcByVmidStatusResumeResponse struct {
}

func (c *Client) PostNodeLxcByVmidStatusResume(Node string, Vmid int) (*PostNodeLxcByVmidStatusResumeResponse, error) {

}

type PostNodeLxcByVmidStatusRebootParameters struct {
	// Wait maximal timeout seconds for the shutdown.
	Timeout *int `json:"timeout"`
}

type PostNodeLxcByVmidStatusRebootResponse struct {
}

func (c *Client) PostNodeLxcByVmidStatusReboot(Node string, Vmid int, parameters PostNodeLxcByVmidStatusRebootParameters) (*PostNodeLxcByVmidStatusRebootResponse, error) {

}

type GetNodeLxcByVmidSnapshotResponse struct {
}

func (c *Client) GetNodeLxcByVmidSnapshot(Node string, Vmid int) (*GetNodeLxcByVmidSnapshotResponse, error) {

}

type PostNodeLxcByVmidSnapshotParameters struct {
	// A textual description or comment.
	Description *string `json:"description"`
	// The name of the snapshot.
	Snapname string `json:"snapname"`
}

type PostNodeLxcByVmidSnapshotResponse struct {
}

func (c *Client) PostNodeLxcByVmidSnapshot(Node string, Vmid int, parameters PostNodeLxcByVmidSnapshotParameters) (*PostNodeLxcByVmidSnapshotResponse, error) {

}

type DeleteNodeLxcByVmidSnapshotBySnapnameParameters struct {
	// For removal from config file, even if removing disk snapshots fails.
	Force *bool `json:"force"`
}

type DeleteNodeLxcByVmidSnapshotBySnapnameResponse struct {
}

func (c *Client) DeleteNodeLxcByVmidSnapshotBySnapname(Node string, Vmid int, Snapname string, parameters DeleteNodeLxcByVmidSnapshotBySnapnameParameters) (*DeleteNodeLxcByVmidSnapshotBySnapnameResponse, error) {

}

type GetNodeLxcByVmidSnapshotBySnapnameResponse struct {
}

func (c *Client) GetNodeLxcByVmidSnapshotBySnapname(Node string, Vmid int, Snapname string) (*GetNodeLxcByVmidSnapshotBySnapnameResponse, error) {

}

type PostNodeLxcByVmidSnapshotBySnapnameRollbackParameters struct {
	// Whether the container should get started after rolling back successfully
	Start *bool `json:"start"`
}

type PostNodeLxcByVmidSnapshotBySnapnameRollbackResponse struct {
}

func (c *Client) PostNodeLxcByVmidSnapshotBySnapnameRollback(Node string, Vmid int, Snapname string, parameters PostNodeLxcByVmidSnapshotBySnapnameRollbackParameters) (*PostNodeLxcByVmidSnapshotBySnapnameRollbackResponse, error) {

}

type GetNodeLxcByVmidSnapshotBySnapnameConfigResponse struct {
}

func (c *Client) GetNodeLxcByVmidSnapshotBySnapnameConfig(Node string, Vmid int, Snapname string) (*GetNodeLxcByVmidSnapshotBySnapnameConfigResponse, error) {

}

type PutNodeLxcByVmidSnapshotBySnapnameConfigParameters struct {
	// A textual description or comment.
	Description *string `json:"description"`
}

type PutNodeLxcByVmidSnapshotBySnapnameConfigResponse struct {
}

func (c *Client) PutNodeLxcByVmidSnapshotBySnapnameConfig(Node string, Vmid int, Snapname string, parameters PutNodeLxcByVmidSnapshotBySnapnameConfigParameters) (*PutNodeLxcByVmidSnapshotBySnapnameConfigResponse, error) {

}

type GetNodeLxcByVmidFirewallResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewall(Node string, Vmid int) (*GetNodeLxcByVmidFirewallResponse, error) {

}

type GetNodeLxcByVmidFirewallRulesResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallRules(Node string, Vmid int) (*GetNodeLxcByVmidFirewallRulesResponse, error) {

}

type PostNodeLxcByVmidFirewallRulesParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Update rule at position <pos>.
	Pos *int `json:"pos"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type FirewallType `json:"type"`
}

type PostNodeLxcByVmidFirewallRulesResponse struct {
}

func (c *Client) PostNodeLxcByVmidFirewallRules(Node string, Vmid int, parameters PostNodeLxcByVmidFirewallRulesParameters) (*PostNodeLxcByVmidFirewallRulesResponse, error) {

}

type DeleteNodeLxcByVmidFirewallRuleParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteNodeLxcByVmidFirewallRuleResponse struct {
}

func (c *Client) DeleteNodeLxcByVmidFirewallRule(Node string, Vmid int, Pos int, parameters DeleteNodeLxcByVmidFirewallRuleParameters) (*DeleteNodeLxcByVmidFirewallRuleResponse, error) {

}

type GetNodeLxcByVmidFirewallRuleResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallRule(Node string, Vmid int, Pos int) (*GetNodeLxcByVmidFirewallRuleResponse, error) {

}

type PutNodeLxcByVmidFirewallRuleParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action *string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Move rule to new position <moveto>. Other arguments are ignored.
	Moveto *int `json:"moveto"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type *FirewallType `json:"type"`
}

type PutNodeLxcByVmidFirewallRuleResponse struct {
}

func (c *Client) PutNodeLxcByVmidFirewallRule(Node string, Vmid int, Pos int, parameters PutNodeLxcByVmidFirewallRuleParameters) (*PutNodeLxcByVmidFirewallRuleResponse, error) {

}

type GetNodeLxcByVmidFirewallAliasesResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallAliases(Node string, Vmid int) (*GetNodeLxcByVmidFirewallAliasesResponse, error) {

}

type PostNodeLxcByVmidFirewallAliasesParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	// Alias name.
	Name string `json:"name"`
}

type PostNodeLxcByVmidFirewallAliasesResponse struct {
}

func (c *Client) PostNodeLxcByVmidFirewallAliases(Node string, Vmid int, parameters PostNodeLxcByVmidFirewallAliasesParameters) (*PostNodeLxcByVmidFirewallAliasesResponse, error) {

}

type DeleteNodeLxcByVmidFirewallAliasParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteNodeLxcByVmidFirewallAliasResponse struct {
}

func (c *Client) DeleteNodeLxcByVmidFirewallAlias(Node string, Vmid int, Name string, parameters DeleteNodeLxcByVmidFirewallAliasParameters) (*DeleteNodeLxcByVmidFirewallAliasResponse, error) {

}

type GetNodeLxcByVmidFirewallAliasResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallAlias(Node string, Vmid int, Name string) (*GetNodeLxcByVmidFirewallAliasResponse, error) {

}

type PutNodeLxcByVmidFirewallAliasParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Rename an existing alias.
	Rename *string `json:"rename"`
}

type PutNodeLxcByVmidFirewallAliasResponse struct {
}

func (c *Client) PutNodeLxcByVmidFirewallAlias(Node string, Vmid int, Name string, parameters PutNodeLxcByVmidFirewallAliasParameters) (*PutNodeLxcByVmidFirewallAliasResponse, error) {

}

type GetNodeLxcByVmidFirewallIpsetResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallIpset(Node string, Vmid int) (*GetNodeLxcByVmidFirewallIpsetResponse, error) {

}

type PostNodeLxcByVmidFirewallIpsetParameters struct {
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// IP set name.
	Name string `json:"name"`
	// Rename an existing IPSet. You can set 'rename' to the same value as 'name' to update the 'comment' of an existing IPSet.
	Rename *string `json:"rename"`
}

type PostNodeLxcByVmidFirewallIpsetResponse struct {
}

func (c *Client) PostNodeLxcByVmidFirewallIpset(Node string, Vmid int, parameters PostNodeLxcByVmidFirewallIpsetParameters) (*PostNodeLxcByVmidFirewallIpsetResponse, error) {

}

type DeleteNodeLxcByVmidFirewallIpsetByNameParameters struct {
	// Delete all members of the IPSet, if there are any.
	Force *bool `json:"force"`
}

type DeleteNodeLxcByVmidFirewallIpsetByNameResponse struct {
}

func (c *Client) DeleteNodeLxcByVmidFirewallIpsetByName(Node string, Vmid int, Name string, parameters DeleteNodeLxcByVmidFirewallIpsetByNameParameters) (*DeleteNodeLxcByVmidFirewallIpsetByNameResponse, error) {

}

type GetNodeLxcByVmidFirewallIpsetByNameResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallIpsetByName(Node string, Vmid int, Name string) (*GetNodeLxcByVmidFirewallIpsetByNameResponse, error) {

}

type PostNodeLxcByVmidFirewallIpsetByNameParameters struct {
	// Network/IP specification in CIDR format.
	Cidr    string  `json:"cidr"`
	Comment *string `json:"comment"`
	Nomatch *bool   `json:"nomatch"`
}

type PostNodeLxcByVmidFirewallIpsetByNameResponse struct {
}

func (c *Client) PostNodeLxcByVmidFirewallIpsetByName(Node string, Vmid int, Name string, parameters PostNodeLxcByVmidFirewallIpsetByNameParameters) (*PostNodeLxcByVmidFirewallIpsetByNameResponse, error) {

}

type DeleteNodeLxcByVmidFirewallIpsetByNameByCidrParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteNodeLxcByVmidFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) DeleteNodeLxcByVmidFirewallIpsetByNameByCidr(Node string, Vmid int, Name string, Cidr string, parameters DeleteNodeLxcByVmidFirewallIpsetByNameByCidrParameters) (*DeleteNodeLxcByVmidFirewallIpsetByNameByCidrResponse, error) {

}

type GetNodeLxcByVmidFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallIpsetByNameByCidr(Node string, Vmid int, Name string, Cidr string) (*GetNodeLxcByVmidFirewallIpsetByNameByCidrResponse, error) {

}

type PutNodeLxcByVmidFirewallIpsetByNameByCidrParameters struct {
	Comment *string `json:"comment"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest  *string `json:"digest"`
	Nomatch *bool   `json:"nomatch"`
}

type PutNodeLxcByVmidFirewallIpsetByNameByCidrResponse struct {
}

func (c *Client) PutNodeLxcByVmidFirewallIpsetByNameByCidr(Node string, Vmid int, Name string, Cidr string, parameters PutNodeLxcByVmidFirewallIpsetByNameByCidrParameters) (*PutNodeLxcByVmidFirewallIpsetByNameByCidrResponse, error) {

}

type GetNodeLxcByVmidFirewallOptionsResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallOptions(Node string, Vmid int) (*GetNodeLxcByVmidFirewallOptionsResponse, error) {

}

type PutNodeLxcByVmidFirewallOptionsParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Enable DHCP.
	Dhcp *bool `json:"dhcp"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Enable/disable firewall rules.
	Enable *bool `json:"enable"`
	// Enable default IP filters. This is equivalent to adding an empty ipfilter-net<id> ipset for every interface. Such ipsets implicitly contain sane default restrictions such as restricting IPv6 link local addresses to the one derived from the interface's MAC address. For containers the configured IP addresses will be implicitly added.
	Ipfilter *bool `json:"ipfilter"`
	// Log level for incoming traffic.
	LogLevelIn *FirewallLog `json:"log_level_in"`
	// Log level for outgoing traffic.
	LogLevelOut *FirewallLog `json:"log_level_out"`
	// Enable/disable MAC address filter.
	Macfilter *bool `json:"macfilter"`
	// Enable NDP (Neighbor Discovery Protocol).
	Ndp *bool `json:"ndp"`
	// Input policy.
	PolicyIn *FirewallOptionsPolicy `json:"policy_in"`
	// Output policy.
	PolicyOut *FirewallOptionsPolicy `json:"policy_out"`
	// Allow sending Router Advertisement.
	Radv *bool `json:"radv"`
}

type PutNodeLxcByVmidFirewallOptionsResponse struct {
}

func (c *Client) PutNodeLxcByVmidFirewallOptions(Node string, Vmid int, parameters PutNodeLxcByVmidFirewallOptionsParameters) (*PutNodeLxcByVmidFirewallOptionsResponse, error) {

}

type GetNodeLxcByVmidFirewallLogParameters struct {
	Limit *int `json:"limit"`
	// Display log since this UNIX epoch.
	Since *int `json:"since"`
	Start *int `json:"start"`
	// Display log until this UNIX epoch.
	Until *int `json:"until"`
}

type GetNodeLxcByVmidFirewallLogResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallLog(Node string, Vmid int, parameters GetNodeLxcByVmidFirewallLogParameters) (*GetNodeLxcByVmidFirewallLogResponse, error) {

}

type GetNodeLxcByVmidFirewallRefsParameters struct {
	// Only list references of specified type.
	Type *FirewallRefsType `json:"type"`
}

type GetNodeLxcByVmidFirewallRefsResponse struct {
}

func (c *Client) GetNodeLxcByVmidFirewallRefs(Node string, Vmid int, parameters GetNodeLxcByVmidFirewallRefsParameters) (*GetNodeLxcByVmidFirewallRefsResponse, error) {

}

type GetNodeLxcByVmidRrdParameters struct {
	// The RRD consolidation function
	Cf *NodeRrdCf `json:"cf"`
	// The list of datasources you want to display.
	Ds string `json:"ds"`
	// Specify the time frame you are interested in.
	Timeframe NodeRrdTimeframe `json:"timeframe"`
}

type GetNodeLxcByVmidRrdResponse struct {
}

func (c *Client) GetNodeLxcByVmidRrd(Node string, Vmid int, parameters GetNodeLxcByVmidRrdParameters) (*GetNodeLxcByVmidRrdResponse, error) {

}

type GetNodeLxcByVmidRrddataParameters struct {
	// The RRD consolidation function
	Cf *NodeRrdCf `json:"cf"`
	// Specify the time frame you are interested in.
	Timeframe NodeRrdTimeframe `json:"timeframe"`
}

type GetNodeLxcByVmidRrddataResponse struct {
}

func (c *Client) GetNodeLxcByVmidRrddata(Node string, Vmid int, parameters GetNodeLxcByVmidRrddataParameters) (*GetNodeLxcByVmidRrddataResponse, error) {

}

type PostNodeLxcByVmidVncproxyParameters struct {
	// sets the height of the console in pixels.
	Height *int `json:"height"`
	// use websocket instead of standard VNC.
	Websocket *bool `json:"websocket"`
	// sets the width of the console in pixels.
	Width *int `json:"width"`
}

type PostNodeLxcByVmidVncproxyResponse struct {
}

func (c *Client) PostNodeLxcByVmidVncproxy(Node string, Vmid int, parameters PostNodeLxcByVmidVncproxyParameters) (*PostNodeLxcByVmidVncproxyResponse, error) {

}

type PostNodeLxcByVmidTermproxyResponse struct {
}

func (c *Client) PostNodeLxcByVmidTermproxy(Node string, Vmid int) (*PostNodeLxcByVmidTermproxyResponse, error) {

}

type GetNodeLxcByVmidVncwebsocketParameters struct {
	// Port number returned by previous vncproxy call.
	Port int `json:"port"`
	// Ticket from previous call to vncproxy.
	Vncticket string `json:"vncticket"`
}

type GetNodeLxcByVmidVncwebsocketResponse struct {
}

func (c *Client) GetNodeLxcByVmidVncwebsocket(Node string, Vmid int, parameters GetNodeLxcByVmidVncwebsocketParameters) (*GetNodeLxcByVmidVncwebsocketResponse, error) {

}

type PostNodeLxcByVmidSpiceproxyParameters struct {
	// SPICE proxy server. This can be used by the client to specify the proxy server. All nodes in a cluster runs 'spiceproxy', so it is up to the client to choose one. By default, we return the node where the VM is currently running. As reasonable setting is to use same node you use to connect to the API (This is window.location.hostname for the JS GUI).
	Proxy *string `json:"proxy"`
}

type PostNodeLxcByVmidSpiceproxyResponse struct {
}

func (c *Client) PostNodeLxcByVmidSpiceproxy(Node string, Vmid int, parameters PostNodeLxcByVmidSpiceproxyParameters) (*PostNodeLxcByVmidSpiceproxyResponse, error) {

}

type PostNodeLxcByVmidRemoteMigrateParameters struct {
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *float64 `json:"bwlimit"`
	// Delete the original CT and related data after successful migration. By default the original CT is kept on the source cluster in a stopped state.
	Delete *bool `json:"delete"`
	// Use online/live migration.
	Online *bool `json:"online"`
	// Use restart migration
	Restart *bool `json:"restart"`
	// Mapping from source to target bridges. Providing only a single bridge ID maps all source bridges to that bridge. Providing the special value '1' will map each source bridge to itself.
	TargetBridge string `json:"target-bridge"`
	// Remote target endpoint
	TargetEndpoint string `json:"target-endpoint"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	TargetStorage string `json:"target-storage"`
	// The (unique) ID of the VM.
	TargetVmid *int `json:"target-vmid"`
	// Timeout in seconds for shutdown for restart migration
	Timeout *int `json:"timeout"`
}

type PostNodeLxcByVmidRemoteMigrateResponse struct {
}

func (c *Client) PostNodeLxcByVmidRemoteMigrate(Node string, Vmid int, parameters PostNodeLxcByVmidRemoteMigrateParameters) (*PostNodeLxcByVmidRemoteMigrateResponse, error) {

}

type PostNodeLxcByVmidMigrateParameters struct {
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *float64 `json:"bwlimit"`
	// Use online/live migration.
	Online *bool `json:"online"`
	// Use restart migration
	Restart *bool `json:"restart"`
	// Target node.
	Target string `json:"target"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage. Providing the special value '1' will map each source storage to itself.
	TargetStorage *string `json:"target-storage"`
	// Timeout in seconds for shutdown for restart migration
	Timeout *int `json:"timeout"`
}

type PostNodeLxcByVmidMigrateResponse struct {
}

func (c *Client) PostNodeLxcByVmidMigrate(Node string, Vmid int, parameters PostNodeLxcByVmidMigrateParameters) (*PostNodeLxcByVmidMigrateResponse, error) {

}

type GetNodeLxcByVmidFeatureParameters struct {
	// Feature to check.
	Feature NodeFeatureFeature `json:"feature"`
	// The name of the snapshot.
	Snapname *string `json:"snapname"`
}

type GetNodeLxcByVmidFeatureResponse struct {
}

func (c *Client) GetNodeLxcByVmidFeature(Node string, Vmid int, parameters GetNodeLxcByVmidFeatureParameters) (*GetNodeLxcByVmidFeatureResponse, error) {

}

type PostNodeLxcByVmidTemplateResponse struct {
}

func (c *Client) PostNodeLxcByVmidTemplate(Node string, Vmid int) (*PostNodeLxcByVmidTemplateResponse, error) {

}

type PostNodeLxcByVmidCloneParameters struct {
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *float64 `json:"bwlimit"`
	// Description for the new CT.
	Description *string `json:"description"`
	// Create a full copy of all disks. This is always done when you clone a normal CT. For CT templates, we try to create a linked clone by default.
	Full *bool `json:"full"`
	// Set a hostname for the new CT.
	Hostname *string `json:"hostname"`
	// VMID for the clone.
	Newid int `json:"newid"`
	// Add the new CT to the specified pool.
	Pool *string `json:"pool"`
	// The name of the snapshot.
	Snapname *string `json:"snapname"`
	// Target storage for full clone.
	Storage *string `json:"storage"`
	// Target node. Only allowed if the original VM is on shared storage.
	Target *string `json:"target"`
}

type PostNodeLxcByVmidCloneResponse struct {
}

func (c *Client) PostNodeLxcByVmidClone(Node string, Vmid int, parameters PostNodeLxcByVmidCloneParameters) (*PostNodeLxcByVmidCloneResponse, error) {

}

type PutNodeLxcByVmidResizeParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// The disk you want to resize.
	Disk NodeLxcResizeDisk `json:"disk"`
	// The new size. With the '+' sign the value is added to the actual size of the volume and without it, the value is taken as an absolute one. Shrinking disk size is not supported.
	Size string `json:"size"`
}

type PutNodeLxcByVmidResizeResponse struct {
}

func (c *Client) PutNodeLxcByVmidResize(Node string, Vmid int, parameters PutNodeLxcByVmidResizeParameters) (*PutNodeLxcByVmidResizeResponse, error) {

}

type PostNodeLxcByVmidMoveVolumeParameters struct {
	// Override I/O bandwidth limit (in KiB/s).
	Bwlimit *float64 `json:"bwlimit"`
	// Delete the original volume after successful copy. By default the original is kept as an unused volume entry.
	Delete *bool `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 " .
	// 		    "digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Target Storage.
	Storage *string `json:"storage"`
	// Prevent changes if current configuration file of the target " .
	// 		    "container has a different SHA1 digest. This can be used to prevent " .
	// 		    "concurrent modifications.
	TargetDigest *string `json:"target-digest"`
	// The (unique) ID of the VM.
	TargetVmid *int `json:"target-vmid"`
	// The config key the volume will be moved to. Default is the source volume key.
	TargetVolume *NodeLxcMoveVolumeVolume `json:"target-volume"`
	// Volume which will be moved.
	Volume NodeLxcMoveVolumeVolume `json:"volume"`
}

type PostNodeLxcByVmidMoveVolumeResponse struct {
}

func (c *Client) PostNodeLxcByVmidMoveVolume(Node string, Vmid int, parameters PostNodeLxcByVmidMoveVolumeParameters) (*PostNodeLxcByVmidMoveVolumeResponse, error) {

}

type GetNodeLxcByVmidPendingResponse struct {
}

func (c *Client) GetNodeLxcByVmidPending(Node string, Vmid int) (*GetNodeLxcByVmidPendingResponse, error) {

}

type PostNodeLxcByVmidMtunnelParameters struct {
	// List of network bridges to check availability. Will be checked again for actually used bridges during migration.
	Bridges *string `json:"bridges"`
	// List of storages to check permission and availability. Will be checked again for all actually used storages during migration.
	Storages *string `json:"storages"`
}

type PostNodeLxcByVmidMtunnelResponse struct {
}

func (c *Client) PostNodeLxcByVmidMtunnel(Node string, Vmid int, parameters PostNodeLxcByVmidMtunnelParameters) (*PostNodeLxcByVmidMtunnelResponse, error) {

}

type GetNodeLxcByVmidMtunnelwebsocketParameters struct {
	// unix socket to forward to
	Socket string `json:"socket"`
	// ticket return by initial 'mtunnel' API call, or retrieved via 'ticket' tunnel command
	Ticket string `json:"ticket"`
}

type GetNodeLxcByVmidMtunnelwebsocketResponse struct {
}

func (c *Client) GetNodeLxcByVmidMtunnelwebsocket(Node string, Vmid int, parameters GetNodeLxcByVmidMtunnelwebsocketParameters) (*GetNodeLxcByVmidMtunnelwebsocketResponse, error) {

}

type GetNodeCephResponse struct {
}

func (c *Client) GetNodeCeph(Node string) (*GetNodeCephResponse, error) {

}

type GetNodeCephCfgResponse struct {
}

func (c *Client) GetNodeCephCfg(Node string) (*GetNodeCephCfgResponse, error) {

}

type GetNodeCephCfgRawResponse struct {
}

func (c *Client) GetNodeCephCfgRaw(Node string) (*GetNodeCephCfgRawResponse, error) {

}

type GetNodeCephCfgDbResponse struct {
}

func (c *Client) GetNodeCephCfgDb(Node string) (*GetNodeCephCfgDbResponse, error) {

}

type GetNodeCephOsdResponse struct {
}

func (c *Client) GetNodeCephOsd(Node string) (*GetNodeCephOsdResponse, error) {

}

type PostNodeCephOsdParameters struct {
	// Set the device class of the OSD in crush.
	CrushDeviceClass *string `json:"crush-device-class"`
	// Block device name for block.db.
	DbDev *string `json:"db_dev"`
	// Size in GiB for block.db.
	DbDevSize *float64 `json:"db_dev_size"`
	// Block device name.
	Dev string `json:"dev"`
	// Enables encryption of the OSD.
	Encrypted *bool `json:"encrypted"`
	// Block device name for block.wal.
	WalDev *string `json:"wal_dev"`
	// Size in GiB for block.wal.
	WalDevSize *float64 `json:"wal_dev_size"`
}

type PostNodeCephOsdResponse struct {
}

func (c *Client) PostNodeCephOsd(Node string, parameters PostNodeCephOsdParameters) (*PostNodeCephOsdResponse, error) {

}

type DeleteNodeCephOsdByOsdidParameters struct {
	// If set, we remove partition table entries.
	Cleanup *bool `json:"cleanup"`
}

type DeleteNodeCephOsdByOsdidResponse struct {
}

func (c *Client) DeleteNodeCephOsdByOsdid(Node string, Osdid int, parameters DeleteNodeCephOsdByOsdidParameters) (*DeleteNodeCephOsdByOsdidResponse, error) {

}

type GetNodeCephOsdByOsdidResponse struct {
}

func (c *Client) GetNodeCephOsdByOsdid(Node string, Osdid int) (*GetNodeCephOsdByOsdidResponse, error) {

}

type GetNodeCephOsdByOsdidMetadataResponse struct {
}

func (c *Client) GetNodeCephOsdByOsdidMetadata(Node string, Osdid int) (*GetNodeCephOsdByOsdidMetadataResponse, error) {

}

type GetNodeCephOsdByOsdidLvInfoParameters struct {
	// OSD device type
	Type *NodeCephOsdLvInfoType `json:"type"`
}

type GetNodeCephOsdByOsdidLvInfoResponse struct {
}

func (c *Client) GetNodeCephOsdByOsdidLvInfo(Node string, Osdid int, parameters GetNodeCephOsdByOsdidLvInfoParameters) (*GetNodeCephOsdByOsdidLvInfoResponse, error) {

}

type PostNodeCephOsdByOsdidInResponse struct {
}

func (c *Client) PostNodeCephOsdByOsdidIn(Node string, Osdid int) (*PostNodeCephOsdByOsdidInResponse, error) {

}

type PostNodeCephOsdByOsdidOutResponse struct {
}

func (c *Client) PostNodeCephOsdByOsdidOut(Node string, Osdid int) (*PostNodeCephOsdByOsdidOutResponse, error) {

}

type PostNodeCephOsdByOsdidScrubParameters struct {
	// If set, instructs a deep scrub instead of a normal one.
	Deep *bool `json:"deep"`
}

type PostNodeCephOsdByOsdidScrubResponse struct {
}

func (c *Client) PostNodeCephOsdByOsdidScrub(Node string, Osdid int, parameters PostNodeCephOsdByOsdidScrubParameters) (*PostNodeCephOsdByOsdidScrubResponse, error) {

}

type GetNodeCephMdsResponse struct {
}

func (c *Client) GetNodeCephMds(Node string) (*GetNodeCephMdsResponse, error) {

}

type DeleteNodeCephMdsByNameResponse struct {
}

func (c *Client) DeleteNodeCephMdsByName(Node string, Name string) (*DeleteNodeCephMdsByNameResponse, error) {

}

type PostNodeCephMdsByNameParameters struct {
	// Determines whether a ceph-mds daemon should poll and replay the log of an active MDS. Faster switch on MDS failure, but needs more idle resources.
	Hotstandby *bool `json:"hotstandby"`
}

type PostNodeCephMdsByNameResponse struct {
}

func (c *Client) PostNodeCephMdsByName(Node string, Name string, parameters PostNodeCephMdsByNameParameters) (*PostNodeCephMdsByNameResponse, error) {

}

type GetNodeCephMgrResponse struct {
}

func (c *Client) GetNodeCephMgr(Node string) (*GetNodeCephMgrResponse, error) {

}

type DeleteNodeCephMgrByIdResponse struct {
}

func (c *Client) DeleteNodeCephMgrById(Node string, Id string) (*DeleteNodeCephMgrByIdResponse, error) {

}

type PostNodeCephMgrByIdResponse struct {
}

func (c *Client) PostNodeCephMgrById(Node string, Id string) (*PostNodeCephMgrByIdResponse, error) {

}

type GetNodeCephMonResponse struct {
}

func (c *Client) GetNodeCephMon(Node string) (*GetNodeCephMonResponse, error) {

}

type DeleteNodeCephMonByMonidResponse struct {
}

func (c *Client) DeleteNodeCephMonByMonid(Node string, Monid string) (*DeleteNodeCephMonByMonidResponse, error) {

}

type PostNodeCephMonByMonidParameters struct {
	// Overwrites autodetected monitor IP address(es). Must be in the public network(s) of Ceph.
	MonAddress *string `json:"mon-address"`
}

type PostNodeCephMonByMonidResponse struct {
}

func (c *Client) PostNodeCephMonByMonid(Node string, Monid string, parameters PostNodeCephMonByMonidParameters) (*PostNodeCephMonByMonidResponse, error) {

}

type GetNodeCephFsResponse struct {
}

func (c *Client) GetNodeCephFs(Node string) (*GetNodeCephFsResponse, error) {

}

type PostNodeCephFsByNameParameters struct {
	// Configure the created CephFS as storage for this cluster.
	AddStorage *bool `json:"add-storage"`
	// Number of placement groups for the backing data pool. The metadata pool will use a quarter of this.
	PgNum *int `json:"pg_num"`
}

type PostNodeCephFsByNameResponse struct {
}

func (c *Client) PostNodeCephFsByName(Node string, Name string, parameters PostNodeCephFsByNameParameters) (*PostNodeCephFsByNameResponse, error) {

}

type GetNodeCephPoolResponse struct {
}

func (c *Client) GetNodeCephPool(Node string) (*GetNodeCephPoolResponse, error) {

}

type PostNodeCephPoolParameters struct {
	// Configure VM and CT storage using the new pool.
	AddStorages *bool `json:"add_storages"`
	// The application of the pool.
	Application *NodeCephPoolApplication `json:"application"`
	// The rule to use for mapping object placement in the cluster.
	CrushRule *string `json:"crush_rule"`
	// Create an erasure coded pool for RBD with an accompaning replicated pool for metadata storage. With EC, the common ceph options 'size', 'min_size' and 'crush_rule' parameters will be applied to the metadata pool.
	ErasureCoding *string `json:"erasure-coding"`
	// Minimum number of replicas per object
	MinSize *int `json:"min_size"`
	// The name of the pool. It must be unique.
	Name string `json:"name"`
	// The automatic PG scaling mode of the pool.
	PgAutoscaleMode *NodeCephPoolPgAutoscaleMode `json:"pg_autoscale_mode"`
	// Number of placement groups.
	PgNum *int `json:"pg_num"`
	// Minimal number of placement groups.
	PgNumMin *int `json:"pg_num_min"`
	// Number of replicas per object
	Size *int `json:"size"`
	// The estimated target size of the pool for the PG autoscaler.
	TargetSize *string `json:"target_size"`
	// The estimated target ratio of the pool for the PG autoscaler.
	TargetSizeRatio *float64 `json:"target_size_ratio"`
}

type PostNodeCephPoolResponse struct {
}

func (c *Client) PostNodeCephPool(Node string, parameters PostNodeCephPoolParameters) (*PostNodeCephPoolResponse, error) {

}

type DeleteNodeCephPoolByNameParameters struct {
	// If true, destroys pool even if in use
	Force *bool `json:"force"`
	// Remove the erasure code profile. Defaults to true, if applicable.
	RemoveEcprofile *bool `json:"remove_ecprofile"`
	// Remove all pveceph-managed storages configured for this pool
	RemoveStorages *bool `json:"remove_storages"`
}

type DeleteNodeCephPoolByNameResponse struct {
}

func (c *Client) DeleteNodeCephPoolByName(Node string, Name string, parameters DeleteNodeCephPoolByNameParameters) (*DeleteNodeCephPoolByNameResponse, error) {

}

type GetNodeCephPoolByNameResponse struct {
}

func (c *Client) GetNodeCephPoolByName(Node string, Name string) (*GetNodeCephPoolByNameResponse, error) {

}

type PutNodeCephPoolByNameParameters struct {
	// The application of the pool.
	Application *NodeCephPoolApplication `json:"application"`
	// The rule to use for mapping object placement in the cluster.
	CrushRule *string `json:"crush_rule"`
	// Minimum number of replicas per object
	MinSize *int `json:"min_size"`
	// The automatic PG scaling mode of the pool.
	PgAutoscaleMode *NodeCephPoolPgAutoscaleMode `json:"pg_autoscale_mode"`
	// Number of placement groups.
	PgNum *int `json:"pg_num"`
	// Minimal number of placement groups.
	PgNumMin *int `json:"pg_num_min"`
	// Number of replicas per object
	Size *int `json:"size"`
	// The estimated target size of the pool for the PG autoscaler.
	TargetSize *string `json:"target_size"`
	// The estimated target ratio of the pool for the PG autoscaler.
	TargetSizeRatio *float64 `json:"target_size_ratio"`
}

type PutNodeCephPoolByNameResponse struct {
}

func (c *Client) PutNodeCephPoolByName(Node string, Name string, parameters PutNodeCephPoolByNameParameters) (*PutNodeCephPoolByNameResponse, error) {

}

type GetNodeCephPoolByNameStatusParameters struct {
	// If enabled, will display additional data(eg. statistics).
	Verbose *bool `json:"verbose"`
}

type GetNodeCephPoolByNameStatusResponse struct {
}

func (c *Client) GetNodeCephPoolByNameStatus(Node string, Name string, parameters GetNodeCephPoolByNameStatusParameters) (*GetNodeCephPoolByNameStatusResponse, error) {

}

type PostNodeCephInitParameters struct {
	// Declare a separate cluster network, OSDs will routeheartbeat, object replication and recovery traffic over it
	ClusterNetwork *string `json:"cluster-network"`
	// Disable cephx authentication.
	//
	// WARNING: cephx is a security feature protecting against man-in-the-middle attacks. Only consider disabling cephx if your network is private!
	DisableCephx *bool `json:"disable_cephx"`
	// Minimum number of available replicas per object to allow I/O
	MinSize *int `json:"min_size"`
	// Use specific network for all ceph related traffic
	Network *string `json:"network"`
	// Placement group bits, used to specify the default number of placement groups.
	//
	// NOTE: 'osd pool default pg num' does not work for default pools.
	PgBits *int `json:"pg_bits"`
	// Targeted number of replicas per object
	Size *int `json:"size"`
}

type PostNodeCephInitResponse struct {
}

func (c *Client) PostNodeCephInit(Node string, parameters PostNodeCephInitParameters) (*PostNodeCephInitResponse, error) {

}

type PostNodeCephStopParameters struct {
	// Ceph service name.
	Service *string `json:"service"`
}

type PostNodeCephStopResponse struct {
}

func (c *Client) PostNodeCephStop(Node string, parameters PostNodeCephStopParameters) (*PostNodeCephStopResponse, error) {

}

type PostNodeCephStartParameters struct {
	// Ceph service name.
	Service *string `json:"service"`
}

type PostNodeCephStartResponse struct {
}

func (c *Client) PostNodeCephStart(Node string, parameters PostNodeCephStartParameters) (*PostNodeCephStartResponse, error) {

}

type PostNodeCephRestartParameters struct {
	// Ceph service name.
	Service *string `json:"service"`
}

type PostNodeCephRestartResponse struct {
}

func (c *Client) PostNodeCephRestart(Node string, parameters PostNodeCephRestartParameters) (*PostNodeCephRestartResponse, error) {

}

type GetNodeCephStatusResponse struct {
}

func (c *Client) GetNodeCephStatus(Node string) (*GetNodeCephStatusResponse, error) {

}

type GetNodeCephCrushResponse struct {
}

func (c *Client) GetNodeCephCrush(Node string) (*GetNodeCephCrushResponse, error) {

}

type GetNodeCephLogParameters struct {
	Limit *int `json:"limit"`
	Start *int `json:"start"`
}

type GetNodeCephLogResponse struct {
}

func (c *Client) GetNodeCephLog(Node string, parameters GetNodeCephLogParameters) (*GetNodeCephLogResponse, error) {

}

type GetNodeCephRulesResponse struct {
}

func (c *Client) GetNodeCephRules(Node string) (*GetNodeCephRulesResponse, error) {

}

type GetNodeCephCmdSafetyParameters struct {
	// Action to check
	Action NodeCephCmdSafetyAction `json:"action"`
	// ID of the service
	Id string `json:"id"`
	// Service type
	Service NodeCephCmdSafetyService `json:"service"`
}

type GetNodeCephCmdSafetyResponse struct {
}

func (c *Client) GetNodeCephCmdSafety(Node string, parameters GetNodeCephCmdSafetyParameters) (*GetNodeCephCmdSafetyResponse, error) {

}

type PostNodeVzdumpParameters struct {
	// Backup all known guest systems on this host.
	All *bool `json:"all"`
	// Limit I/O bandwidth (in KiB/s).
	Bwlimit *int `json:"bwlimit"`
	// Compress dump file.
	Compress *Compress `json:"compress"`
	// Store resulting files to specified directory.
	Dumpdir *string `json:"dumpdir"`
	// Exclude specified guest systems (assumes --all)
	Exclude *string `json:"exclude"`
	// Exclude certain files/directories (shell globs). Paths starting with '/' are anchored to the container's root,  other paths match relative to each subdirectory.
	ExcludePath *any `json:"exclude-path"`
	// Set IO priority when using the BFQ scheduler. For snapshot and suspend mode backups of VMs, this only affects the compressor. A value of 8 means the idle priority is used, otherwise the best-effort priority is used with the specified value.
	Ionice *int `json:"ionice"`
	// Maximal time to wait for the global lock (minutes).
	Lockwait *int `json:"lockwait"`
	// Specify when to send an email
	Mailnotification *Mailnotification `json:"mailnotification"`
	// Comma-separated list of email addresses or users that should receive email notifications.
	Mailto *string `json:"mailto"`
	// Deprecated: use 'prune-backups' instead. Maximal number of backup files per guest system.
	Maxfiles *int `json:"maxfiles"`
	// Backup mode.
	Mode *Mode `json:"mode"`
	// Template string for generating notes for the backup(s). It can contain variables which will be replaced by their values. Currently supported are {{cluster}}, {{guestname}}, {{node}}, and {{vmid}}, but more might be added in the future. Needs to be a single line, newline and backslash need to be escaped as '\n' and '\\' respectively.
	NotesTemplate *string `json:"notes-template"`
	// Other performance-related settings.
	Performance *string `json:"performance"`
	// Use pigz instead of gzip when N>0. N=1 uses half of cores, N>1 uses N as thread count.
	Pigz *int `json:"pigz"`
	// Backup all known guest systems included in the specified pool.
	Pool *string `json:"pool"`
	// If true, mark backup(s) as protected.
	Protected *bool `json:"protected"`
	// Use these retention options instead of those from the storage configuration.
	PruneBackups *string `json:"prune-backups"`
	// Be quiet.
	Quiet *bool `json:"quiet"`
	// Prune older backups according to 'prune-backups'.
	Remove *bool `json:"remove"`
	// Use specified hook script.
	Script *string `json:"script"`
	// Exclude temporary files and logs.
	Stdexcludes *bool `json:"stdexcludes"`
	// Write tar to stdout, not to a file.
	Stdout *bool `json:"stdout"`
	// Stop running backup jobs on this host.
	Stop *bool `json:"stop"`
	// Maximal time to wait until a guest system is stopped (minutes).
	Stopwait *int `json:"stopwait"`
	// Store resulting file to this storage.
	Storage *string `json:"storage"`
	// Store temporary files to specified directory.
	Tmpdir *string `json:"tmpdir"`
	// The ID of the guest system you want to backup.
	Vmid *string `json:"vmid"`
	// Zstd threads. N=0 uses half of the available cores, N>0 uses N as thread count.
	Zstd *int `json:"zstd"`
}

type PostNodeVzdumpResponse struct {
}

func (c *Client) PostNodeVzdump(Node string, parameters PostNodeVzdumpParameters) (*PostNodeVzdumpResponse, error) {

}

type GetNodeVzdumpDefaultsParameters struct {
	// The storage identifier.
	Storage *string `json:"storage"`
}

type GetNodeVzdumpDefaultsResponse struct {
}

func (c *Client) GetNodeVzdumpDefaults(Node string, parameters GetNodeVzdumpDefaultsParameters) (*GetNodeVzdumpDefaultsResponse, error) {

}

type GetNodeVzdumpExtractconfigParameters struct {
	// Volume identifier
	Volume string `json:"volume"`
}

type GetNodeVzdumpExtractconfigResponse struct {
}

func (c *Client) GetNodeVzdumpExtractconfig(Node string, parameters GetNodeVzdumpExtractconfigParameters) (*GetNodeVzdumpExtractconfigResponse, error) {

}

type GetNodeServicesResponse struct {
}

func (c *Client) GetNodeServices(Node string) (*GetNodeServicesResponse, error) {

}

type GetNodeServiceResponse struct {
}

func (c *Client) GetNodeService(Node string, Service NodeServiceService) (*GetNodeServiceResponse, error) {

}

type GetNodeServiceStateResponse struct {
}

func (c *Client) GetNodeServiceState(Node string, Service NodeServiceService) (*GetNodeServiceStateResponse, error) {

}

type PostNodeServiceStartResponse struct {
}

func (c *Client) PostNodeServiceStart(Node string, Service NodeServiceService) (*PostNodeServiceStartResponse, error) {

}

type PostNodeServiceStopResponse struct {
}

func (c *Client) PostNodeServiceStop(Node string, Service NodeServiceService) (*PostNodeServiceStopResponse, error) {

}

type PostNodeServiceRestartResponse struct {
}

func (c *Client) PostNodeServiceRestart(Node string, Service NodeServiceService) (*PostNodeServiceRestartResponse, error) {

}

type PostNodeServiceReloadResponse struct {
}

func (c *Client) PostNodeServiceReload(Node string, Service NodeServiceService) (*PostNodeServiceReloadResponse, error) {

}

type DeleteNodeSubscriptionResponse struct {
}

func (c *Client) DeleteNodeSubscription(Node string) (*DeleteNodeSubscriptionResponse, error) {

}

type GetNodeSubscriptionResponse struct {
}

func (c *Client) GetNodeSubscription(Node string) (*GetNodeSubscriptionResponse, error) {

}

type PostNodeSubscriptionParameters struct {
	// Always connect to server, even if local cache is still valid.
	Force *bool `json:"force"`
}

type PostNodeSubscriptionResponse struct {
}

func (c *Client) PostNodeSubscription(Node string, parameters PostNodeSubscriptionParameters) (*PostNodeSubscriptionResponse, error) {

}

type PutNodeSubscriptionParameters struct {
	// Proxmox VE subscription key
	Key string `json:"key"`
}

type PutNodeSubscriptionResponse struct {
}

func (c *Client) PutNodeSubscription(Node string, parameters PutNodeSubscriptionParameters) (*PutNodeSubscriptionResponse, error) {

}

type DeleteNodeNetworkResponse struct {
}

func (c *Client) DeleteNodeNetwork(Node string) (*DeleteNodeNetworkResponse, error) {

}

type GetNodeNetworkParameters struct {
	// Only list specific interface types.
	Type *NodeNetworkType `json:"type"`
}

type GetNodeNetworkResponse struct {
}

func (c *Client) GetNodeNetwork(Node string, parameters GetNodeNetworkParameters) (*GetNodeNetworkResponse, error) {

}

type PostNodeNetworkParameters struct {
	// IP address.
	Address *string `json:"address"`
	// IP address.
	Address6 *string `json:"address6"`
	// Automatically start interface on boot.
	Autostart *bool `json:"autostart"`
	// Specify the primary interface for active-backup bond.
	BondPrimary *string `json:"bond-primary"`
	// Bonding mode.
	BondMode *NodeNetworkBondMode `json:"bond_mode"`
	// Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes.
	BondXmitHashPolicy *NodeNetworkBondXmitHashPolicy `json:"bond_xmit_hash_policy"`
	// Specify the interfaces you want to add to your bridge.
	BridgePorts *string `json:"bridge_ports"`
	// Enable bridge vlan support.
	BridgeVlanAware *bool `json:"bridge_vlan_aware"`
	// IPv4 CIDR.
	Cidr *string `json:"cidr"`
	// IPv6 CIDR.
	Cidr6 *string `json:"cidr6"`
	// Comments
	Comments *string `json:"comments"`
	// Comments
	Comments6 *string `json:"comments6"`
	// Default gateway address.
	Gateway *string `json:"gateway"`
	// Default ipv6 gateway address.
	Gateway6 *string `json:"gateway6"`
	// Network interface name.
	Iface string `json:"iface"`
	// MTU.
	Mtu *int `json:"mtu"`
	// Network mask.
	Netmask *string `json:"netmask"`
	// Network mask.
	Netmask6 *int `json:"netmask6"`
	// Specify the interfaces used by the bonding device.
	OvsBonds *string `json:"ovs_bonds"`
	// The OVS bridge associated with a OVS port. This is required when you create an OVS port.
	OvsBridge *string `json:"ovs_bridge"`
	// OVS interface options.
	OvsOptions *string `json:"ovs_options"`
	// Specify the interfaces you want to add to your bridge.
	OvsPorts *string `json:"ovs_ports"`
	// Specify a VLan tag (used by OVSPort, OVSIntPort, OVSBond)
	OvsTag *int `json:"ovs_tag"`
	// Specify the interfaces used by the bonding device.
	Slaves *string `json:"slaves"`
	// Network interface type
	Type NodeNetworkParameterType `json:"type"`
	// vlan-id for a custom named vlan interface (ifupdown2 only).
	VlanId *int `json:"vlan-id"`
	// Specify the raw interface for the vlan interface.
	VlanRawDevice *string `json:"vlan-raw-device"`
}

type PostNodeNetworkResponse struct {
}

func (c *Client) PostNodeNetwork(Node string, parameters PostNodeNetworkParameters) (*PostNodeNetworkResponse, error) {

}

type PutNodeNetworkResponse struct {
}

func (c *Client) PutNodeNetwork(Node string) (*PutNodeNetworkResponse, error) {

}

type DeleteNodeNetworkByIfaceResponse struct {
}

func (c *Client) DeleteNodeNetworkByIface(Node string, Iface string) (*DeleteNodeNetworkByIfaceResponse, error) {

}

type GetNodeNetworkByIfaceResponse struct {
}

func (c *Client) GetNodeNetworkByIface(Node string, Iface string) (*GetNodeNetworkByIfaceResponse, error) {

}

type PutNodeNetworkByIfaceParameters struct {
	// IP address.
	Address *string `json:"address"`
	// IP address.
	Address6 *string `json:"address6"`
	// Automatically start interface on boot.
	Autostart *bool `json:"autostart"`
	// Specify the primary interface for active-backup bond.
	BondPrimary *string `json:"bond-primary"`
	// Bonding mode.
	BondMode *NodeNetworkBondMode `json:"bond_mode"`
	// Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes.
	BondXmitHashPolicy *NodeNetworkBondXmitHashPolicy `json:"bond_xmit_hash_policy"`
	// Specify the interfaces you want to add to your bridge.
	BridgePorts *string `json:"bridge_ports"`
	// Enable bridge vlan support.
	BridgeVlanAware *bool `json:"bridge_vlan_aware"`
	// IPv4 CIDR.
	Cidr *string `json:"cidr"`
	// IPv6 CIDR.
	Cidr6 *string `json:"cidr6"`
	// Comments
	Comments *string `json:"comments"`
	// Comments
	Comments6 *string `json:"comments6"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Default gateway address.
	Gateway *string `json:"gateway"`
	// Default ipv6 gateway address.
	Gateway6 *string `json:"gateway6"`
	// MTU.
	Mtu *int `json:"mtu"`
	// Network mask.
	Netmask *string `json:"netmask"`
	// Network mask.
	Netmask6 *int `json:"netmask6"`
	// Specify the interfaces used by the bonding device.
	OvsBonds *string `json:"ovs_bonds"`
	// The OVS bridge associated with a OVS port. This is required when you create an OVS port.
	OvsBridge *string `json:"ovs_bridge"`
	// OVS interface options.
	OvsOptions *string `json:"ovs_options"`
	// Specify the interfaces you want to add to your bridge.
	OvsPorts *string `json:"ovs_ports"`
	// Specify a VLan tag (used by OVSPort, OVSIntPort, OVSBond)
	OvsTag *int `json:"ovs_tag"`
	// Specify the interfaces used by the bonding device.
	Slaves *string `json:"slaves"`
	// Network interface type
	Type NodeNetworkParameterType `json:"type"`
	// vlan-id for a custom named vlan interface (ifupdown2 only).
	VlanId *int `json:"vlan-id"`
	// Specify the raw interface for the vlan interface.
	VlanRawDevice *string `json:"vlan-raw-device"`
}

type PutNodeNetworkByIfaceResponse struct {
}

func (c *Client) PutNodeNetworkByIface(Node string, Iface string, parameters PutNodeNetworkByIfaceParameters) (*PutNodeNetworkByIfaceResponse, error) {

}

type GetNodeTasksParameters struct {
	// Only list tasks with a status of ERROR.
	Errors *bool `json:"errors"`
	// Only list this amount of tasks.
	Limit *int `json:"limit"`
	// Only list tasks since this UNIX epoch.
	Since *int `json:"since"`
	// List archived, active or all tasks.
	Source *NodeTasksSource `json:"source"`
	// List tasks beginning from this offset.
	Start *int `json:"start"`
	// List of Task States that should be returned.
	Statusfilter *string `json:"statusfilter"`
	// Only list tasks of this type (e.g., vzstart, vzdump).
	Typefilter *string `json:"typefilter"`
	// Only list tasks until this UNIX epoch.
	Until *int `json:"until"`
	// Only list tasks from this user.
	Userfilter *string `json:"userfilter"`
	// Only list tasks for this VM.
	Vmid *int `json:"vmid"`
}

type GetNodeTasksResponse struct {
}

func (c *Client) GetNodeTasks(Node string, parameters GetNodeTasksParameters) (*GetNodeTasksResponse, error) {

}

type DeleteNodeTaskResponse struct {
}

func (c *Client) DeleteNodeTask(Node string, Upid string) (*DeleteNodeTaskResponse, error) {

}

type GetNodeTaskResponse struct {
}

func (c *Client) GetNodeTask(Node string, Upid string) (*GetNodeTaskResponse, error) {

}

type GetNodeTaskLogParameters struct {
	// Whether the tasklog file should be downloaded. This parameter can't be used in conjunction with other parameters
	Download *bool `json:"download"`
	// The amount of lines to read from the tasklog.
	Limit *int `json:"limit"`
	// Start at this line when reading the tasklog
	Start *int `json:"start"`
}

type GetNodeTaskLogResponse struct {
}

func (c *Client) GetNodeTaskLog(Node string, Upid string, parameters GetNodeTaskLogParameters) (*GetNodeTaskLogResponse, error) {

}

type GetNodeTaskStatusResponse struct {
}

func (c *Client) GetNodeTaskStatus(Node string, Upid string) (*GetNodeTaskStatusResponse, error) {

}

type GetNodeScanResponse struct {
}

func (c *Client) GetNodeScan(Node string) (*GetNodeScanResponse, error) {

}

type GetNodeScanNfsParameters struct {
	// The server address (name or IP).
	Server string `json:"server"`
}

type GetNodeScanNfsResponse struct {
}

func (c *Client) GetNodeScanNfs(Node string, parameters GetNodeScanNfsParameters) (*GetNodeScanNfsResponse, error) {

}

type GetNodeScanCifsParameters struct {
	// SMB domain (Workgroup).
	Domain *string `json:"domain"`
	// User password.
	Password *string `json:"password"`
	// The server address (name or IP).
	Server string `json:"server"`
	// User name.
	Username *string `json:"username"`
}

type GetNodeScanCifsResponse struct {
}

func (c *Client) GetNodeScanCifs(Node string, parameters GetNodeScanCifsParameters) (*GetNodeScanCifsResponse, error) {

}

type GetNodeScanPbsParameters struct {
	// Certificate SHA 256 fingerprint.
	Fingerprint *string `json:"fingerprint"`
	// User password or API token secret.
	Password string `json:"password"`
	// Optional port.
	Port *int `json:"port"`
	// The server address (name or IP).
	Server string `json:"server"`
	// User-name or API token-ID.
	Username string `json:"username"`
}

type GetNodeScanPbsResponse struct {
}

func (c *Client) GetNodeScanPbs(Node string, parameters GetNodeScanPbsParameters) (*GetNodeScanPbsResponse, error) {

}

type GetNodeScanGlusterfsParameters struct {
	// The server address (name or IP).
	Server string `json:"server"`
}

type GetNodeScanGlusterfsResponse struct {
}

func (c *Client) GetNodeScanGlusterfs(Node string, parameters GetNodeScanGlusterfsParameters) (*GetNodeScanGlusterfsResponse, error) {

}

type GetNodeScanIscsiParameters struct {
	// The iSCSI portal (IP or DNS name with optional port).
	Portal string `json:"portal"`
}

type GetNodeScanIscsiResponse struct {
}

func (c *Client) GetNodeScanIscsi(Node string, parameters GetNodeScanIscsiParameters) (*GetNodeScanIscsiResponse, error) {

}

type GetNodeScanLvmResponse struct {
}

func (c *Client) GetNodeScanLvm(Node string) (*GetNodeScanLvmResponse, error) {

}

type GetNodeScanLvmthinParameters struct {
	Vg string `json:"vg"`
}

type GetNodeScanLvmthinResponse struct {
}

func (c *Client) GetNodeScanLvmthin(Node string, parameters GetNodeScanLvmthinParameters) (*GetNodeScanLvmthinResponse, error) {

}

type GetNodeScanZfsResponse struct {
}

func (c *Client) GetNodeScanZfs(Node string) (*GetNodeScanZfsResponse, error) {

}

type GetNodeHardwareResponse struct {
}

func (c *Client) GetNodeHardware(Node string) (*GetNodeHardwareResponse, error) {

}

type GetNodeHardwarePciParameters struct {
	// A list of blacklisted PCI classes, which will not be returned. Following are filtered by default: Memory Controller (05), Bridge (06) and Processor (0b).
	PciClassBlacklist *string `json:"pci-class-blacklist"`
	// If disabled, does only print the PCI IDs. Otherwise, additional information like vendor and device will be returned.
	Verbose *bool `json:"verbose"`
}

type GetNodeHardwarePciResponse struct {
}

func (c *Client) GetNodeHardwarePci(Node string, parameters GetNodeHardwarePciParameters) (*GetNodeHardwarePciResponse, error) {

}

type GetNodeHardwarePciByPciidResponse struct {
}

func (c *Client) GetNodeHardwarePciByPciid(Node string, Pciid string) (*GetNodeHardwarePciByPciidResponse, error) {

}

type GetNodeHardwarePciByPciidMdevResponse struct {
}

func (c *Client) GetNodeHardwarePciByPciidMdev(Node string, Pciid string) (*GetNodeHardwarePciByPciidMdevResponse, error) {

}

type GetNodeHardwareUsbResponse struct {
}

func (c *Client) GetNodeHardwareUsb(Node string) (*GetNodeHardwareUsbResponse, error) {

}

type GetNodeCapabilitiesResponse struct {
}

func (c *Client) GetNodeCapabilities(Node string) (*GetNodeCapabilitiesResponse, error) {

}

type GetNodeCapabilitiesQemuResponse struct {
}

func (c *Client) GetNodeCapabilitiesQemu(Node string) (*GetNodeCapabilitiesQemuResponse, error) {

}

type GetNodeCapabilitiesQemuCpuResponse struct {
}

func (c *Client) GetNodeCapabilitiesQemuCpu(Node string) (*GetNodeCapabilitiesQemuCpuResponse, error) {

}

type GetNodeCapabilitiesQemuMachinesResponse struct {
}

func (c *Client) GetNodeCapabilitiesQemuMachines(Node string) (*GetNodeCapabilitiesQemuMachinesResponse, error) {

}

type GetNodeStorageParameters struct {
	// Only list stores which support this content type.
	Content *string `json:"content"`
	// Only list stores which are enabled (not disabled in config).
	Enabled *bool `json:"enabled"`
	// Include information about formats
	Format *bool `json:"format"`
	// Only list status for  specified storage
	Storage *string `json:"storage"`
	// If target is different to 'node', we only lists shared storages which content is accessible on this 'node' and the specified 'target' node.
	Target *string `json:"target"`
}

type GetNodeStorageResponse struct {
}

func (c *Client) GetNodeStorage(Node string, parameters GetNodeStorageParameters) (*GetNodeStorageResponse, error) {

}

type GetNodeStorageByStorageResponse struct {
}

func (c *Client) GetNodeStorageByStorage(Node string, Storage string) (*GetNodeStorageByStorageResponse, error) {

}

type DeleteNodeStorageByStoragePrunebackupsParameters struct {
	// Use these retention options instead of those from the storage configuration.
	PruneBackups *string `json:"prune-backups"`
	// Either 'qemu' or 'lxc'. Only consider backups for guests of this type.
	Type *NodeStoragePrunebackupsType `json:"type"`
	// Only prune backups for this VM.
	Vmid *int `json:"vmid"`
}

type DeleteNodeStorageByStoragePrunebackupsResponse struct {
}

func (c *Client) DeleteNodeStorageByStoragePrunebackups(Node string, Storage string, parameters DeleteNodeStorageByStoragePrunebackupsParameters) (*DeleteNodeStorageByStoragePrunebackupsResponse, error) {

}

type GetNodeStorageByStoragePrunebackupsParameters struct {
	// Use these retention options instead of those from the storage configuration.
	PruneBackups *string `json:"prune-backups"`
	// Either 'qemu' or 'lxc'. Only consider backups for guests of this type.
	Type *NodeStoragePrunebackupsType `json:"type"`
	// Only consider backups for this guest.
	Vmid *int `json:"vmid"`
}

type GetNodeStorageByStoragePrunebackupsResponse struct {
}

func (c *Client) GetNodeStorageByStoragePrunebackups(Node string, Storage string, parameters GetNodeStorageByStoragePrunebackupsParameters) (*GetNodeStorageByStoragePrunebackupsResponse, error) {

}

type GetNodeStorageByStorageContentParameters struct {
	// Only list content of this type.
	Content *string `json:"content"`
	// Only list images for this VM
	Vmid *int `json:"vmid"`
}

type GetNodeStorageByStorageContentResponse struct {
}

func (c *Client) GetNodeStorageByStorageContent(Node string, Storage string, parameters GetNodeStorageByStorageContentParameters) (*GetNodeStorageByStorageContentResponse, error) {

}

type PostNodeStorageByStorageContentParameters struct {
	// The name of the file to create.
	Filename string                    `json:"filename"`
	Format   *NodeStorageContentFormat `json:"format"`
	// Size in kilobyte (1024 bytes). Optional suffixes 'M' (megabyte, 1024K) and 'G' (gigabyte, 1024M)
	Size string `json:"size"`
	// Specify owner VM
	Vmid int `json:"vmid"`
}

type PostNodeStorageByStorageContentResponse struct {
}

func (c *Client) PostNodeStorageByStorageContent(Node string, Storage string, parameters PostNodeStorageByStorageContentParameters) (*PostNodeStorageByStorageContentResponse, error) {

}

type DeleteNodeStorageByStorageContentByVolumeParameters struct {
	// Time to wait for the task to finish. We return 'null' if the task finish within that time.
	Delay *int `json:"delay"`
}

type DeleteNodeStorageByStorageContentByVolumeResponse struct {
}

func (c *Client) DeleteNodeStorageByStorageContentByVolume(Node string, Storage string, Volume string, parameters DeleteNodeStorageByStorageContentByVolumeParameters) (*DeleteNodeStorageByStorageContentByVolumeResponse, error) {

}

type GetNodeStorageByStorageContentByVolumeResponse struct {
}

func (c *Client) GetNodeStorageByStorageContentByVolume(Node string, Storage string, Volume string) (*GetNodeStorageByStorageContentByVolumeResponse, error) {

}

type PostNodeStorageByStorageContentByVolumeParameters struct {
	// Target volume identifier
	Target string `json:"target"`
	// Target node. Default is local node.
	TargetNode *string `json:"target_node"`
}

type PostNodeStorageByStorageContentByVolumeResponse struct {
}

func (c *Client) PostNodeStorageByStorageContentByVolume(Node string, Storage string, Volume string, parameters PostNodeStorageByStorageContentByVolumeParameters) (*PostNodeStorageByStorageContentByVolumeResponse, error) {

}

type PutNodeStorageByStorageContentByVolumeParameters struct {
	// The new notes.
	Notes *string `json:"notes"`
	// Protection status. Currently only supported for backups.
	Protected *bool `json:"protected"`
}

type PutNodeStorageByStorageContentByVolumeResponse struct {
}

func (c *Client) PutNodeStorageByStorageContentByVolume(Node string, Storage string, Volume string, parameters PutNodeStorageByStorageContentByVolumeParameters) (*PutNodeStorageByStorageContentByVolumeResponse, error) {

}

type GetNodeStorageByStorageFileRestoreListParameters struct {
	// base64-path to the directory or file being listed, or "/".
	Filepath string `json:"filepath"`
	// Backup volume ID or name. Currently only PBS snapshots are supported.
	Volume string `json:"volume"`
}

type GetNodeStorageByStorageFileRestoreListResponse struct {
}

func (c *Client) GetNodeStorageByStorageFileRestoreList(Node string, Storage string, parameters GetNodeStorageByStorageFileRestoreListParameters) (*GetNodeStorageByStorageFileRestoreListResponse, error) {

}

type GetNodeStorageByStorageFileRestoreDownloadParameters struct {
	// base64-path to the directory or file to download.
	Filepath string `json:"filepath"`
	// Backup volume ID or name. Currently only PBS snapshots are supported.
	Volume string `json:"volume"`
}

type GetNodeStorageByStorageFileRestoreDownloadResponse struct {
}

func (c *Client) GetNodeStorageByStorageFileRestoreDownload(Node string, Storage string, parameters GetNodeStorageByStorageFileRestoreDownloadParameters) (*GetNodeStorageByStorageFileRestoreDownloadResponse, error) {

}

type GetNodeStorageByStorageStatusResponse struct {
}

func (c *Client) GetNodeStorageByStorageStatus(Node string, Storage string) (*GetNodeStorageByStorageStatusResponse, error) {

}

type GetNodeStorageByStorageRrdParameters struct {
	// The RRD consolidation function
	Cf *NodeRrdCf `json:"cf"`
	// The list of datasources you want to display.
	Ds string `json:"ds"`
	// Specify the time frame you are interested in.
	Timeframe NodeRrdTimeframe `json:"timeframe"`
}

type GetNodeStorageByStorageRrdResponse struct {
}

func (c *Client) GetNodeStorageByStorageRrd(Node string, Storage string, parameters GetNodeStorageByStorageRrdParameters) (*GetNodeStorageByStorageRrdResponse, error) {

}

type GetNodeStorageByStorageRrddataParameters struct {
	// The RRD consolidation function
	Cf *NodeRrdCf `json:"cf"`
	// Specify the time frame you are interested in.
	Timeframe NodeRrdTimeframe `json:"timeframe"`
}

type GetNodeStorageByStorageRrddataResponse struct {
}

func (c *Client) GetNodeStorageByStorageRrddata(Node string, Storage string, parameters GetNodeStorageByStorageRrddataParameters) (*GetNodeStorageByStorageRrddataResponse, error) {

}

type PostNodeStorageByStorageUploadParameters struct {
	// The expected checksum of the file.
	Checksum *string `json:"checksum"`
	// The algorithm to calculate the checksum of the file.
	ChecksumAlgorithm *NodeStorageChecksumAlgorithm `json:"checksum-algorithm"`
	// Content type.
	Content NodeStorageContent `json:"content"`
	// The name of the file to create. Caution: This will be normalized!
	Filename string `json:"filename"`
	// The source file name. This parameter is usually set by the REST handler. You can only overwrite it when connecting to the trusted port on localhost.
	Tmpfilename *string `json:"tmpfilename"`
}

type PostNodeStorageByStorageUploadResponse struct {
}

func (c *Client) PostNodeStorageByStorageUpload(Node string, Storage string, parameters PostNodeStorageByStorageUploadParameters) (*PostNodeStorageByStorageUploadResponse, error) {

}

type PostNodeStorageByStorageDownloadUrlParameters struct {
	// The expected checksum of the file.
	Checksum *string `json:"checksum"`
	// The algorithm to calculate the checksum of the file.
	ChecksumAlgorithm *NodeStorageChecksumAlgorithm `json:"checksum-algorithm"`
	// Content type.
	Content NodeStorageContent `json:"content"`
	// The name of the file to create. Caution: This will be normalized!
	Filename string `json:"filename"`
	// The URL to download the file from.
	Url string `json:"url"`
	// If false, no SSL/TLS certificates will be verified.
	VerifyCertificates *bool `json:"verify-certificates"`
}

type PostNodeStorageByStorageDownloadUrlResponse struct {
}

func (c *Client) PostNodeStorageByStorageDownloadUrl(Node string, Storage string, parameters PostNodeStorageByStorageDownloadUrlParameters) (*PostNodeStorageByStorageDownloadUrlResponse, error) {

}

type GetNodeDisksResponse struct {
}

func (c *Client) GetNodeDisks(Node string) (*GetNodeDisksResponse, error) {

}

type GetNodeDisksLvmResponse struct {
}

func (c *Client) GetNodeDisksLvm(Node string) (*GetNodeDisksLvmResponse, error) {

}

type PostNodeDisksLvmParameters struct {
	// Configure storage using the Volume Group
	AddStorage *bool `json:"add_storage"`
	// The block device you want to create the volume group on
	Device string `json:"device"`
	// The storage identifier.
	Name string `json:"name"`
}

type PostNodeDisksLvmResponse struct {
}

func (c *Client) PostNodeDisksLvm(Node string, parameters PostNodeDisksLvmParameters) (*PostNodeDisksLvmResponse, error) {

}

type DeleteNodeDisksLvmByNameParameters struct {
	// Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupConfig *bool `json:"cleanup-config"`
	// Also wipe disks so they can be repurposed afterwards.
	CleanupDisks *bool `json:"cleanup-disks"`
}

type DeleteNodeDisksLvmByNameResponse struct {
}

func (c *Client) DeleteNodeDisksLvmByName(Node string, Name string, parameters DeleteNodeDisksLvmByNameParameters) (*DeleteNodeDisksLvmByNameResponse, error) {

}

type GetNodeDisksLvmthinResponse struct {
}

func (c *Client) GetNodeDisksLvmthin(Node string) (*GetNodeDisksLvmthinResponse, error) {

}

type PostNodeDisksLvmthinParameters struct {
	// Configure storage using the thinpool.
	AddStorage *bool `json:"add_storage"`
	// The block device you want to create the thinpool on.
	Device string `json:"device"`
	// The storage identifier.
	Name string `json:"name"`
}

type PostNodeDisksLvmthinResponse struct {
}

func (c *Client) PostNodeDisksLvmthin(Node string, parameters PostNodeDisksLvmthinParameters) (*PostNodeDisksLvmthinResponse, error) {

}

type DeleteNodeDisksLvmthinByNameParameters struct {
	// Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupConfig *bool `json:"cleanup-config"`
	// Also wipe disks so they can be repurposed afterwards.
	CleanupDisks *bool `json:"cleanup-disks"`
	// The storage identifier.
	VolumeGroup string `json:"volume-group"`
}

type DeleteNodeDisksLvmthinByNameResponse struct {
}

func (c *Client) DeleteNodeDisksLvmthinByName(Node string, Name string, parameters DeleteNodeDisksLvmthinByNameParameters) (*DeleteNodeDisksLvmthinByNameResponse, error) {

}

type GetNodeDisksDirectoryResponse struct {
}

func (c *Client) GetNodeDisksDirectory(Node string) (*GetNodeDisksDirectoryResponse, error) {

}

type PostNodeDisksDirectoryParameters struct {
	// Configure storage using the directory.
	AddStorage *bool `json:"add_storage"`
	// The block device you want to create the filesystem on.
	Device string `json:"device"`
	// The desired filesystem.
	Filesystem *NodeDisksDirectoryFilesystem `json:"filesystem"`
	// The storage identifier.
	Name string `json:"name"`
}

type PostNodeDisksDirectoryResponse struct {
}

func (c *Client) PostNodeDisksDirectory(Node string, parameters PostNodeDisksDirectoryParameters) (*PostNodeDisksDirectoryResponse, error) {

}

type DeleteNodeDisksDirectoryByNameParameters struct {
	// Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupConfig *bool `json:"cleanup-config"`
	// Also wipe disk so it can be repurposed afterwards.
	CleanupDisks *bool `json:"cleanup-disks"`
}

type DeleteNodeDisksDirectoryByNameResponse struct {
}

func (c *Client) DeleteNodeDisksDirectoryByName(Node string, Name string, parameters DeleteNodeDisksDirectoryByNameParameters) (*DeleteNodeDisksDirectoryByNameResponse, error) {

}

type GetNodeDisksZfsResponse struct {
}

func (c *Client) GetNodeDisksZfs(Node string) (*GetNodeDisksZfsResponse, error) {

}

type PostNodeDisksZfsParameters struct {
	// Configure storage using the zpool.
	AddStorage *bool `json:"add_storage"`
	// Pool sector size exponent.
	Ashift *int `json:"ashift"`
	// The compression algorithm to use.
	Compression *NodeDisksZfsCompression `json:"compression"`
	// The block devices you want to create the zpool on.
	Devices     string  `json:"devices"`
	DraidConfig *string `json:"draid-config"`
	// The storage identifier.
	Name string `json:"name"`
	// The RAID level to use.
	Raidlevel NodeDisksZfsRaidlevel `json:"raidlevel"`
}

type PostNodeDisksZfsResponse struct {
}

func (c *Client) PostNodeDisksZfs(Node string, parameters PostNodeDisksZfsParameters) (*PostNodeDisksZfsResponse, error) {

}

type DeleteNodeDisksZfsByNameParameters struct {
	// Marks associated storage(s) as not available on this node anymore or removes them from the configuration (if configured for this node only).
	CleanupConfig *bool `json:"cleanup-config"`
	// Also wipe disks so they can be repurposed afterwards.
	CleanupDisks *bool `json:"cleanup-disks"`
}

type DeleteNodeDisksZfsByNameResponse struct {
}

func (c *Client) DeleteNodeDisksZfsByName(Node string, Name string, parameters DeleteNodeDisksZfsByNameParameters) (*DeleteNodeDisksZfsByNameResponse, error) {

}

type GetNodeDisksZfsByNameResponse struct {
}

func (c *Client) GetNodeDisksZfsByName(Node string, Name string) (*GetNodeDisksZfsByNameResponse, error) {

}

type GetNodeDisksListParameters struct {
	// Also include partitions.
	IncludePartitions *bool `json:"include-partitions"`
	// Skip smart checks.
	Skipsmart *bool `json:"skipsmart"`
	// Only list specific types of disks.
	Type *NodeDisksListType `json:"type"`
}

type GetNodeDisksListResponse struct {
}

func (c *Client) GetNodeDisksList(Node string, parameters GetNodeDisksListParameters) (*GetNodeDisksListResponse, error) {

}

type GetNodeDisksSmartParameters struct {
	// Block device name
	Disk string `json:"disk"`
	// If true returns only the health status
	Healthonly *bool `json:"healthonly"`
}

type GetNodeDisksSmartResponse struct {
}

func (c *Client) GetNodeDisksSmart(Node string, parameters GetNodeDisksSmartParameters) (*GetNodeDisksSmartResponse, error) {

}

type PostNodeDisksInitgptParameters struct {
	// Block device name
	Disk string `json:"disk"`
	// UUID for the GPT table
	Uuid *string `json:"uuid"`
}

type PostNodeDisksInitgptResponse struct {
}

func (c *Client) PostNodeDisksInitgpt(Node string, parameters PostNodeDisksInitgptParameters) (*PostNodeDisksInitgptResponse, error) {

}

type PutNodeDisksWipediskParameters struct {
	// Block device name
	Disk string `json:"disk"`
}

type PutNodeDisksWipediskResponse struct {
}

func (c *Client) PutNodeDisksWipedisk(Node string, parameters PutNodeDisksWipediskParameters) (*PutNodeDisksWipediskResponse, error) {

}

type GetNodeAptResponse struct {
}

func (c *Client) GetNodeApt(Node string) (*GetNodeAptResponse, error) {

}

type GetNodeAptUpdateResponse struct {
}

func (c *Client) GetNodeAptUpdate(Node string) (*GetNodeAptUpdateResponse, error) {

}

type PostNodeAptUpdateParameters struct {
	// Send notification mail about new packages (to email address specified for user 'root@pam').
	Notify *bool `json:"notify"`
	// Only produces output suitable for logging, omitting progress indicators.
	Quiet *bool `json:"quiet"`
}

type PostNodeAptUpdateResponse struct {
}

func (c *Client) PostNodeAptUpdate(Node string, parameters PostNodeAptUpdateParameters) (*PostNodeAptUpdateResponse, error) {

}

type GetNodeAptChangelogParameters struct {
	// Package name.
	Name string `json:"name"`
	// Package version.
	Version *string `json:"version"`
}

type GetNodeAptChangelogResponse struct {
}

func (c *Client) GetNodeAptChangelog(Node string, parameters GetNodeAptChangelogParameters) (*GetNodeAptChangelogResponse, error) {

}

type GetNodeAptRepositoriesResponse struct {
}

func (c *Client) GetNodeAptRepositories(Node string) (*GetNodeAptRepositoriesResponse, error) {

}

type PostNodeAptRepositoriesParameters struct {
	// Digest to detect modifications.
	Digest *string `json:"digest"`
	// Whether the repository should be enabled or not.
	Enabled *bool `json:"enabled"`
	// Index within the file (starting from 0).
	Index int `json:"index"`
	// Path to the containing file.
	Path string `json:"path"`
}

type PostNodeAptRepositoriesResponse struct {
}

func (c *Client) PostNodeAptRepositories(Node string, parameters PostNodeAptRepositoriesParameters) (*PostNodeAptRepositoriesResponse, error) {

}

type PutNodeAptRepositoriesParameters struct {
	// Digest to detect modifications.
	Digest *string `json:"digest"`
	// Handle that identifies a repository.
	Handle string `json:"handle"`
}

type PutNodeAptRepositoriesResponse struct {
}

func (c *Client) PutNodeAptRepositories(Node string, parameters PutNodeAptRepositoriesParameters) (*PutNodeAptRepositoriesResponse, error) {

}

type GetNodeAptVersionsResponse struct {
}

func (c *Client) GetNodeAptVersions(Node string) (*GetNodeAptVersionsResponse, error) {

}

type GetNodeFirewallResponse struct {
}

func (c *Client) GetNodeFirewall(Node string) (*GetNodeFirewallResponse, error) {

}

type GetNodeFirewallRulesResponse struct {
}

func (c *Client) GetNodeFirewallRules(Node string) (*GetNodeFirewallRulesResponse, error) {

}

type PostNodeFirewallRulesParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Update rule at position <pos>.
	Pos *int `json:"pos"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type FirewallType `json:"type"`
}

type PostNodeFirewallRulesResponse struct {
}

func (c *Client) PostNodeFirewallRules(Node string, parameters PostNodeFirewallRulesParameters) (*PostNodeFirewallRulesResponse, error) {

}

type DeleteNodeFirewallRuleParameters struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type DeleteNodeFirewallRuleResponse struct {
}

func (c *Client) DeleteNodeFirewallRule(Node string, Pos int, parameters DeleteNodeFirewallRuleParameters) (*DeleteNodeFirewallRuleResponse, error) {

}

type GetNodeFirewallRuleResponse struct {
}

func (c *Client) GetNodeFirewallRule(Node string, Pos int) (*GetNodeFirewallRuleResponse, error) {

}

type PutNodeFirewallRuleParameters struct {
	// Rule action ('ACCEPT', 'DROP', 'REJECT') or security group name.
	Action *string `json:"action"`
	// Descriptive comment.
	Comment *string `json:"comment"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Restrict packet destination address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Dest *string `json:"dest"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Restrict TCP/UDP destination port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Dport *string `json:"dport"`
	// Flag to enable/disable a rule.
	Enable *int `json:"enable"`
	// Specify icmp-type. Only valid if proto equals 'icmp' or 'icmpv6'/'ipv6-icmp'.
	IcmpType *string `json:"icmp-type"`
	// Network interface name. You have to use network configuration key names for VMs and containers ('net\d+'). Host related rules can use arbitrary strings.
	Iface *string `json:"iface"`
	// Log level for firewall rule.
	Log *FirewallLog `json:"log"`
	// Use predefined standard macro.
	Macro *string `json:"macro"`
	// Move rule to new position <moveto>. Other arguments are ignored.
	Moveto *int `json:"moveto"`
	// IP protocol. You can use protocol names ('tcp'/'udp') or simple numbers, as defined in '/etc/protocols'.
	Proto *string `json:"proto"`
	// Restrict packet source address. This can refer to a single IP address, an IP set ('+ipsetname') or an IP alias definition. You can also specify an address range like '20.34.101.207-201.3.9.99', or a list of IP addresses and networks (entries are separated by comma). Please do not mix IPv4 and IPv6 addresses inside such lists.
	Source *string `json:"source"`
	// Restrict TCP/UDP source port. You can use service names or simple numbers (0-65535), as defined in '/etc/services'. Port ranges can be specified with '\d+:\d+', for example '80:85', and you can use comma separated list to match several ports or ranges.
	Sport *string `json:"sport"`
	// Rule type.
	Type *FirewallType `json:"type"`
}

type PutNodeFirewallRuleResponse struct {
}

func (c *Client) PutNodeFirewallRule(Node string, Pos int, parameters PutNodeFirewallRuleParameters) (*PutNodeFirewallRuleResponse, error) {

}

type GetNodeFirewallOptionsResponse struct {
}

func (c *Client) GetNodeFirewallOptions(Node string) (*GetNodeFirewallOptionsResponse, error) {

}

type PutNodeFirewallOptionsParameters struct {
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Enable host firewall rules.
	Enable *bool `json:"enable"`
	// Log level for incoming traffic.
	LogLevelIn *FirewallLog `json:"log_level_in"`
	// Log level for outgoing traffic.
	LogLevelOut *FirewallLog `json:"log_level_out"`
	// Enable logging of conntrack information.
	LogNfConntrack *bool `json:"log_nf_conntrack"`
	// Enable NDP (Neighbor Discovery Protocol).
	Ndp *bool `json:"ndp"`
	// Allow invalid packets on connection tracking.
	NfConntrackAllowInvalid *bool `json:"nf_conntrack_allow_invalid"`
	// Enable conntrack helpers for specific protocols. Supported protocols: amanda, ftp, irc, netbios-ns, pptp, sane, sip, snmp, tftp
	NfConntrackHelpers *string `json:"nf_conntrack_helpers"`
	// Maximum number of tracked connections.
	NfConntrackMax *int `json:"nf_conntrack_max"`
	// Conntrack established timeout.
	NfConntrackTcpTimeoutEstablished *int `json:"nf_conntrack_tcp_timeout_established"`
	// Conntrack syn recv timeout.
	NfConntrackTcpTimeoutSynRecv *int `json:"nf_conntrack_tcp_timeout_syn_recv"`
	// Enable SMURFS filter.
	Nosmurfs *bool `json:"nosmurfs"`
	// Enable synflood protection
	ProtectionSynflood *bool `json:"protection_synflood"`
	// Synflood protection rate burst by ip src.
	ProtectionSynfloodBurst *int `json:"protection_synflood_burst"`
	// Synflood protection rate syn/sec by ip src.
	ProtectionSynfloodRate *int `json:"protection_synflood_rate"`
	// Log level for SMURFS filter.
	SmurfLogLevel *FirewallLog `json:"smurf_log_level"`
	// Log level for illegal tcp flags filter.
	TcpFlagsLogLevel *FirewallLog `json:"tcp_flags_log_level"`
	// Filter illegal combinations of TCP flags.
	Tcpflags *bool `json:"tcpflags"`
}

type PutNodeFirewallOptionsResponse struct {
}

func (c *Client) PutNodeFirewallOptions(Node string, parameters PutNodeFirewallOptionsParameters) (*PutNodeFirewallOptionsResponse, error) {

}

type GetNodeFirewallLogParameters struct {
	Limit *int `json:"limit"`
	// Display log since this UNIX epoch.
	Since *int `json:"since"`
	Start *int `json:"start"`
	// Display log until this UNIX epoch.
	Until *int `json:"until"`
}

type GetNodeFirewallLogResponse struct {
}

func (c *Client) GetNodeFirewallLog(Node string, parameters GetNodeFirewallLogParameters) (*GetNodeFirewallLogResponse, error) {

}

type GetNodeReplicationParameters struct {
	// Only list replication jobs for this guest.
	Guest *int `json:"guest"`
}

type GetNodeReplicationResponse struct {
}

func (c *Client) GetNodeReplication(Node string, parameters GetNodeReplicationParameters) (*GetNodeReplicationResponse, error) {

}

type GetNodeReplicationByIdResponse struct {
}

func (c *Client) GetNodeReplicationById(Node string, Id string) (*GetNodeReplicationByIdResponse, error) {

}

type GetNodeReplicationByIdStatusResponse struct {
}

func (c *Client) GetNodeReplicationByIdStatus(Node string, Id string) (*GetNodeReplicationByIdStatusResponse, error) {

}

type GetNodeReplicationByIdLogParameters struct {
	Limit *int `json:"limit"`
	Start *int `json:"start"`
}

type GetNodeReplicationByIdLogResponse struct {
}

func (c *Client) GetNodeReplicationByIdLog(Node string, Id string, parameters GetNodeReplicationByIdLogParameters) (*GetNodeReplicationByIdLogResponse, error) {

}

type PostNodeReplicationByIdScheduleNowResponse struct {
}

func (c *Client) PostNodeReplicationByIdScheduleNow(Node string, Id string) (*PostNodeReplicationByIdScheduleNowResponse, error) {

}

type GetNodeCertificatesResponse struct {
}

func (c *Client) GetNodeCertificates(Node string) (*GetNodeCertificatesResponse, error) {

}

type GetNodeCertificatesAcmeResponse struct {
}

func (c *Client) GetNodeCertificatesAcme(Node string) (*GetNodeCertificatesAcmeResponse, error) {

}

type DeleteNodeCertificatesAcmeCertificateResponse struct {
}

func (c *Client) DeleteNodeCertificatesAcmeCertificate(Node string) (*DeleteNodeCertificatesAcmeCertificateResponse, error) {

}

type PostNodeCertificatesAcmeCertificateParameters struct {
	// Overwrite existing custom certificate.
	Force *bool `json:"force"`
}

type PostNodeCertificatesAcmeCertificateResponse struct {
}

func (c *Client) PostNodeCertificatesAcmeCertificate(Node string, parameters PostNodeCertificatesAcmeCertificateParameters) (*PostNodeCertificatesAcmeCertificateResponse, error) {

}

type PutNodeCertificatesAcmeCertificateParameters struct {
	// Force renewal even if expiry is more than 30 days away.
	Force *bool `json:"force"`
}

type PutNodeCertificatesAcmeCertificateResponse struct {
}

func (c *Client) PutNodeCertificatesAcmeCertificate(Node string, parameters PutNodeCertificatesAcmeCertificateParameters) (*PutNodeCertificatesAcmeCertificateResponse, error) {

}

type GetNodeCertificatesInfoResponse struct {
}

func (c *Client) GetNodeCertificatesInfo(Node string) (*GetNodeCertificatesInfoResponse, error) {

}

type DeleteNodeCertificatesCustomParameters struct {
	// Restart pveproxy.
	Restart *bool `json:"restart"`
}

type DeleteNodeCertificatesCustomResponse struct {
}

func (c *Client) DeleteNodeCertificatesCustom(Node string, parameters DeleteNodeCertificatesCustomParameters) (*DeleteNodeCertificatesCustomResponse, error) {

}

type PostNodeCertificatesCustomParameters struct {
	// PEM encoded certificate (chain).
	Certificates string `json:"certificates"`
	// Overwrite existing custom or ACME certificate files.
	Force *bool `json:"force"`
	// PEM encoded private key.
	Key *string `json:"key"`
	// Restart pveproxy.
	Restart *bool `json:"restart"`
}

type PostNodeCertificatesCustomResponse struct {
}

func (c *Client) PostNodeCertificatesCustom(Node string, parameters PostNodeCertificatesCustomParameters) (*PostNodeCertificatesCustomResponse, error) {

}

type GetNodeConfigParameters struct {
	// Return only a specific property from the node configuration.
	Property *NodeConfigProperty `json:"property"`
}

type GetNodeConfigResponse struct {
}

func (c *Client) GetNodeConfig(Node string, parameters GetNodeConfigParameters) (*GetNodeConfigResponse, error) {

}

type PutNodeConfigParameters struct {
	// Node specific ACME settings.
	Acme *string `json:"acme"`
	// ACME domain and validation plugin
	AcmedomainN *string `json:"acmedomain[n]"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Description for the Node. Shown in the web-interface node notes panel. This is saved as comment inside the configuration file.
	Description *string `json:"description"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Initial delay in seconds, before starting all the Virtual Guests with on-boot enabled.
	StartallOnbootDelay *int `json:"startall-onboot-delay"`
	// MAC address for wake on LAN
	Wakeonlan *string `json:"wakeonlan"`
}

type PutNodeConfigResponse struct {
}

func (c *Client) PutNodeConfig(Node string, parameters PutNodeConfigParameters) (*PutNodeConfigResponse, error) {

}

type GetNodeSdnResponse struct {
}

func (c *Client) GetNodeSdn(Node string) (*GetNodeSdnResponse, error) {

}

type GetNodeSdnZonesResponse struct {
}

func (c *Client) GetNodeSdnZones(Node string) (*GetNodeSdnZonesResponse, error) {

}

type GetNodeSdnZoneResponse struct {
}

func (c *Client) GetNodeSdnZone(Node string, Zone string) (*GetNodeSdnZoneResponse, error) {

}

type GetNodeSdnZoneContentResponse struct {
}

func (c *Client) GetNodeSdnZoneContent(Node string, Zone string) (*GetNodeSdnZoneContentResponse, error) {

}

type GetNodeVersionResponse struct {
}

func (c *Client) GetNodeVersion(Node string) (*GetNodeVersionResponse, error) {

}

type GetNodeStatusResponse struct {
}

func (c *Client) GetNodeStatus(Node string) (*GetNodeStatusResponse, error) {

}

type PostNodeStatusParameters struct {
	// Specify the command.
	Command NodeStatusCommand `json:"command"`
}

type PostNodeStatusResponse struct {
}

func (c *Client) PostNodeStatus(Node string, parameters PostNodeStatusParameters) (*PostNodeStatusResponse, error) {

}

type GetNodeNetstatResponse struct {
}

func (c *Client) GetNodeNetstat(Node string) (*GetNodeNetstatResponse, error) {

}

type PostNodeExecuteParameters struct {
	// JSON encoded array of commands.
	Commands string `json:"commands"`
}

type PostNodeExecuteResponse struct {
}

func (c *Client) PostNodeExecute(Node string, parameters PostNodeExecuteParameters) (*PostNodeExecuteResponse, error) {

}

type PostNodeWakeonlanResponse struct {
}

func (c *Client) PostNodeWakeonlan(Node string) (*PostNodeWakeonlanResponse, error) {

}

type GetNodeRrdParameters struct {
	// The RRD consolidation function
	Cf *NodeRrdCf `json:"cf"`
	// The list of datasources you want to display.
	Ds string `json:"ds"`
	// Specify the time frame you are interested in.
	Timeframe NodeRrdTimeframe `json:"timeframe"`
}

type GetNodeRrdResponse struct {
}

func (c *Client) GetNodeRrd(Node string, parameters GetNodeRrdParameters) (*GetNodeRrdResponse, error) {

}

type GetNodeRrddataParameters struct {
	// The RRD consolidation function
	Cf *NodeRrdCf `json:"cf"`
	// Specify the time frame you are interested in.
	Timeframe NodeRrdTimeframe `json:"timeframe"`
}

type GetNodeRrddataResponse struct {
}

func (c *Client) GetNodeRrddata(Node string, parameters GetNodeRrddataParameters) (*GetNodeRrddataResponse, error) {

}

type GetNodeSyslogParameters struct {
	Limit *int `json:"limit"`
	// Service ID
	Service *string `json:"service"`
	// Display all log since this date-time string.
	Since *string `json:"since"`
	Start *int    `json:"start"`
	// Display all log until this date-time string.
	Until *string `json:"until"`
}

type GetNodeSyslogResponse struct {
}

func (c *Client) GetNodeSyslog(Node string, parameters GetNodeSyslogParameters) (*GetNodeSyslogResponse, error) {

}

type GetNodeJournalParameters struct {
	// End before the given Cursor. Conflicts with 'until'
	Endcursor *string `json:"endcursor"`
	// Limit to the last X lines. Conflicts with a range.
	Lastentries *int `json:"lastentries"`
	// Display all log since this UNIX epoch. Conflicts with 'startcursor'.
	Since *int `json:"since"`
	// Start after the given Cursor. Conflicts with 'since'
	Startcursor *string `json:"startcursor"`
	// Display all log until this UNIX epoch. Conflicts with 'endcursor'.
	Until *int `json:"until"`
}

type GetNodeJournalResponse struct {
}

func (c *Client) GetNodeJournal(Node string, parameters GetNodeJournalParameters) (*GetNodeJournalResponse, error) {

}

type PostNodeVncshellParameters struct {
	// Run specific command or default to login.
	Cmd *NodeCmd `json:"cmd"`
	// Add parameters to a command. Encoded as null terminated strings.
	CmdOpts *string `json:"cmd-opts"`
	// sets the height of the console in pixels.
	Height *int `json:"height"`
	// use websocket instead of standard vnc.
	Websocket *bool `json:"websocket"`
	// sets the width of the console in pixels.
	Width *int `json:"width"`
}

type PostNodeVncshellResponse struct {
}

func (c *Client) PostNodeVncshell(Node string, parameters PostNodeVncshellParameters) (*PostNodeVncshellResponse, error) {

}

type PostNodeTermproxyParameters struct {
	// Run specific command or default to login.
	Cmd *NodeCmd `json:"cmd"`
	// Add parameters to a command. Encoded as null terminated strings.
	CmdOpts *string `json:"cmd-opts"`
}

type PostNodeTermproxyResponse struct {
}

func (c *Client) PostNodeTermproxy(Node string, parameters PostNodeTermproxyParameters) (*PostNodeTermproxyResponse, error) {

}

type GetNodeVncwebsocketParameters struct {
	// Port number returned by previous vncproxy call.
	Port int `json:"port"`
	// Ticket from previous call to vncproxy.
	Vncticket string `json:"vncticket"`
}

type GetNodeVncwebsocketResponse struct {
}

func (c *Client) GetNodeVncwebsocket(Node string, parameters GetNodeVncwebsocketParameters) (*GetNodeVncwebsocketResponse, error) {

}

type PostNodeSpiceshellParameters struct {
	// Run specific command or default to login.
	Cmd *NodeCmd `json:"cmd"`
	// Add parameters to a command. Encoded as null terminated strings.
	CmdOpts *string `json:"cmd-opts"`
	// SPICE proxy server. This can be used by the client to specify the proxy server. All nodes in a cluster runs 'spiceproxy', so it is up to the client to choose one. By default, we return the node where the VM is currently running. As reasonable setting is to use same node you use to connect to the API (This is window.location.hostname for the JS GUI).
	Proxy *string `json:"proxy"`
}

type PostNodeSpiceshellResponse struct {
}

func (c *Client) PostNodeSpiceshell(Node string, parameters PostNodeSpiceshellParameters) (*PostNodeSpiceshellResponse, error) {

}

type GetNodeDnsResponse struct {
}

func (c *Client) GetNodeDns(Node string) (*GetNodeDnsResponse, error) {

}

type PutNodeDnsParameters struct {
	// First name server IP address.
	Dns1 *string `json:"dns1"`
	// Second name server IP address.
	Dns2 *string `json:"dns2"`
	// Third name server IP address.
	Dns3 *string `json:"dns3"`
	// Search domain for host-name lookup.
	Search string `json:"search"`
}

type PutNodeDnsResponse struct {
}

func (c *Client) PutNodeDns(Node string, parameters PutNodeDnsParameters) (*PutNodeDnsResponse, error) {

}

type GetNodeTimeResponse struct {
}

func (c *Client) GetNodeTime(Node string) (*GetNodeTimeResponse, error) {

}

type PutNodeTimeParameters struct {
	// Time zone. The file '/usr/share/zoneinfo/zone.tab' contains the list of valid names.
	Timezone string `json:"timezone"`
}

type PutNodeTimeResponse struct {
}

func (c *Client) PutNodeTime(Node string, parameters PutNodeTimeParameters) (*PutNodeTimeResponse, error) {

}

type GetNodeAplinfoResponse struct {
}

func (c *Client) GetNodeAplinfo(Node string) (*GetNodeAplinfoResponse, error) {

}

type PostNodeAplinfoParameters struct {
	// The storage where the template will be stored
	Storage string `json:"storage"`
	// The template which will downloaded
	Template string `json:"template"`
}

type PostNodeAplinfoResponse struct {
}

func (c *Client) PostNodeAplinfo(Node string, parameters PostNodeAplinfoParameters) (*PostNodeAplinfoResponse, error) {

}

type GetNodeQueryUrlMetadataParameters struct {
	// The URL to query the metadata from.
	Url string `json:"url"`
	// If false, no SSL/TLS certificates will be verified.
	VerifyCertificates *bool `json:"verify-certificates"`
}

type GetNodeQueryUrlMetadataResponse struct {
}

func (c *Client) GetNodeQueryUrlMetadata(Node string, parameters GetNodeQueryUrlMetadataParameters) (*GetNodeQueryUrlMetadataResponse, error) {

}

type GetNodeReportResponse struct {
}

func (c *Client) GetNodeReport(Node string) (*GetNodeReportResponse, error) {

}

type PostNodeStartallParameters struct {
	// Issue start command even if virtual guest have 'onboot' not set or set to off.
	Force *bool `json:"force"`
	// Only consider guests from this comma separated list of VMIDs.
	Vms *string `json:"vms"`
}

type PostNodeStartallResponse struct {
}

func (c *Client) PostNodeStartall(Node string, parameters PostNodeStartallParameters) (*PostNodeStartallResponse, error) {

}

type PostNodeStopallParameters struct {
	// Force a hard-stop after the timeout.
	ForceStop *bool `json:"force-stop"`
	// Timeout for each guest shutdown task. Depending on `force-stop`, the shutdown gets then simply aborted or a hard-stop is forced.
	Timeout *int `json:"timeout"`
	// Only consider Guests with these IDs.
	Vms *string `json:"vms"`
}

type PostNodeStopallResponse struct {
}

func (c *Client) PostNodeStopall(Node string, parameters PostNodeStopallParameters) (*PostNodeStopallResponse, error) {

}

type PostNodeMigrateallParameters struct {
	// Maximal number of parallel migration job. If not set, uses'max_workers' from datacenter.cfg. One of both must be set!
	Maxworkers *int `json:"maxworkers"`
	// Target node.
	Target string `json:"target"`
	// Only consider Guests with these IDs.
	Vms *string `json:"vms"`
	// Enable live storage migration for local disk
	WithLocalDisks *bool `json:"with-local-disks"`
}

type PostNodeMigrateallResponse struct {
}

func (c *Client) PostNodeMigrateall(Node string, parameters PostNodeMigrateallParameters) (*PostNodeMigrateallResponse, error) {

}

type GetNodeHostsResponse struct {
}

func (c *Client) GetNodeHosts(Node string) (*GetNodeHostsResponse, error) {

}

type PostNodeHostsParameters struct {
	// The target content of /etc/hosts.
	Data string `json:"data"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
}

type PostNodeHostsResponse struct {
}

func (c *Client) PostNodeHosts(Node string, parameters PostNodeHostsParameters) (*PostNodeHostsResponse, error) {

}

type GetStorageParameters struct {
	// Only list storage of specific type
	Type *StorageType `json:"type"`
}

type GetStorageResponse struct {
}

func (c *Client) GetStorage(parameters GetStorageParameters) (*GetStorageResponse, error) {

}

type PostStorageParameters struct {
	// Authsupported.
	Authsupported *string `json:"authsupported"`
	// Base volume. This volume is automatically activated.
	Base *string `json:"base"`
	// block size
	Blocksize *string `json:"blocksize"`
	// Set I/O bandwidth limit for various operations (in KiB/s).
	Bwlimit *string `json:"bwlimit"`
	// host group for comstar views
	ComstarHg *string `json:"comstar_hg"`
	// target group for comstar views
	ComstarTg *string `json:"comstar_tg"`
	// Allowed content types.
	//
	// NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs.
	//
	Content *string `json:"content"`
	// Overrides for default content type directories.
	ContentDirs *string `json:"content-dirs"`
	// Create the base directory if it doesn't exist.
	CreateBasePath *bool `json:"create-base-path"`
	// Populate the directory with the default structure.
	CreateSubdirs *bool `json:"create-subdirs"`
	// Data Pool (for erasure coding only)
	DataPool *string `json:"data-pool"`
	// Proxmox Backup Server datastore name.
	Datastore *string `json:"datastore"`
	// Flag to disable the storage.
	Disable *bool `json:"disable"`
	// CIFS domain.
	Domain *string `json:"domain"`
	// Encryption key. Use 'autogen' to generate one automatically without passphrase.
	EncryptionKey *string `json:"encryption-key"`
	// NFS export path.
	Export *string `json:"export"`
	// Certificate SHA 256 fingerprint.
	Fingerprint *string `json:"fingerprint"`
	// Default image format.
	Format *string `json:"format"`
	// The Ceph filesystem name.
	FsName *string `json:"fs-name"`
	// Mount CephFS through FUSE.
	Fuse *bool `json:"fuse"`
	// Assume the given path is an externally managed mountpoint and consider the storage offline if it is not mounted. Using a boolean (yes/no) value serves as a shortcut to using the target path in this field.
	IsMountpoint *string `json:"is_mountpoint"`
	// iscsi provider
	Iscsiprovider *string `json:"iscsiprovider"`
	// Client keyring contents (for external clusters).
	Keyring *string `json:"keyring"`
	// Always access rbd through krbd kernel module.
	Krbd *bool `json:"krbd"`
	// target portal group for Linux LIO targets
	LioTpg *string `json:"lio_tpg"`
	// Base64-encoded, PEM-formatted public RSA key. Used to encrypt a copy of the encryption-key which will be added to each encrypted backup.
	MasterPubkey *string `json:"master-pubkey"`
	// Maximal number of protected backups per guest. Use '-1' for unlimited.
	MaxProtectedBackups *int `json:"max-protected-backups"`
	// Deprecated: use 'prune-backups' instead. Maximal number of backup files per VM. Use '0' for unlimited.
	Maxfiles *int `json:"maxfiles"`
	// Create the directory if it doesn't exist and populate it with default sub-dirs. NOTE: Deprecated, use the 'create-base-path' and 'create-subdirs' options instead.
	Mkdir *bool `json:"mkdir"`
	// IP addresses of monitors (for external clusters).
	Monhost *string `json:"monhost"`
	// mount point
	Mountpoint *string `json:"mountpoint"`
	// Namespace.
	Namespace *string `json:"namespace"`
	// Set the NOCOW flag on files. Disables data checksumming and causes data errors to be unrecoverable from while allowing direct I/O. Only use this if data does not need to be any more safe than on a single ext4 formatted disk with no underlying raid system.
	Nocow *bool `json:"nocow"`
	// List of cluster node names.
	Nodes *string `json:"nodes"`
	// disable write caching on the target
	Nowritecache *bool `json:"nowritecache"`
	// NFS/CIFS mount options (see 'man nfs' or 'man mount.cifs')
	Options *string `json:"options"`
	// Password for accessing the share/datastore.
	Password *string `json:"password"`
	// File system path.
	Path *string `json:"path"`
	// Pool.
	Pool *string `json:"pool"`
	// For non default port.
	Port *int `json:"port"`
	// iSCSI portal (IP or DNS name with optional port).
	Portal *string `json:"portal"`
	// Preallocation mode for raw and qcow2 images. Using 'metadata' on raw images results in preallocation=off.
	Preallocation *StoragePreallocation `json:"preallocation"`
	// The retention options with shorter intervals are processed first with --keep-last being the very first one. Each option covers a specific period of time. We say that backups within this period are covered by this option. The next option does not take care of already covered backups and only considers older backups.
	PruneBackups *string `json:"prune-backups"`
	// Zero-out data when removing LVs.
	Saferemove *bool `json:"saferemove"`
	// Wipe throughput (cstream -t parameter value).
	SaferemoveThroughput *string `json:"saferemove_throughput"`
	// Server IP or DNS name.
	Server *string `json:"server"`
	// Backup volfile server IP or DNS name.
	Server2 *string `json:"server2"`
	// CIFS share.
	Share *string `json:"share"`
	// Mark storage as shared.
	Shared *bool `json:"shared"`
	// SMB protocol version. 'default' if not set, negotiates the highest SMB2+ version supported by both the client and server.
	Smbversion *StorageSmbversion `json:"smbversion"`
	// use sparse volumes
	Sparse *bool `json:"sparse"`
	// The storage identifier.
	Storage string `json:"storage"`
	// Subdir to mount.
	Subdir *string `json:"subdir"`
	// Only use logical volumes tagged with 'pve-vm-ID'.
	TaggedOnly *bool `json:"tagged_only"`
	// iSCSI target.
	Target *string `json:"target"`
	// LVM thin pool LV name.
	Thinpool *string `json:"thinpool"`
	// Gluster transport: tcp or rdma
	Transport *StorageTransport `json:"transport"`
	// Storage type.
	Type StorageType `json:"type"`
	// RBD Id.
	Username *string `json:"username"`
	// Volume group name.
	Vgname *string `json:"vgname"`
	// Glusterfs Volume.
	Volume *string `json:"volume"`
}

type PostStorageResponse struct {
}

func (c *Client) PostStorage(parameters PostStorageParameters) (*PostStorageResponse, error) {

}

type DeleteStorageByStorageResponse struct {
}

func (c *Client) DeleteStorageByStorage(Storage string) (*DeleteStorageByStorageResponse, error) {

}

type GetStorageByStorageResponse struct {
}

func (c *Client) GetStorageByStorage(Storage string) (*GetStorageByStorageResponse, error) {

}

type PutStorageByStorageParameters struct {
	// block size
	Blocksize *string `json:"blocksize"`
	// Set I/O bandwidth limit for various operations (in KiB/s).
	Bwlimit *string `json:"bwlimit"`
	// host group for comstar views
	ComstarHg *string `json:"comstar_hg"`
	// target group for comstar views
	ComstarTg *string `json:"comstar_tg"`
	// Allowed content types.
	//
	// NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs.
	//
	Content *string `json:"content"`
	// Overrides for default content type directories.
	ContentDirs *string `json:"content-dirs"`
	// Create the base directory if it doesn't exist.
	CreateBasePath *bool `json:"create-base-path"`
	// Populate the directory with the default structure.
	CreateSubdirs *bool `json:"create-subdirs"`
	// Data Pool (for erasure coding only)
	DataPool *string `json:"data-pool"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// Flag to disable the storage.
	Disable *bool `json:"disable"`
	// CIFS domain.
	Domain *string `json:"domain"`
	// Encryption key. Use 'autogen' to generate one automatically without passphrase.
	EncryptionKey *string `json:"encryption-key"`
	// Certificate SHA 256 fingerprint.
	Fingerprint *string `json:"fingerprint"`
	// Default image format.
	Format *string `json:"format"`
	// The Ceph filesystem name.
	FsName *string `json:"fs-name"`
	// Mount CephFS through FUSE.
	Fuse *bool `json:"fuse"`
	// Assume the given path is an externally managed mountpoint and consider the storage offline if it is not mounted. Using a boolean (yes/no) value serves as a shortcut to using the target path in this field.
	IsMountpoint *string `json:"is_mountpoint"`
	// Client keyring contents (for external clusters).
	Keyring *string `json:"keyring"`
	// Always access rbd through krbd kernel module.
	Krbd *bool `json:"krbd"`
	// target portal group for Linux LIO targets
	LioTpg *string `json:"lio_tpg"`
	// Base64-encoded, PEM-formatted public RSA key. Used to encrypt a copy of the encryption-key which will be added to each encrypted backup.
	MasterPubkey *string `json:"master-pubkey"`
	// Maximal number of protected backups per guest. Use '-1' for unlimited.
	MaxProtectedBackups *int `json:"max-protected-backups"`
	// Deprecated: use 'prune-backups' instead. Maximal number of backup files per VM. Use '0' for unlimited.
	Maxfiles *int `json:"maxfiles"`
	// Create the directory if it doesn't exist and populate it with default sub-dirs. NOTE: Deprecated, use the 'create-base-path' and 'create-subdirs' options instead.
	Mkdir *bool `json:"mkdir"`
	// IP addresses of monitors (for external clusters).
	Monhost *string `json:"monhost"`
	// mount point
	Mountpoint *string `json:"mountpoint"`
	// Namespace.
	Namespace *string `json:"namespace"`
	// Set the NOCOW flag on files. Disables data checksumming and causes data errors to be unrecoverable from while allowing direct I/O. Only use this if data does not need to be any more safe than on a single ext4 formatted disk with no underlying raid system.
	Nocow *bool `json:"nocow"`
	// List of cluster node names.
	Nodes *string `json:"nodes"`
	// disable write caching on the target
	Nowritecache *bool `json:"nowritecache"`
	// NFS/CIFS mount options (see 'man nfs' or 'man mount.cifs')
	Options *string `json:"options"`
	// Password for accessing the share/datastore.
	Password *string `json:"password"`
	// Pool.
	Pool *string `json:"pool"`
	// For non default port.
	Port *int `json:"port"`
	// Preallocation mode for raw and qcow2 images. Using 'metadata' on raw images results in preallocation=off.
	Preallocation *StoragePreallocation `json:"preallocation"`
	// The retention options with shorter intervals are processed first with --keep-last being the very first one. Each option covers a specific period of time. We say that backups within this period are covered by this option. The next option does not take care of already covered backups and only considers older backups.
	PruneBackups *string `json:"prune-backups"`
	// Zero-out data when removing LVs.
	Saferemove *bool `json:"saferemove"`
	// Wipe throughput (cstream -t parameter value).
	SaferemoveThroughput *string `json:"saferemove_throughput"`
	// Server IP or DNS name.
	Server *string `json:"server"`
	// Backup volfile server IP or DNS name.
	Server2 *string `json:"server2"`
	// Mark storage as shared.
	Shared *bool `json:"shared"`
	// SMB protocol version. 'default' if not set, negotiates the highest SMB2+ version supported by both the client and server.
	Smbversion *StorageSmbversion `json:"smbversion"`
	// use sparse volumes
	Sparse *bool `json:"sparse"`
	// Subdir to mount.
	Subdir *string `json:"subdir"`
	// Only use logical volumes tagged with 'pve-vm-ID'.
	TaggedOnly *bool `json:"tagged_only"`
	// Gluster transport: tcp or rdma
	Transport *StorageTransport `json:"transport"`
	// RBD Id.
	Username *string `json:"username"`
}

type PutStorageByStorageResponse struct {
}

func (c *Client) PutStorageByStorage(Storage string, parameters PutStorageByStorageParameters) (*PutStorageByStorageResponse, error) {

}

type GetAccessResponse struct {
}

func (c *Client) GetAccess() (*GetAccessResponse, error) {

}

type GetAccessUsersParameters struct {
	// Optional filter for enable property.
	Enabled *bool `json:"enabled"`
	// Include group and token information.
	Full *bool `json:"full"`
}

type GetAccessUsersResponse struct {
}

func (c *Client) GetAccessUsers(parameters GetAccessUsersParameters) (*GetAccessUsersResponse, error) {

}

type PostAccessUsersParameters struct {
	Comment *string `json:"comment"`
	Email   *string `json:"email"`
	// Enable the account (default). You can set this to '0' to disable the account
	Enable *bool `json:"enable"`
	// Account expiration date (seconds since epoch). '0' means no expiration date.
	Expire    *int    `json:"expire"`
	Firstname *string `json:"firstname"`
	Groups    *string `json:"groups"`
	// Keys for two factor auth (yubico).
	Keys     *string `json:"keys"`
	Lastname *string `json:"lastname"`
	// Initial password.
	Password *string `json:"password"`
	// Full User ID, in the `name@realm` format.
	Userid string `json:"userid"`
}

type PostAccessUsersResponse struct {
}

func (c *Client) PostAccessUsers(parameters PostAccessUsersParameters) (*PostAccessUsersResponse, error) {

}

type DeleteAccessUserResponse struct {
}

func (c *Client) DeleteAccessUser(Userid string) (*DeleteAccessUserResponse, error) {

}

type GetAccessUserResponse struct {
}

func (c *Client) GetAccessUser(Userid string) (*GetAccessUserResponse, error) {

}

type PutAccessUserParameters struct {
	Append  *bool   `json:"append"`
	Comment *string `json:"comment"`
	Email   *string `json:"email"`
	// Enable the account (default). You can set this to '0' to disable the account
	Enable *bool `json:"enable"`
	// Account expiration date (seconds since epoch). '0' means no expiration date.
	Expire    *int    `json:"expire"`
	Firstname *string `json:"firstname"`
	Groups    *string `json:"groups"`
	// Keys for two factor auth (yubico).
	Keys     *string `json:"keys"`
	Lastname *string `json:"lastname"`
}

type PutAccessUserResponse struct {
}

func (c *Client) PutAccessUser(Userid string, parameters PutAccessUserParameters) (*PutAccessUserResponse, error) {

}

type GetAccessUserTfaParameters struct {
	// Request all entries as an array.
	Multiple *bool `json:"multiple"`
}

type GetAccessUserTfaResponse struct {
}

func (c *Client) GetAccessUserTfa(Userid string, parameters GetAccessUserTfaParameters) (*GetAccessUserTfaResponse, error) {

}

type PutAccessUserUnlockTfaResponse struct {
}

func (c *Client) PutAccessUserUnlockTfa(Userid string) (*PutAccessUserUnlockTfaResponse, error) {

}

type GetAccessUserTokenResponse struct {
}

func (c *Client) GetAccessUserToken(Userid string) (*GetAccessUserTokenResponse, error) {

}

type DeleteAccessUserTokenByTokenidResponse struct {
}

func (c *Client) DeleteAccessUserTokenByTokenid(Userid string, Tokenid string) (*DeleteAccessUserTokenByTokenidResponse, error) {

}

type GetAccessUserTokenByTokenidResponse struct {
}

func (c *Client) GetAccessUserTokenByTokenid(Userid string, Tokenid string) (*GetAccessUserTokenByTokenidResponse, error) {

}

type PostAccessUserTokenByTokenidParameters struct {
	Comment *string `json:"comment"`
	// API token expiration date (seconds since epoch). '0' means no expiration date.
	Expire *int `json:"expire"`
	// Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
	Privsep *bool `json:"privsep"`
}

type PostAccessUserTokenByTokenidResponse struct {
}

func (c *Client) PostAccessUserTokenByTokenid(Userid string, Tokenid string, parameters PostAccessUserTokenByTokenidParameters) (*PostAccessUserTokenByTokenidResponse, error) {

}

type PutAccessUserTokenByTokenidParameters struct {
	Comment *string `json:"comment"`
	// API token expiration date (seconds since epoch). '0' means no expiration date.
	Expire *int `json:"expire"`
	// Restrict API token privileges with separate ACLs (default), or give full privileges of corresponding user.
	Privsep *bool `json:"privsep"`
}

type PutAccessUserTokenByTokenidResponse struct {
}

func (c *Client) PutAccessUserTokenByTokenid(Userid string, Tokenid string, parameters PutAccessUserTokenByTokenidParameters) (*PutAccessUserTokenByTokenidResponse, error) {

}

type GetAccessGroupsResponse struct {
}

func (c *Client) GetAccessGroups() (*GetAccessGroupsResponse, error) {

}

type PostAccessGroupsParameters struct {
	Comment *string `json:"comment"`
	Groupid string  `json:"groupid"`
}

type PostAccessGroupsResponse struct {
}

func (c *Client) PostAccessGroups(parameters PostAccessGroupsParameters) (*PostAccessGroupsResponse, error) {

}

type DeleteAccessGroupResponse struct {
}

func (c *Client) DeleteAccessGroup(Groupid string) (*DeleteAccessGroupResponse, error) {

}

type GetAccessGroupResponse struct {
}

func (c *Client) GetAccessGroup(Groupid string) (*GetAccessGroupResponse, error) {

}

type PutAccessGroupParameters struct {
	Comment *string `json:"comment"`
}

type PutAccessGroupResponse struct {
}

func (c *Client) PutAccessGroup(Groupid string, parameters PutAccessGroupParameters) (*PutAccessGroupResponse, error) {

}

type GetAccessRolesResponse struct {
}

func (c *Client) GetAccessRoles() (*GetAccessRolesResponse, error) {

}

type PostAccessRolesParameters struct {
	Privs  *string `json:"privs"`
	Roleid string  `json:"roleid"`
}

type PostAccessRolesResponse struct {
}

func (c *Client) PostAccessRoles(parameters PostAccessRolesParameters) (*PostAccessRolesResponse, error) {

}

type DeleteAccessRoleResponse struct {
}

func (c *Client) DeleteAccessRole(Roleid string) (*DeleteAccessRoleResponse, error) {

}

type GetAccessRoleResponse struct {
}

func (c *Client) GetAccessRole(Roleid string) (*GetAccessRoleResponse, error) {

}

type PutAccessRoleParameters struct {
	Append *bool   `json:"append"`
	Privs  *string `json:"privs"`
}

type PutAccessRoleResponse struct {
}

func (c *Client) PutAccessRole(Roleid string, parameters PutAccessRoleParameters) (*PutAccessRoleResponse, error) {

}

type GetAccessAclResponse struct {
}

func (c *Client) GetAccessAcl() (*GetAccessAclResponse, error) {

}

type PutAccessAclParameters struct {
	// Remove permissions (instead of adding it).
	Delete *bool `json:"delete"`
	// List of groups.
	Groups *string `json:"groups"`
	// Access control path
	Path string `json:"path"`
	// Allow to propagate (inherit) permissions.
	Propagate *bool `json:"propagate"`
	// List of roles.
	Roles string `json:"roles"`
	// List of API tokens.
	Tokens *string `json:"tokens"`
	// List of users.
	Users *string `json:"users"`
}

type PutAccessAclResponse struct {
}

func (c *Client) PutAccessAcl(parameters PutAccessAclParameters) (*PutAccessAclResponse, error) {

}

type GetAccessDomainsResponse struct {
}

func (c *Client) GetAccessDomains() (*GetAccessDomainsResponse, error) {

}

type PostAccessDomainsParameters struct {
	// Specifies the Authentication Context Class Reference values that theAuthorization Server is being requested to use for the Auth Request.
	AcrValues *string `json:"acr-values"`
	// Automatically create users if they do not exist.
	Autocreate *bool `json:"autocreate"`
	// LDAP base domain name
	BaseDn *string `json:"base_dn"`
	// LDAP bind domain name
	BindDn *string `json:"bind_dn"`
	// Path to the CA certificate store
	Capath *string `json:"capath"`
	// username is case-sensitive
	CaseSensitive *bool `json:"case-sensitive"`
	// Path to the client certificate
	Cert *string `json:"cert"`
	// Path to the client certificate key
	Certkey *string `json:"certkey"`
	// OpenID Client ID
	ClientId *string `json:"client-id"`
	// OpenID Client Key
	ClientKey *string `json:"client-key"`
	// Description.
	Comment *string `json:"comment"`
	// Use this as default realm
	Default *bool `json:"default"`
	// AD domain name
	Domain *string `json:"domain"`
	// LDAP filter for user sync.
	Filter *string `json:"filter"`
	// The objectclasses for groups.
	GroupClasses *string `json:"group_classes"`
	// LDAP base domain name for group sync. If not set, the base_dn will be used.
	GroupDn *string `json:"group_dn"`
	// LDAP filter for group sync.
	GroupFilter *string `json:"group_filter"`
	// LDAP attribute representing a groups name. If not set or found, the first value of the DN will be used as name.
	GroupNameAttr *string `json:"group_name_attr"`
	// OpenID Issuer Url
	IssuerUrl *string `json:"issuer-url"`
	// LDAP protocol mode.
	Mode *AccessDomainMode `json:"mode"`
	// LDAP bind password. Will be stored in '/etc/pve/priv/realm/<REALM>.pw'.
	Password *string `json:"password"`
	// Server port.
	Port *int `json:"port"`
	// Specifies whether the Authorization Server prompts the End-User for reauthentication and consent.
	Prompt *string `json:"prompt"`
	// Authentication domain ID
	Realm string `json:"realm"`
	// Specifies the scopes (user details) that should be authorized and returned, for example 'email' or 'profile'.
	Scopes *string `json:"scopes"`
	// Use secure LDAPS protocol. DEPRECATED: use 'mode' instead.
	Secure *bool `json:"secure"`
	// Server IP address (or DNS name)
	Server1 *string `json:"server1"`
	// Fallback Server IP address (or DNS name)
	Server2 *string `json:"server2"`
	// LDAPS TLS/SSL version. It's not recommended to use version older than 1.2!
	Sslversion *AccessDomainSslversion `json:"sslversion"`
	// The default options for behavior of synchronizations.
	SyncDefaultsOptions *string `json:"sync-defaults-options"`
	// Comma separated list of key=value pairs for specifying which LDAP attributes map to which PVE user field. For example, to map the LDAP attribute 'mail' to PVEs 'email', write  'email=mail'. By default, each PVE user field is represented  by an LDAP attribute of the same name.
	SyncAttributes *string `json:"sync_attributes"`
	// Use Two-factor authentication.
	Tfa *string `json:"tfa"`
	// Realm type.
	Type AccessDomainsType `json:"type"`
	// LDAP user attribute name
	UserAttr *string `json:"user_attr"`
	// The objectclasses for users.
	UserClasses *string `json:"user_classes"`
	// OpenID claim used to generate the unique username.
	UsernameClaim *string `json:"username-claim"`
	// Verify the server's SSL certificate
	Verify *bool `json:"verify"`
}

type PostAccessDomainsResponse struct {
}

func (c *Client) PostAccessDomains(parameters PostAccessDomainsParameters) (*PostAccessDomainsResponse, error) {

}

type DeleteAccessDomainResponse struct {
}

func (c *Client) DeleteAccessDomain(Realm string) (*DeleteAccessDomainResponse, error) {

}

type GetAccessDomainResponse struct {
}

func (c *Client) GetAccessDomain(Realm string) (*GetAccessDomainResponse, error) {

}

type PutAccessDomainParameters struct {
	// Specifies the Authentication Context Class Reference values that theAuthorization Server is being requested to use for the Auth Request.
	AcrValues *string `json:"acr-values"`
	// Automatically create users if they do not exist.
	Autocreate *bool `json:"autocreate"`
	// LDAP base domain name
	BaseDn *string `json:"base_dn"`
	// LDAP bind domain name
	BindDn *string `json:"bind_dn"`
	// Path to the CA certificate store
	Capath *string `json:"capath"`
	// username is case-sensitive
	CaseSensitive *bool `json:"case-sensitive"`
	// Path to the client certificate
	Cert *string `json:"cert"`
	// Path to the client certificate key
	Certkey *string `json:"certkey"`
	// OpenID Client ID
	ClientId *string `json:"client-id"`
	// OpenID Client Key
	ClientKey *string `json:"client-key"`
	// Description.
	Comment *string `json:"comment"`
	// Use this as default realm
	Default *bool `json:"default"`
	// A list of settings you want to delete.
	Delete *string `json:"delete"`
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest"`
	// AD domain name
	Domain *string `json:"domain"`
	// LDAP filter for user sync.
	Filter *string `json:"filter"`
	// The objectclasses for groups.
	GroupClasses *string `json:"group_classes"`
	// LDAP base domain name for group sync. If not set, the base_dn will be used.
	GroupDn *string `json:"group_dn"`
	// LDAP filter for group sync.
	GroupFilter *string `json:"group_filter"`
	// LDAP attribute representing a groups name. If not set or found, the first value of the DN will be used as name.
	GroupNameAttr *string `json:"group_name_attr"`
	// OpenID Issuer Url
	IssuerUrl *string `json:"issuer-url"`
	// LDAP protocol mode.
	Mode *AccessDomainMode `json:"mode"`
	// LDAP bind password. Will be stored in '/etc/pve/priv/realm/<REALM>.pw'.
	Password *string `json:"password"`
	// Server port.
	Port *int `json:"port"`
	// Specifies whether the Authorization Server prompts the End-User for reauthentication and consent.
	Prompt *string `json:"prompt"`
	// Specifies the scopes (user details) that should be authorized and returned, for example 'email' or 'profile'.
	Scopes *string `json:"scopes"`
	// Use secure LDAPS protocol. DEPRECATED: use 'mode' instead.
	Secure *bool `json:"secure"`
	// Server IP address (or DNS name)
	Server1 *string `json:"server1"`
	// Fallback Server IP address (or DNS name)
	Server2 *string `json:"server2"`
	// LDAPS TLS/SSL version. It's not recommended to use version older than 1.2!
	Sslversion *AccessDomainSslversion `json:"sslversion"`
	// The default options for behavior of synchronizations.
	SyncDefaultsOptions *string `json:"sync-defaults-options"`
	// Comma separated list of key=value pairs for specifying which LDAP attributes map to which PVE user field. For example, to map the LDAP attribute 'mail' to PVEs 'email', write  'email=mail'. By default, each PVE user field is represented  by an LDAP attribute of the same name.
	SyncAttributes *string `json:"sync_attributes"`
	// Use Two-factor authentication.
	Tfa *string `json:"tfa"`
	// LDAP user attribute name
	UserAttr *string `json:"user_attr"`
	// The objectclasses for users.
	UserClasses *string `json:"user_classes"`
	// Verify the server's SSL certificate
	Verify *bool `json:"verify"`
}

type PutAccessDomainResponse struct {
}

func (c *Client) PutAccessDomain(Realm string, parameters PutAccessDomainParameters) (*PutAccessDomainResponse, error) {

}

type PostAccessDomainSyncParameters struct {
	// If set, does not write anything.
	DryRun *bool `json:"dry-run"`
	// Enable newly synced users immediately.
	EnableNew *bool `json:"enable-new"`
	// DEPRECATED: use 'remove-vanished' instead. If set, uses the LDAP Directory as source of truth, deleting users or groups not returned from the sync and removing all locally modified properties of synced users. If not set, only syncs information which is present in the synced data, and does not delete or modify anything else.
	Full *bool `json:"full"`
	// DEPRECATED: use 'remove-vanished' instead. Remove ACLs for users or groups which were removed from the config during a sync.
	Purge *bool `json:"purge"`
	// A semicolon-seperated list of things to remove when they or the user vanishes during a sync. The following values are possible: 'entry' removes the user/group when not returned from the sync. 'properties' removes the set properties on existing user/group that do not appear in the source (even custom ones). 'acl' removes acls when the user/group is not returned from the sync. Instead of a list it also can be 'none' (the default).
	RemoveVanished *string `json:"remove-vanished"`
	// Select what to sync.
	Scope *SyncScope `json:"scope"`
}

type PostAccessDomainSyncResponse struct {
}

func (c *Client) PostAccessDomainSync(Realm string, parameters PostAccessDomainSyncParameters) (*PostAccessDomainSyncResponse, error) {

}

type GetAccessOpenidResponse struct {
}

func (c *Client) GetAccessOpenid() (*GetAccessOpenidResponse, error) {

}

type PostAccessOpenidAuthUrlParameters struct {
	// Authentication domain ID
	Realm string `json:"realm"`
	// Redirection Url. The client should set this to the used server url (location.origin).
	RedirectUrl string `json:"redirect-url"`
}

type PostAccessOpenidAuthUrlResponse struct {
}

func (c *Client) PostAccessOpenidAuthUrl(parameters PostAccessOpenidAuthUrlParameters) (*PostAccessOpenidAuthUrlResponse, error) {

}

type PostAccessOpenidLoginParameters struct {
	// OpenId authorization code.
	Code string `json:"code"`
	// Redirection Url. The client should set this to the used server url (location.origin).
	RedirectUrl string `json:"redirect-url"`
	// OpenId state.
	State string `json:"state"`
}

type PostAccessOpenidLoginResponse struct {
}

func (c *Client) PostAccessOpenidLogin(parameters PostAccessOpenidLoginParameters) (*PostAccessOpenidLoginResponse, error) {

}

type GetAccessTfaResponse struct {
}

func (c *Client) GetAccessTfa() (*GetAccessTfaResponse, error) {

}

type GetAccessTfaByUseridResponse struct {
}

func (c *Client) GetAccessTfaByUserid(Userid string) (*GetAccessTfaByUseridResponse, error) {

}

type PostAccessTfaByUseridParameters struct {
	// When responding to a u2f challenge: the original challenge string
	Challenge *string `json:"challenge"`
	// A description to distinguish multiple entries from one another
	Description *string `json:"description"`
	// The current password.
	Password *string `json:"password"`
	// A totp URI.
	Totp *string `json:"totp"`
	// TFA Entry Type.
	Type AccessTfaType `json:"type"`
	// The current value for the provided totp URI, or a Webauthn/U2F challenge response
	Value *string `json:"value"`
}

type PostAccessTfaByUseridResponse struct {
}

func (c *Client) PostAccessTfaByUserid(Userid string, parameters PostAccessTfaByUseridParameters) (*PostAccessTfaByUseridResponse, error) {

}

type DeleteAccessTfaByUseridByIdParameters struct {
	// The current password.
	Password *string `json:"password"`
}

type DeleteAccessTfaByUseridByIdResponse struct {
}

func (c *Client) DeleteAccessTfaByUseridById(Userid string, Id string, parameters DeleteAccessTfaByUseridByIdParameters) (*DeleteAccessTfaByUseridByIdResponse, error) {

}

type GetAccessTfaByUseridByIdResponse struct {
}

func (c *Client) GetAccessTfaByUseridById(Userid string, Id string) (*GetAccessTfaByUseridByIdResponse, error) {

}

type PutAccessTfaByUseridByIdParameters struct {
	// A description to distinguish multiple entries from one another
	Description *string `json:"description"`
	// Whether the entry should be enabled for login.
	Enable *bool `json:"enable"`
	// The current password.
	Password *string `json:"password"`
}

type PutAccessTfaByUseridByIdResponse struct {
}

func (c *Client) PutAccessTfaByUseridById(Userid string, Id string, parameters PutAccessTfaByUseridByIdParameters) (*PutAccessTfaByUseridByIdResponse, error) {

}

type GetAccessTicketResponse struct {
}

func (c *Client) GetAccessTicket() (*GetAccessTicketResponse, error) {

}

type PostAccessTicketParameters struct {
	// This parameter is now ignored and assumed to be 1.
	NewFormat *bool `json:"new-format"`
	// One-time password for Two-factor authentication.
	Otp *string `json:"otp"`
	// The secret password. This can also be a valid ticket.
	Password string `json:"password"`
	// Verify ticket, and check if user have access 'privs' on 'path'
	Path *string `json:"path"`
	// Verify ticket, and check if user have access 'privs' on 'path'
	Privs *string `json:"privs"`
	// You can optionally pass the realm using this parameter. Normally the realm is simply added to the username <username>@<relam>.
	Realm *string `json:"realm"`
	// The signed TFA challenge string the user wants to respond to.
	TfaChallenge *string `json:"tfa-challenge"`
	// User name
	Username string `json:"username"`
}

type PostAccessTicketResponse struct {
}

func (c *Client) PostAccessTicket(parameters PostAccessTicketParameters) (*PostAccessTicketResponse, error) {

}

type PutAccessPasswordParameters struct {
	// The new password.
	Password string `json:"password"`
	// Full User ID, in the `name@realm` format.
	Userid string `json:"userid"`
}

type PutAccessPasswordResponse struct {
}

func (c *Client) PutAccessPassword(parameters PutAccessPasswordParameters) (*PutAccessPasswordResponse, error) {

}

type GetAccessPermissionsParameters struct {
	// Only dump this specific path, not the whole tree.
	Path *string `json:"path"`
	// User ID or full API token ID
	Userid *string `json:"userid"`
}

type GetAccessPermissionsResponse struct {
}

func (c *Client) GetAccessPermissions(parameters GetAccessPermissionsParameters) (*GetAccessPermissionsResponse, error) {

}

type GetPoolsResponse struct {
}

func (c *Client) GetPools() (*GetPoolsResponse, error) {

}

type PostPoolsParameters struct {
	Comment *string `json:"comment"`
	Poolid  string  `json:"poolid"`
}

type PostPoolsResponse struct {
}

func (c *Client) PostPools(parameters PostPoolsParameters) (*PostPoolsResponse, error) {

}

type DeletePoolResponse struct {
}

func (c *Client) DeletePool(Poolid string) (*DeletePoolResponse, error) {

}

type GetPoolParameters struct {
	Type *PoolType `json:"type"`
}

type GetPoolResponse struct {
}

func (c *Client) GetPool(Poolid string, parameters GetPoolParameters) (*GetPoolResponse, error) {

}

type PutPoolParameters struct {
	Comment *string `json:"comment"`
	// Remove vms/storage (instead of adding it).
	Delete *bool `json:"delete"`
	// List of storage IDs.
	Storage *string `json:"storage"`
	// List of virtual machines.
	Vms *string `json:"vms"`
}

type PutPoolResponse struct {
}

func (c *Client) PutPool(Poolid string, parameters PutPoolParameters) (*PutPoolResponse, error) {

}

type GetVersionResponse struct {
}

func (c *Client) GetVersion() (*GetVersionResponse, error) {

}
