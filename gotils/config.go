package gotils

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gopkg.in/gcfg.v1"
)

var (
	//Version is a linker injected variable for a git revision info used as version info
	Version = "Unknown build"
	/*Buildstamp is a linker injected variable for a buildtime timestamp used in version info */
	Buildstamp = "unknown build timestamp."
)

// Config contains the configuration data.
type Config struct {
	DataBase struct {
		Database string
		User     string
		Password string
	}
	Log struct {
		LogLevel       string
		ActiveLogLevel LogLevel
		LogFileName    string
		Logger         *Logger
	}
	ConfigFile string

	BuildTimeStamp time.Time
	GitVersion     string
}

// Flags ...
type Flags struct {
	ConfigFile  string
	Database    string
	User        string
	Password    string
	LogLevel    string
	LogFileName string
}

func commandLineParsing(flags *Flags) {

	flag.StringVar(&flags.ConfigFile, "Ini", "shorty.ini", "The configfile to read parameters from.)")
	flag.StringVar(&flags.Database, "Db", "", "The database name to connect to.")
	flag.StringVar(&flags.User, "User", "", "The database user to use for the db connection.")
	flag.StringVar(&flags.Password, "Password", "", "The database users password.")
	flag.StringVar(&flags.LogLevel, "LogLevel", "All", "Determines logging verbosity. [Off|Info|Debug|Warnings|Error|Fatal|All].")
	flag.StringVar(&flags.LogFileName, "LogFile", "", "Sets the name of the logfile. Uses STDOUT if empty.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n"+
			"  %s [options] \n\n"+
			"The commandline options are overruling environment variales, that are overriding the configfile "+
			"settings. You may use the flag names with te the prefix 'SHORTY' as environment variables. \n"+
			"Example: To set the database user you may set the variable 'SHORTYDb=demouser'\n\n"+
			"Example configuration file:\n<snip>\n "+
			"   ; Diana Config file\n"+
			"   [Database]\n"+
			"     Database = shortydb\n"+
			"     User     = mabuse\n"+
			"     Password = fkhdb4322rb\n"+
			"   [Log]\n"+
			"     LogLevel = Debug\n"+
			"     LogFileName = shorty.log\n</snip>\n"+
			"Options:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

// isFlagPassed determines if a flag was set on commandline or if the flag has its value set by default.
// Upgly "brute force" run, but flags are a of a very low number so it does not matter.
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// configPrecedence sets a configs value using the following priorities from low to high:
// defaults, config file, environment variables, flags
func setConfigByPrecedence(place *string, key string, cfg string) {
	f := flag.Lookup(key)
	if f != nil {
		if isFlagPassed(key) {
			*place = f.Value.String()
		} else {
			envvar := "$SHORTY" + key
			if "" != os.ExpandEnv(envvar) {
				*place = os.ExpandEnv(envvar)
			} else if cfg != "" {
				*place = cfg
			} else {
				*place = f.DefValue
			}
		}
	}
}

// ReadConfig ...
func ReadConfig(cfg *Config, supersedefilename ...string) {

	var flags = new(Flags)
	commandLineParsing(flags)
	setConfigByPrecedence(&cfg.ConfigFile, "Ini", "")
	if len(supersedefilename) > 0 {
		cfg.ConfigFile = supersedefilename[0]
	}

	LogDebug("Reading config file '%s'.", cfg.ConfigFile)
	err := gcfg.ReadFileInto(cfg, cfg.ConfigFile)
	if err != nil {
		err := gcfg.ReadFileInto(cfg, fmt.Sprintf("%s", cfg.ConfigFile))
		if err != nil {
			LogFatal(err.Error())
		}
	}

	if flags.LogLevel != "" {
		setConfigByPrecedence(&cfg.Log.LogLevel, "LogLevel", cfg.Log.LogLevel)
		cfg.Log.ActiveLogLevel, err = LogLevelString(cfg.Log.LogLevel)
		if err != nil {
			LogError("Error in config, Loglevel '%s' not existing in tools/loglevel.go. Setting LogLevel to 'All'", cfg.Log.LogLevel)
			cfg.Log.ActiveLogLevel = All
			cfg.Log.LogLevel = "All"
		}
	}

	setConfigByPrecedence(&cfg.Log.LogFileName, "LogFile", cfg.Log.LogFileName)
	cfg.Log.Logger = NewLogger(cfg.Log.LogFileName, cfg.Log.ActiveLogLevel)
	cfg.Log.Logger.SetConvenienceLogger()

	setConfigByPrecedence(&cfg.DataBase.Database, "Db", cfg.DataBase.Database)
	setConfigByPrecedence(&cfg.DataBase.User, "User", cfg.DataBase.User)
	setConfigByPrecedence(&cfg.DataBase.Password, "Password", cfg.DataBase.Password)

	btime, err := time.Parse("2006-01-02_15:04:05_MST", Buildstamp)
	if err != nil {
		cfg.BuildTimeStamp = time.Now()
	} else {
		cfg.BuildTimeStamp = btime
	}
	cfg.GitVersion = Version
	cfg.Log.Logger.Info("Version: %s of %s \n", cfg.GitVersion, cfg.BuildTimeStamp)

}
