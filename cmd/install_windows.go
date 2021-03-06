package cmd

import (
	"fmt"
	"strconv"

	"github.com/drakkan/sftpgo/service"
	"github.com/drakkan/sftpgo/utils"
	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install SFTPGo as Windows Service",
		Long: `To install the SFTPGo Windows Service with the default values for the command line flags simply use:

sftpgo service install

Please take a look at the usage below to customize the startup options`,
		Run: func(cmd *cobra.Command, args []string) {
			s := service.Service{
				ConfigDir:     utils.CleanDirInput(configDir),
				ConfigFile:    configFile,
				LogFilePath:   logFilePath,
				LogMaxSize:    logMaxSize,
				LogMaxBackups: logMaxBackups,
				LogMaxAge:     logMaxAge,
				LogCompress:   logCompress,
				LogVerbose:    logVerbose,
				Shutdown:      make(chan bool),
			}
			winService := service.WindowsService{
				Service: s,
			}
			serviceArgs := []string{"service", "start"}
			customFlags := getCustomServeFlags()
			if len(customFlags) > 0 {
				serviceArgs = append(serviceArgs, customFlags...)
			}
			err := winService.Install(serviceArgs...)
			if err != nil {
				fmt.Printf("Error installing service: %v\r\n", err)
			} else {
				fmt.Printf("Service installed!\r\n")
			}
		},
	}
)

func init() {
	serviceCmd.AddCommand(installCmd)
	addServeFlags(installCmd)
}

func getCustomServeFlags() []string {
	result := []string{}
	if configDir != defaultConfigDir {
		configDir = utils.CleanDirInput(configDir)
		result = append(result, "--"+configDirFlag)
		result = append(result, configDir)
	}
	if configFile != defaultConfigName {
		result = append(result, "--"+configFileFlag)
		result = append(result, configFile)
	}
	if logFilePath != defaultLogFile {
		result = append(result, "--"+logFilePathFlag)
		result = append(result, logFilePath)
	}
	if logMaxSize != defaultLogMaxSize {
		result = append(result, "--"+logMaxSizeFlag)
		result = append(result, strconv.Itoa(logMaxSize))
	}
	if logMaxBackups != defaultLogMaxBackup {
		result = append(result, "--"+logMaxBackupFlag)
		result = append(result, strconv.Itoa(logMaxBackups))
	}
	if logMaxAge != defaultLogMaxAge {
		result = append(result, "--"+logMaxAgeFlag)
		result = append(result, strconv.Itoa(logMaxAge))
	}
	if logVerbose != defaultLogVerbose {
		result = append(result, "--"+logVerboseFlag+"=false")
	}
	if logCompress != defaultLogCompress {
		result = append(result, "--"+logCompressFlag+"=true")
	}
	if profiler != defaultProfiler {
		result = append(result, "--"+profilerFlag+"=true")
	}
	return result
}
