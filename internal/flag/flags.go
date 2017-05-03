package flag

import "flag"
import "strings"

const (
	//FLAG_SERVER_PORT_NAME httpmock服务启动端口参数名称，默认是10105
	FLAG_SERVER_PORT_NAME          = "port"
	FLAG_SERVER_PORT_DEFAULT_VALUE = "10105"

	//FLAG_SERVER_BIND_NAME httpmock服务启动绑定ip参数名称，默认是0.0.0.0
	FLAG_SERVER_BIND_NAME          = "ip"
	FLAG_SERVER_BIND_DEFAULT_VALUE = "0.0.0.0"

	//FLAG_SERVER_VERSION_NAME httpmock版本号
	FLAG_SERVER_VERSION_NAME  = "version"
	FLAG_SERVER_VERSION_VALUE = false

	//FLAG_SERVER_CONFIG_NAME httpmock服务启动配置文件路径，默认是conf/api_gateway.ini
	FLAG_SERVER_CONFIG_NAME  = "config"
	FLAG_SERVER_CONFIG_VALUE = "conf/devel_httpmock.ini"

	//FLAG_ENVIRONMENT_NAME 启动选择环境，默认是devel
	FLAG_ENVIRONMENT_NAME          = "env"
	FLAG_ENVIRONMENT_DEFAULT_VALUE = "devel"
)

type FlagInit struct {
	ConfigPath     string
	ServiceAddress string
	ServicePort    string
	Adapters       string
}

/*初始化程序运行参数
  当返回值is_version为真时，表示查询程序版本号,并记录version的相关信息；其他情况isVersion为false
*/
func InitFlags() (bool, *FlagInit, error) {
	//启动参数读取
	version := flag.Bool(FLAG_SERVER_VERSION_NAME, FLAG_SERVER_VERSION_VALUE, "查询运行httpmock版本号")
	port := flag.String(FLAG_SERVER_PORT_NAME, FLAG_SERVER_PORT_DEFAULT_VALUE, "httpmock服务启动端口")
	address := flag.String(FLAG_SERVER_BIND_NAME, FLAG_SERVER_BIND_DEFAULT_VALUE, "httpmock服务绑定启动IP地址")
	configuration := flag.String(FLAG_SERVER_CONFIG_NAME, FLAG_SERVER_CONFIG_VALUE, "指定配置文件路径")
	env := flag.String(FLAG_ENVIRONMENT_NAME, FLAG_ENVIRONMENT_DEFAULT_VALUE, "选择启动的环境(devel、test、pre、prod)")
	flag.Parse()

	if *version {
		return true, nil, nil
	}

	if err := checkflags(*port, *address, *configuration, *env); err != nil {
		return false, nil, err
	}

	if strings.TrimSpace(*env) != "" {
		*configuration = strings.Replace(*configuration, "devel", *env, 1)
	}

	flagInit := &FlagInit{
		ConfigPath:     *configuration,
		ServiceAddress: *address,
		ServicePort:    *port,
	}
	return false, flagInit, nil
}

func checkflags(port, address, configuration, env string) error {
	var err error
	//判断端口是否符合规范

	//判断address是否符合规范

	//判断configuration文件是否存在

	return err
}
