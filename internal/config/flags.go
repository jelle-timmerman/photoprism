package config

import (
	"github.com/klauspost/cpuid/v2"
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/face"
	"github.com/photoprism/photoprism/internal/i18n"
)

// GlobalFlags describes global command-line parameters and flags.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "admin-password",
		Usage:  "initial admin `PASSWORD`, minimum 4 characters",
		EnvVar: "PHOTOPRISM_ADMIN_PASSWORD",
	},
	cli.StringFlag{
		Name:   "log-level, l",
		Usage:  "trace, debug, info, warning, error, fatal, or panic",
		Value:  "info",
		EnvVar: "PHOTOPRISM_LOG_LEVEL",
	},
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "enable debug mode, show additional log messages",
		EnvVar: "PHOTOPRISM_DEBUG",
	},
	cli.BoolFlag{
		Name:   "test",
		Hidden: true,
		Usage:  "enable test mode",
	},
	cli.BoolFlag{
		Name:   "unsafe",
		Hidden: true,
		Usage:  "enable unsafe mode",
		EnvVar: "PHOTOPRISM_UNSAFE",
	},
	cli.BoolFlag{
		Name:   "demo",
		Hidden: true,
		Usage:  "enable demo mode",
		EnvVar: "PHOTOPRISM_DEMO",
	},
	cli.BoolFlag{
		Name:   "sponsor",
		Hidden: true,
		Usage:  "your continuous support helps to pay for development and operating expenses",
		EnvVar: "PHOTOPRISM_SPONSOR",
	},
	cli.BoolFlag{
		Name:   "public, p",
		Usage:  "disable password authentication, WebDAV, and the advanced settings page",
		EnvVar: "PHOTOPRISM_PUBLIC",
	},
	cli.BoolFlag{
		Name:   "read-only, r",
		Usage:  "disable import, upload, delete, and all other operations that require write permissions",
		EnvVar: "PHOTOPRISM_READONLY",
	},
	cli.BoolFlag{
		Name:   "experimental, e",
		Usage:  "enable experimental features",
		EnvVar: "PHOTOPRISM_EXPERIMENTAL",
	},
	cli.StringFlag{
		Name:   "partner-id",
		Hidden: true,
		Usage:  "hosting partner id",
		EnvVar: "PHOTOPRISM_PARTNER_ID",
	},
	cli.StringFlag{
		Name:   "config-file, c",
		Usage:  "load config options from `FILENAME`",
		EnvVar: "PHOTOPRISM_CONFIG_FILE",
	},
	cli.StringFlag{
		Name:   "config-path",
		Usage:  "config `PATH` to be searched for additional configuration and settings files",
		EnvVar: "PHOTOPRISM_CONFIG_PATH",
	},
	cli.StringFlag{
		Name:   "originals-path",
		Usage:  "storage `PATH` of your original media files (photos and videos)",
		EnvVar: "PHOTOPRISM_ORIGINALS_PATH",
	},
	cli.IntFlag{
		Name:   "originals-limit",
		Value:  1000,
		Usage:  "file size limit in `MB`",
		EnvVar: "PHOTOPRISM_ORIGINALS_LIMIT",
	},
	cli.StringFlag{
		Name:   "storage-path",
		Usage:  "writable storage `PATH` for cache, database, and sidecar files",
		EnvVar: "PHOTOPRISM_STORAGE_PATH",
	},
	cli.StringFlag{
		Name:   "import-path",
		Usage:  "base `PATH` from which files can be imported to originals (optional)",
		EnvVar: "PHOTOPRISM_IMPORT_PATH",
	},
	cli.StringFlag{
		Name:   "cache-path",
		Usage:  "custom cache `PATH` for sessions and thumbnail files (optional)",
		EnvVar: "PHOTOPRISM_CACHE_PATH",
	},
	cli.StringFlag{
		Name:   "sidecar-path",
		Usage:  "custom relative or absolute sidecar `PATH` (optional)",
		EnvVar: "PHOTOPRISM_SIDECAR_PATH",
	},
	cli.StringFlag{
		Name:   "temp-path",
		Usage:  "custom temporary file `PATH` (optional)",
		EnvVar: "PHOTOPRISM_TEMP_PATH",
	},
	cli.StringFlag{
		Name:   "backup-path",
		Usage:  "custom backup `PATH` for index backup files (optional)",
		EnvVar: "PHOTOPRISM_BACKUP_PATH",
	},
	cli.StringFlag{
		Name:   "assets-path",
		Usage:  "assets `PATH` containing static resources like icons, models, and translations",
		EnvVar: "PHOTOPRISM_ASSETS_PATH",
	},
	cli.IntFlag{
		Name:   "workers, w",
		Usage:  "maximum `NUMBER` of indexing workers, default depends on the number of physical cores",
		EnvVar: "PHOTOPRISM_WORKERS",
		Value:  cpuid.CPU.PhysicalCores / 2,
	},
	cli.IntFlag{
		Name:   "wakeup-interval",
		Usage:  "metadata, share & sync background worker wakeup interval in `SECONDS` (1-604800)",
		Value:  DefaultWakeupIntervalSeconds,
		EnvVar: "PHOTOPRISM_WAKEUP_INTERVAL",
	},
	cli.IntFlag{
		Name:   "auto-index",
		Usage:  "WebDAV auto index safety delay in `SECONDS`, disable with -1",
		Value:  DefaultAutoIndexDelay,
		EnvVar: "PHOTOPRISM_AUTO_INDEX",
	},
	cli.IntFlag{
		Name:   "auto-import",
		Usage:  "WebDAV auto import safety delay in `SECONDS`, disable with -1",
		Value:  DefaultAutoImportDelay,
		EnvVar: "PHOTOPRISM_AUTO_IMPORT",
	},
	cli.BoolFlag{
		Name:   "disable-webdav",
		Usage:  "disable built-in WebDAV server",
		EnvVar: "PHOTOPRISM_DISABLE_WEBDAV",
	},
	cli.BoolFlag{
		Name:   "disable-settings",
		Usage:  "disable settings UI and API",
		EnvVar: "PHOTOPRISM_DISABLE_SETTINGS",
	},
	cli.BoolFlag{
		Name:   "disable-places",
		Usage:  "disable reverse geocoding and maps",
		EnvVar: "PHOTOPRISM_DISABLE_PLACES",
	},
	cli.BoolFlag{
		Name:   "disable-backups",
		Usage:  "disable creating YAML metadata files",
		EnvVar: "PHOTOPRISM_DISABLE_BACKUPS",
	},
	cli.BoolFlag{
		Name:   "disable-exiftool",
		Usage:  "disable creating JSON metadata sidecar files with ExifTool",
		EnvVar: "PHOTOPRISM_DISABLE_EXIFTOOL",
	},
	cli.BoolFlag{
		Name:   "disable-ffmpeg",
		Usage:  "disable video transcoding and thumbnail extraction with FFmpeg",
		EnvVar: "PHOTOPRISM_DISABLE_FFMPEG",
	},
	cli.BoolFlag{
		Name:   "disable-darktable",
		Usage:  "disable converting RAW files with Darktable",
		EnvVar: "PHOTOPRISM_DISABLE_DARKTABLE",
	},
	cli.BoolFlag{
		Name:   "disable-rawtherapee",
		Usage:  "disable converting RAW files with RawTherapee",
		EnvVar: "PHOTOPRISM_DISABLE_RAWTHERAPEE",
	},
	cli.BoolFlag{
		Name:   "disable-sips",
		Usage:  "disable converting RAW files with Sips (macOS only)",
		EnvVar: "PHOTOPRISM_DISABLE_SIPS",
	},
	cli.BoolFlag{
		Name:   "disable-heifconvert",
		Usage:  "disable converting HEIC/HEIF files",
		EnvVar: "PHOTOPRISM_DISABLE_HEIFCONVERT",
	},
	cli.BoolFlag{
		Name:   "disable-tensorflow",
		Usage:  "disable all features depending on TensorFlow",
		EnvVar: "PHOTOPRISM_DISABLE_TENSORFLOW",
	},
	cli.BoolFlag{
		Name:   "disable-faces",
		Usage:  "disable facial recognition",
		EnvVar: "PHOTOPRISM_DISABLE_FACES",
	},
	cli.BoolFlag{
		Name:   "disable-classification",
		Usage:  "disable image classification",
		EnvVar: "PHOTOPRISM_DISABLE_CLASSIFICATION",
	},
	cli.BoolFlag{
		Name:   "detect-nsfw",
		Usage:  "flag photos as private that may be offensive (requires TensorFlow)",
		EnvVar: "PHOTOPRISM_DETECT_NSFW",
	},
	cli.BoolFlag{
		Name:   "upload-nsfw",
		Usage:  "allow uploads that may be offensive",
		EnvVar: "PHOTOPRISM_UPLOAD_NSFW",
	},
	cli.StringFlag{
		Name:   "default-theme",
		Usage:  "standard user interface theme `NAME`",
		Hidden: true,
		EnvVar: "PHOTOPRISM_DEFAULT_THEME",
	},
	cli.StringFlag{
		Name:   "default-locale",
		Usage:  "standard user interface language `CODE`",
		Value:  i18n.Default.Locale(),
		EnvVar: "PHOTOPRISM_DEFAULT_LOCALE",
	},
	cli.StringFlag{
		Name:   "app-icon",
		Usage:  "web app `ICON` (logo, app, crisp, mint, bold)",
		EnvVar: "PHOTOPRISM_APP_ICON",
	},
	cli.StringFlag{
		Name:   "app-name",
		Usage:  "web app `NAME` when installed on a device",
		Value:  "PhotoPrism",
		EnvVar: "PHOTOPRISM_APP_NAME",
	},
	cli.StringFlag{
		Name:   "app-mode",
		Usage:  "web app `MODE` (fullscreen, standalone, minimal-ui, browser)",
		Value:  "standalone",
		EnvVar: "PHOTOPRISM_APP_MODE",
	},
	cli.StringFlag{
		Name:   "cdn-url",
		Usage:  "content delivery network `URL` (optional)",
		EnvVar: "PHOTOPRISM_CDN_URL",
	},
	cli.StringFlag{
		Name:   "site-url",
		Usage:  "public site `URL`",
		Value:  "http://localhost:2342/",
		EnvVar: "PHOTOPRISM_SITE_URL",
	},
	cli.StringFlag{
		Name:   "site-author",
		Usage:  "site `OWNER`, copyright, or artist",
		EnvVar: "PHOTOPRISM_SITE_AUTHOR",
	},
	cli.StringFlag{
		Name:   "site-title",
		Usage:  "site `TITLE`",
		Value:  "PhotoPrism",
		EnvVar: "PHOTOPRISM_SITE_TITLE",
	},
	cli.StringFlag{
		Name:   "site-caption",
		Usage:  "site `CAPTION`",
		Value:  "Browse Your Life",
		EnvVar: "PHOTOPRISM_SITE_CAPTION",
	},
	cli.StringFlag{
		Name:   "site-description",
		Usage:  "site `DESCRIPTION` (optional)",
		EnvVar: "PHOTOPRISM_SITE_DESCRIPTION",
	},
	cli.StringFlag{
		Name:   "site-preview",
		Usage:  "site preview image `URL` (optional)",
		EnvVar: "PHOTOPRISM_SITE_PREVIEW",
	},
	cli.StringFlag{
		Name:   "imprint",
		Usage:  "legal `INFO`, displayed in the page footer",
		Value:  "",
		EnvVar: "PHOTOPRISM_IMPRINT",
	},
	cli.StringFlag{
		Name:   "imprint-url",
		Usage:  "legal info `URL` (optional)",
		Value:  "",
		EnvVar: "PHOTOPRISM_IMPRINT_URL",
	},
	cli.IntFlag{
		Name:   "http-port",
		Value:  2342,
		Usage:  "http server port `NUMBER`",
		EnvVar: "PHOTOPRISM_HTTP_PORT",
	},
	cli.StringFlag{
		Name:   "http-host",
		Usage:  "http server `IP` address",
		EnvVar: "PHOTOPRISM_HTTP_HOST",
	},
	cli.StringFlag{
		Name:   "http-mode, m",
		Usage:  "http server `MODE` (debug, release, or test)",
		EnvVar: "PHOTOPRISM_HTTP_MODE",
	},
	cli.StringFlag{
		Name:   "http-compression, z",
		Usage:  "http server compression `METHOD` (none or gzip)",
		EnvVar: "PHOTOPRISM_HTTP_COMPRESSION",
	},
	cli.StringFlag{
		Name:   "database-driver",
		Usage:  "database `DRIVER` (sqlite or mysql)",
		Value:  "sqlite",
		EnvVar: "PHOTOPRISM_DATABASE_DRIVER",
	},
	cli.StringFlag{
		Name:   "database-dsn",
		Usage:  "sqlite file name, providing a `DSN` is optional for other drivers",
		EnvVar: "PHOTOPRISM_DATABASE_DSN",
	},
	cli.StringFlag{
		Name:   "database-server",
		Usage:  "database server `HOST` with optional port e.g. mariadb:3306",
		EnvVar: "PHOTOPRISM_DATABASE_SERVER",
	},
	cli.StringFlag{
		Name:   "database-name",
		Value:  "photoprism",
		Usage:  "database schema `NAME`",
		EnvVar: "PHOTOPRISM_DATABASE_NAME",
	},
	cli.StringFlag{
		Name:   "database-user",
		Value:  "photoprism",
		Usage:  "database user `NAME`",
		EnvVar: "PHOTOPRISM_DATABASE_USER",
	},
	cli.StringFlag{
		Name:   "database-password",
		Usage:  "database user `PASSWORD`",
		EnvVar: "PHOTOPRISM_DATABASE_PASSWORD",
	},
	cli.IntFlag{
		Name:   "database-conns",
		Usage:  "maximum `NUMBER` of open database connections",
		EnvVar: "PHOTOPRISM_DATABASE_CONNS",
	},
	cli.IntFlag{
		Name:   "database-conns-idle",
		Usage:  "maximum `NUMBER` of idle database connections",
		EnvVar: "PHOTOPRISM_DATABASE_CONNS_IDLE",
	},
	cli.BoolFlag{
		Name:   "raw-presets",
		Usage:  "enable RAW file converter presets (may reduce performance)",
		EnvVar: "PHOTOPRISM_RAW_PRESETS",
	},
	cli.StringFlag{
		Name:   "darktable-bin",
		Usage:  "Darktable CLI `COMMAND` for RAW image conversion",
		Value:  "darktable-cli",
		EnvVar: "PHOTOPRISM_DARKTABLE_BIN",
	},
	cli.StringFlag{
		Name:   "darktable-blacklist",
		Usage:  "file `EXTENSIONS` incompatible with Darktable",
		Value:  "dng",
		EnvVar: "PHOTOPRISM_DARKTABLE_BLACKLIST",
	},
	cli.StringFlag{
		Name:   "rawtherapee-bin",
		Usage:  "RawTherapee CLI `COMMAND` for RAW image conversion",
		Value:  "rawtherapee-cli",
		EnvVar: "PHOTOPRISM_RAWTHERAPEE_BIN",
	},
	cli.StringFlag{
		Name:   "rawtherapee-blacklist",
		Usage:  "file `EXTENSIONS` incompatible with RawTherapee",
		Value:  "",
		EnvVar: "PHOTOPRISM_RAWTHERAPEE_BLACKLIST",
	},
	cli.StringFlag{
		Name:   "sips-bin",
		Usage:  "Sips `COMMAND` for RAW image conversion (macOS only)",
		Value:  "sips",
		EnvVar: "PHOTOPRISM_SIPS_BIN",
	},
	cli.StringFlag{
		Name:   "heifconvert-bin",
		Usage:  "HEIC/HEIF image convert `COMMAND`",
		Value:  "heif-convert",
		EnvVar: "PHOTOPRISM_HEIFCONVERT_BIN",
	},
	cli.StringFlag{
		Name:   "ffmpeg-bin",
		Usage:  "FFmpeg `COMMAND` for video transcoding and thumbnail extraction",
		Value:  "ffmpeg",
		EnvVar: "PHOTOPRISM_FFMPEG_BIN",
	},
	cli.StringFlag{
		Name:   "ffmpeg-encoder",
		Usage:  "FFmpeg AVC encoder `NAME`",
		Value:  "libx264",
		EnvVar: "PHOTOPRISM_FFMPEG_ENCODER",
	},
	cli.IntFlag{
		Name:   "ffmpeg-bitrate",
		Usage:  "maximum FFmpeg encoding `BITRATE` (Mbit/s)",
		Value:  50,
		EnvVar: "PHOTOPRISM_FFMPEG_BITRATE",
	},
	cli.IntFlag{
		Name:   "ffmpeg-buffers",
		Usage:  "`NUMBER` of FFmpeg capture buffers",
		Value:  32,
		EnvVar: "PHOTOPRISM_FFMPEG_BUFFERS",
	},
	cli.StringFlag{
		Name:   "exiftool-bin",
		Usage:  "ExifTool `COMMAND` for extracting metadata",
		Value:  "exiftool",
		EnvVar: "PHOTOPRISM_EXIFTOOL_BIN",
	},
	cli.StringFlag{
		Name:   "download-token",
		Usage:  "`SECRET` download URL token for originals (default: random)",
		EnvVar: "PHOTOPRISM_DOWNLOAD_TOKEN",
	},
	cli.StringFlag{
		Name:   "preview-token",
		Usage:  "`SECRET` thumbnail and video streaming URL token (default: random)",
		EnvVar: "PHOTOPRISM_PREVIEW_TOKEN",
	},
	cli.StringFlag{
		Name:   "thumb-filter",
		Usage:  "thumbnail downscaling `FILTER` (best to worst: blackman, lanczos, cubic, linear)",
		Value:  "lanczos",
		EnvVar: "PHOTOPRISM_THUMB_FILTER",
	},
	cli.IntFlag{
		Name:   "thumb-size, s",
		Usage:  "maximum pre-cached thumbnail image size in `PIXELS` (720-7680)",
		Value:  2048,
		EnvVar: "PHOTOPRISM_THUMB_SIZE",
	},
	cli.BoolFlag{
		Name:   "thumb-uncached, u",
		Usage:  "enable on-demand thumbnail generation (high memory and cpu usage)",
		EnvVar: "PHOTOPRISM_THUMB_UNCACHED",
	},
	cli.IntFlag{
		Name:   "thumb-size-uncached, x",
		Usage:  "maximum size of on-demand generated thumbnails in `PIXELS` (720-7680)",
		Value:  7680,
		EnvVar: "PHOTOPRISM_THUMB_SIZE_UNCACHED",
	},
	cli.IntFlag{
		Name:   "jpeg-size",
		Usage:  "maximum size of generated JPEG images in `PIXELS` (720-30000)",
		Value:  7680,
		EnvVar: "PHOTOPRISM_JPEG_SIZE",
	},
	cli.IntFlag{
		Name:   "jpeg-quality, q",
		Usage:  "`QUALITY` of generated JPEG images, a higher value reduces compression (25-100)",
		Value:  92,
		EnvVar: "PHOTOPRISM_JPEG_QUALITY",
	},
	cli.IntFlag{
		Name:   "face-size",
		Usage:  "minimum face size in `PIXELS`",
		Value:  face.SizeThreshold,
		EnvVar: "PHOTOPRISM_FACE_SIZE",
	},
	cli.Float64Flag{
		Name:   "face-score",
		Usage:  "minimum face `QUALITY` score",
		Value:  face.ScoreThreshold,
		EnvVar: "PHOTOPRISM_FACE_SCORE",
	},
	cli.IntFlag{
		Name:   "face-overlap",
		Usage:  "face area overlap threshold in `PERCENT`",
		Value:  face.OverlapThreshold,
		EnvVar: "PHOTOPRISM_FACE_OVERLAP",
	},
	cli.IntFlag{
		Name:   "face-cluster-size",
		Usage:  "minimum size of automatically clustered faces in `PIXELS`",
		Value:  face.ClusterSizeThreshold,
		EnvVar: "PHOTOPRISM_FACE_CLUSTER_SIZE",
	},
	cli.IntFlag{
		Name:   "face-cluster-score",
		Usage:  "minimum `QUALITY` score of automatically clustered faces",
		Value:  face.ClusterScoreThreshold,
		EnvVar: "PHOTOPRISM_FACE_CLUSTER_SCORE",
	},
	cli.IntFlag{
		Name:   "face-cluster-core",
		Usage:  "`NUMBER` of faces forming a cluster core",
		Value:  face.ClusterCore,
		EnvVar: "PHOTOPRISM_FACE_CLUSTER_CORE",
	},
	cli.Float64Flag{
		Name:   "face-cluster-dist",
		Usage:  "similarity `DISTANCE` of faces forming a cluster core",
		Value:  face.ClusterDist,
		EnvVar: "PHOTOPRISM_FACE_CLUSTER_DIST",
	},
	cli.Float64Flag{
		Name:   "face-match-dist",
		Usage:  "similarity `OFFSET` for matching faces with existing clusters",
		Value:  face.MatchDist,
		EnvVar: "PHOTOPRISM_FACE_MATCH_DIST",
	},
	cli.StringFlag{
		Name:   "pid-filename",
		Usage:  "process id `FILENAME` (daemon mode only)",
		EnvVar: "PHOTOPRISM_PID_FILENAME",
	},
	cli.StringFlag{
		Name:   "log-filename",
		Usage:  "server log `FILENAME` (daemon mode only)",
		EnvVar: "PHOTOPRISM_LOG_FILENAME",
		Value:  "",
	},
}
