package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const debugArgName = "debug"

func InitLog() {
	if viper.GetBool(debugArgName) {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
		logrus.Debug("已开启debug模式...")
		logrus.Debugf("config:%v",Instance)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	Instance.Debug = viper.GetBool(debugArgName)
}

func BindParameter(cmd *cobra.Command) {
	viper.SetEnvPrefix("authz")
	viper.AutomaticEnv()

	cmd.PersistentFlags().BoolVarP(&Instance.Debug, debugArgName, "v", false, "debug mod")
	cmd.PersistentFlags().BoolVarP(&Instance.Allow, "allow", "a", true, "是否允许")

	//_ = viper.BindPFlag(debugArgName, cmd.PersistentFlags().Lookup(debugArgName))
	//
	//_ = viper.BindPFlag("mongo-address", cmd.PersistentFlags().Lookup("mongo-address"))
	//_ = viper.BindPFlag("mongo-port", cmd.PersistentFlags().Lookup("mongo-port"))
	//_ = viper.BindPFlag("mongo-Username", cmd.PersistentFlags().Lookup("mongo-Username"))
	//_ = viper.BindPFlag("mongo-Password", cmd.PersistentFlags().Lookup("mongo-Password"))
	//_ = viper.BindPFlag("mongo-LocalThreshold", cmd.PersistentFlags().Lookup("mongo-LocalThreshold"))
	//_ = viper.BindPFlag("mongo-MaxPoolSize", cmd.PersistentFlags().Lookup("mongo-MaxPoolSize"))
	//_ = viper.BindPFlag("mongo-MaxConnIdleTime", cmd.PersistentFlags().Lookup("mongo-MaxConnIdleTime"))
	//_ = viper.BindPFlag("mongo-DbName", cmd.PersistentFlags().Lookup("mongo-DbName"))
	//_ = viper.BindPFlag("mongo-EventCollectionName", cmd.PersistentFlags().Lookup("mongo-EventCollectionName"))
	//_ = viper.BindPFlag("mongo-SnapshotCollectionName", cmd.PersistentFlags().Lookup("mongo-SnapshotCollectionName"))
	_ = viper.BindPFlags(cmd.Flags())
}

type Config struct {
	Debug bool
	Allow bool
}

var Instance = &Config{
}
