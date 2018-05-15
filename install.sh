#!/bin/bash

cli_url=https://github.com/ingrammicro/concerto/raw/master/binaries/concerto.x64
api_url=${CONCERTO_ENDPOINT:=https://clients.IMCO_DOMAIN:886/}
cli_command=concerto
cli_fullpath=/usr/local/bin/$cli_command
conf_path=$HOME/.concerto
cli_conf=$conf_path/client.xml
cli_conf_exists=false
cli_fullpath_exists=false
cacert_exists=false
cert_exists=false
key_exists=false
force_bin=false
force_keys=false
force_conf=false
verbose=false
LOGO_H_SIZE=127

parseArgs(){
	printf "Parse arguments ..."

	for arg in "$@"
	do
    if [ "$arg" == "fb" ]; then force_bin=true; fi
		if [ "$arg" == "fk" ]; then force_keys=true; fi
		if [ "$arg" == "fc" ]; then force_conf=true; fi
		if [ "$arg" == "f" ]; then force_bin=true; force_keys=true; force_conf=true; fi
		if [ "$arg" == "v" ]; then verbose=true; fi
	done
	printf " OK\n"
}

initializeIMCOCLI(){
	printf "Initializing ..."

	case "$(uname -m)" in
		*64)
			;;
		*)
			echo >&2 -e "IMCO CLI is available for 64 bit systems only\n"
			exit 1
			;;
	esac

	case "$(uname -s)" in
		Darwin)
			cli_url="$cli_url.darwin"
			;;
		Linux)
			cli_url="$cli_url.linux"
			;;
		*)
			$verbose && printf " (OS could not be detected. Assuming linux) "
			cli_url="$cli_url.linux"
			;;
	esac

	[ ]
	printf " OK\n"

	getInstallationState
}

getInstallationState(){
	[ -f $cli_conf ] && cli_conf_exists=true || cli_conf_exists=false
	[ -f $cli_fullpath ] && cli_fullpath_exists=true || cli_fullpath_exists=false
	[ -f $conf_path/ssl/ca_cert.pem ] && cacert_exists=true || cacert_exists=false
	[ -f $conf_path/ssl/cert.crt ] && cert_exists=true || cert_exists=false
	[ -f $conf_path/ssl/private/cert.key ] && key_exists=true || key_exists=false
}

writeDefaultConfig(){
	printf "Writing IMCO CLI configuration ..."

	if ! $force_conf && $cli_conf_exists;
	then
		$verbose && printf " (configuration found at '$cli_conf'. Use '-fc' to force update binary)"
		printf " Skipped\n"
		return
	fi

	mkdir -p "${conf_path}"
	cat <<EOF > $cli_conf
<concerto version="1.0" server="$api_url" log_file="/var/log/concerto-client.log" log_level="info">
	<ssl cert="$conf_path/ssl/cert.crt" key="$conf_path/ssl/private/cert.key" server_ca="$conf_path/ssl/ca_cert.pem" />
</concerto>
EOF

	printf " OK\n"
}

installIMCOCLI(){
	printf "Installing IMCO CLI ..."

	if ! $force_bin && $cli_fullpath_exists;
	then
		$verbose && printf " (IMCO CLI exists. Use '-fb' to force update binary)"
		printf " Skipped\n"
		return
	fi

	command -v curl > /dev/null &&  dwld="curl -sSL -o" || \
	{ command -v wget > /dev/null && dwld="wget -qO"; } || \
	{ echo ' (curl or wget are needed to install IMCO CLI.) Failed'; exit 1; }
	printf " (you might be asked for your password to sudo now)\n"
	if ! sudo $dwld  $cli_fullpath $cli_url;
	then
		echo "(IMCO CLI Binary download failed). Failed"
		exit 1
	fi

	if ! sudo chmod +x $cli_fullpath;
	then
		echo "(IMCO CLI Binary execution flag assigment failed). Failed"
		exit 1
	fi

	echo "Binary has been installed. OK"

	current_concerto=$(command -v $cli_fullpath)
	[ $current_concerto != $cli_fullpath ] && echo "WARNING: IMCO CLI is being run from '$current_concerto'. Please, update your path to execute from $cli_fullpath"

}

installAPIKeys(){
	printf "Installing API keys ..."

	if ! $force_keys && $cacert_exists && $cert_exists && $key_exists;
	then
	 	$verbose && printf " (IMCO CLI keys already exists. Use '-fk' to force update API keys)"
		printf " Skipped\n"
	else
		echo
		#concerto setup api_keys < /dev/tty

		if [ $? -ne 0 ];
		then
			printf " (error downloading IMCO CLI keys. Try downloading manually). Failed\n"
			certsInstructions
			exit 1
		fi
	fi

	# # if certs not there
	#  ! $cacert_exists || ! $cert_exists || ! $key_exists ] && ! $cert_exists && concerto setup api_keys
	#  getInstallationState
	#  ! $cacert_exists || ! $cert_exists || ! $cert_exists ] && ! $cert_exists && certsInstructions || echo "IMCO CLI installed. Type 'concerto' to access CLI help"

}

certsInstructions(){
cat <<EOF
IMCO CLI uses an API Key that you can download from IMCO's Web through 'Settings' > 'User Details' > 'New API Key'
Uncompress the downloaded file and copy as follows:
$cli_conf
└── ssl
    ├── ca_cert.pem
    ├── cert.crt
    └── private
        └── cert.key
EOF
}

installedMessage(){
	printf "\n IMCO CLI is installed.\n Type 'concerto' to access IMCO commands\n\n"
}

showLogo(){
	[ $LOGO_H_SIZE -lt $(tput cols) ] && logoSimple || logoSimple
	echo "Executing IM Cloud Orchestrator CLI install"
}

logoSimple(){
cat <<EOF
                                 :::::::::::                    
                              :::::::::::::::::                 
                            :::::::::::::::::::::               
                           ::::::           .::::::             
                    .::::::::::                :::::            
                 :::::::::::::                  ::::            
               :::::::::  :::                   :::::           
              ::::::                            ::::            
       :::::::::::                              :::::::::       
    :::::::::::::               :: ::::::      :::::::::::::    
  :::::::                      ::::::::::::::.         :::::::  
 :::::                        ::::::::::::::::.           ::::: 
:::::                       ::::::::    ::::::             :::::
::::                        ::::::        ::::::            ::::
::::                         ::::::      .::::::            ::::
::::                         ::::::::::::::::::             ::::
:::::                        :::::::::::::::::             .::::
::::::                .:::::: .::::::::::::::             ::::::
 :::::::      ::::::::::::::::     :::::. : ::          ::::::: 
   ::::::::: :::::::::::::::::::::::   .:::::::::: ::::::::::   
     ::::::: ::::::::::::::::::::::::.:::::::::::: ::::::::     
         :::::::::::::::::::::::::::: ::::    .::::::::         
         :::::::::::        :::::::::  ::::   :::::             
         :::::::::            :::::::: ::::::::::               
         :::::::::            :::::::::.  :::::::               
          ::::::::            :::::::::                         
           :::::::::         ::::::::::                         
           :::::::::::::::::::::::::    Ingram Micro Cloud Orchestrator
           ::::::::::::::::::::::::     https://start.concerto.io                       
            :::::::::::::::::::::::                             
                  ::::::::::::::::                              
                   :::::::                                      
EOF
}






showLogo
parseArgs $@
initializeIMCOCLI
installIMCOCLI
writeDefaultConfig
installAPIKeys
installedMessage
