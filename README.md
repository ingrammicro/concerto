# IMCO CLI / Go Library

[![Go Report Card](http://goreportcard.com/badge/ingrammicro/concerto)](http://goreportcard.com/report/ingrammicro/concerto)
[![Build Status](https://drone.io/github.com/ingrammicro/concerto/status.png)][cli_build] [![GoDoc](https://godoc.org/github.com/ingrammicro/concerto?status.png)](https://godoc.org/github.com/ingrammicro/concerto)
[![codecov.io](https://codecov.io/github/ingrammicro/concerto/coverage.svg?branch=master)](https://codecov.io/github/ingrammicro/concerto?branch=master)
[![Build Status](https://travis-ci.org/ingrammicro/concerto.svg?branch=master)](https://travis-ci.org/ingrammicro/concerto)

Ingram Micro Cloud Orchestrator Command Line Interface (aka IMCO CLI) allows you to interact with IMCO features, and build your own scripts calling IMCO's API.

If you are already using IMCO CLI, and only want to obtain the latest version, download IMCO CLI for:

- [Linux][cli_linux]
- [OSX][cli_darwin]
- [Windows][cli_windows]

> NOTE: IMCO CLI is named as `concerto` in terms of the binary and executable.

If you want to build the CLI using the source code, please, take into account that the master branch is the adequate one to be used for latest stable and published version of IMCO CLI.

## Table of Contents

- [Setup](#setup)
  - [Pre-requisites](#pre-requisites)
  - [Setup script](#setup-script)
  - [Manual Setup](#manual-setup)
  - [Linux and OSX](#linux-and-osx)
    - [Configuration](#configuration)
    - [Binaries](#binaries)
  - [Environment variables](#environment-variables)
  - [Troubleshooting](#troubleshooting)
- [Usage](#usage)
  - [Wizard](#wizard)
    - [Wizard Use Case](#wizard-use-case)
  - [Blueprint](#blueprint)
    - [Blueprint Use Case](#blueprint-use-case)
      - [Template OS](#template-os)
      - [Service List](#service-list)
      - [Instantiate a server](#instantiate-a-server)
  - [Firewall Management](#firewall-management)
    - [Firewall Update Case](#firewall-update-case)
  - [Blueprint Update](#blueprint-update)
    - [Blueprint Update Case](#blueprint-update-case)
- [Contribute](#contribute)

## Setup

### Pre-requisites

Before setting up the CLI, you will need a IMCO account, and an API key associated with your account.

> NOTE: The API Endpoint server value depends on the targeted IMCO platform domain: https://clients.IMCO_DOMAIN:886

Once your account have been provisioned, if you are a linux or OS X, we recommend you to execute the automated setup script. Otherwise, follow the [manual process](#manual-setup)

### Setup script

Open a terminal window and execute

```bash
$ curl -sSL goo.gl/ujPLzA | sh
```

The script will drive you through:

- IMCO CLI Binary download
- Configuration
- API keys creation

The setup script can take these arguments:

- `fb` forces the binary to be overwriten
- `fc` forces the configuration file to be overwriten
- `fk` forces the API keys to be overwriten
- `f` forces binary, configuration and API keys to be overwriten
- `v` verbose mode

Example

```bash
$ curl -sSL goo.gl/ujPLzA | sh -s f
```

## Manual Setup

Use IMCO's Web UI to navigate the menus to `Settings` > `User Details` and scroll down until you find the `New API Key` button.

<img src="./docs/images/newAPIkey.png" alt="API Key" width="500px" >

Pressing `New API Key` will download a compressed file that contains the necessary files to authenticate with IMCO API and manage your infrastructure. `Keep it safe`.

Extract the contents with your zip compressor of choice and continue using the setup guide for your O.S.

## Linux and OSX

### Configuration

IMCO CLI configuration will usually be located in your personal folder under `.concerto`. If you are using root, CLI will look for contiguration files under `/etc/imco`.
We will assume that you are not root, so create the folder and drop the certificates to this location:

```bash
$ mkdir -p ~/.concerto/ssl/
$ unzip -x api-key.zip -d ~/.concerto/ssl
```

IMCO CLI expects a configuration file to be present containing:

- API Endpoint
- Log file
- Log level
- Certificate location

This command will generate the file `~/.concerto/client.xml` with suitable contents for most users:

```bash
$ cat <<EOF > ~/.concerto/client.xml
<concerto version="1.0" server="https://clients.IMCO_DOMAIN:886/" log_file="/var/log/concerto-client.log" log_level="info">
 <ssl cert="$HOME/.concerto/ssl/cert.crt" key="$HOME/.concerto/ssl/private/cert.key" server_ca="$HOME/.concerto/ssl/ca_cert.pem" />
</concerto>
EOF
```

We should have in your `.concerto` folder this structure:

```bash
$HOME/.concerto
├── client.xml
└── ssl
    ├── ca_cert.pem
    ├── cert.crt
    └── private
        └── cert.key
```

### Binaries

Download linux binaries for [Linux][cli_linux] or for [OSX][cli_darwin] and place it in your path.

Linux:

```bash
$ sudo curl -o /usr/local/bin/concerto https://github.com/ingrammicro/concerto/raw/master/binaries/concerto.x64.linux
$ sudo chmod +x /usr/local/bin/concerto
```

OSX:

```bash
$ sudo curl -o /usr/local/bin/concerto https://github.com/ingrammicro/concerto/raw/master/binaries/concerto.x64.darwin
$ sudo chmod +x /usr/local/bin/concerto
```

To test the binary execute `concerto` without parameters

```bash
$ concerto
NAME:
   concerto - Manages communication between Host and Concerto Platform

USAGE:
   concerto [global options] command [command options] [arguments...]

VERSION:
   0.6.0

AUTHOR:
   Concerto Contributors <https://github.com/ingrammicro/concerto>

COMMANDS:
     setup, se              Configures and setups concerto cli enviroment
     nodes, no              Manages Docker Nodes
     cluster, clu           Manages a Kubernetes Cluster
     reports, rep           Provides historical uptime of servers
     events, ev             Events allow the user to track their actions and the state of their servers
     blueprint, bl          Manages blueprint commands for scripts, services and templates
     cloud, clo             Manages cloud related commands for workspaces, servers, generic images, ssh profiles, cloud providers, server plans and Saas providers
     licensee_reports, lic  Provides information about licensee reports
     network, net           Manages network related commands for firewall profiles
     settings, set          Provides settings for cloud and Saas accounts as well as reports
     wizard, wiz            Manages wizard related commands for apps, locations, cloud providers, server plans
...
```

To test that certificates are valid, and that we can communicate with IMCO server, obtain the list of workspaces at your IMCO account using this command

```bash
$ concerto cloud workspaces list
ID                         NAME                  DEFAULT        SSH_PROFILE_ID             FIREWALL_PROFILE_ID
5aabb7521de0240abb00000e   default               true           5aabb7521de0240abb00000d   5aabb7521de0240abb00000c
5aeae3fafbd05409ef000005   Wordpress_workspace   false          5aabb7521de0240abb00000d   5aeae3fafbd05409ef000003
```

## Environment variables

When using IMCO CLI you can override configuration parameters using the following environment variables:

Env. Variable | Descripcion
------------------------|---------------------
`CONCERTO_ENDPOINT` | IMCO API endpoint
`CONCERTO_CA_CERT` | CA certificate used with the API endpoint.
`CONCERTO_CLIENT_CERT` | Client certificate used with the API endpoint.
`CONCERTO_CLIENT_KEY` | Client key used with the API endpoint.
`CONCERTO_CONFIG` | Config file to be read by Concerto CLI.
`CONCERTO_URL` | IMCO web site URL.

## Troubleshooting

If you got an error executing IMCO CLI:

- execute `which concerto` to make sure that the binary is installed.
- execute `ls -l /path/to/concerto` with the output from the previous command, and check that you have execute permissions.
- execute `$PATH` and search for the path where `concerto` is installed. If `concerto` isn't in the path, move it to a `$PATH` location.
- check that your internet connection can reach `clients.IMCO_DOMAIN`
- make sure that your firewall lets you access to <https://clients.IMCO_DOMAIN:886>
- check that `client.xml` is pointing to the correct certificates location
- if `concerto` executes but only shows server commands, you are probably trying to use `concerto` from a commissioned server, and the configuration is being read from `/etc/imco`. If that's the case, you should leave `concerto` configuration untouched so that server commands are available for our remote management.

## Usage

We include the most common use cases here. If you feel there is a missing a use case here, open an issue or contact us at <enquiries@concerto.io>.

## Wizard

The Wizard command for IMCO CLI is the command line version of our `Quick add server` in the IMCO's Web UI.

<img src="./docs/images/webwizard.png" alt="Web Wizard" width="500px" >

Wizard is the quickest way to install a well known stack in a cloud server. You can get an idea of what the wizard does using the command `concerto wizard` without further subcommands:

```bash
$ concerto wizard
NAME:
    - Manages wizard related commands for apps, locations, cloud providers, server plans

USAGE:
    command [command options] [arguments...]

COMMANDS:
     apps             Provides information about apps
     cloud_providers  Provides information about cloud providers
     locations        Provides information about locations
     server_plans     Provides information about server plans
...
```

IMCO CLI Wizard lets you select the application layer, the location, the cloud provider account for that location, and finally the hostname. IMCO CLI Wizard takes care of the details.

If you haven't configured your cloud provider accounts yet, you can do it from the IMCO's Web UI, or using `concerto settings cloud_accounts` commands

### Wizard Use Case

Let's type `concerto wizard apps list` to check what servers can I instantiate using IMCO CLI wizard.

```bash
$ concerto wizard apps list
ID                         NAME              FLAVOUR_REQUIREMENTS   GENERIC_IMAGE_ID
5aabb75a1de0240abb000185   Ubuntu 14.04      {}                     5aabb7551de0240abb000064
5aabb75a1de0240abb000186   Ubuntu 16.04      {}                     5aabb7551de0240abb000065
5aabb75a1de0240abb000187   Windows 2012 R2   {"memory":4096}        5aabb7551de0240abb000066
5aabb75a1de0240abb000188   Joomla            {"memory":1024}        5aabb7551de0240abb000064
5aabb75a1de0240abb000189   Magento           {"memory":1024}        5aabb7551de0240abb000064
5aabb75b1de0240abb00018a   MongoDB           {}                     5aabb7551de0240abb000064
5aabb75b1de0240abb00018b   Wordpress         {"memory":1024}        5aabb7551de0240abb000064
5aabb75b1de0240abb00018c   Docker            {"memory":2048}        5aabb7551de0240abb000064
```

You can choose whatever application/stack is fine for your purpose, we choose `Wordpress`. Take note of the application identifier, `5aabb75b1de0240abb00018b` for `Wordpress`.

We will also need the location where we want our server to be instantiated. Execute `concerto wizard locations list` to get the possible locations and its identifier.

```bash
$ concerto wizard locations list
ID                         NAME
5aabb7551de0240abb000060   North America
5aabb7551de0240abb000061   Europe
5aabb7551de0240abb000062   Asia Pacific
5aabb7551de0240abb000063   South America
```

Take note of your preferred location. We will use `5aabb7551de0240abb000060` for `North America`.

When using IMCO's Web UI, the wizard takes care of filtering appropriate cloud accounts for that provider and location. However, using the CLI is the user's responsibility to chose a provider cloud account for that application/stack and location; and a server plan capable of instantiating the stack in that location.
To show all possible cloud accounts execute this command:

```bash
$ concerto wizard cloud_providers list --app_id 5aabb75b1de0240abb00018b --location_id 5aabb7551de0240abb000060
ID                         NAME                  REQUIRED_CREDENTIALS
5aabb7511de0240abb000001   AWS                   [access_key_id secret_access_key]
5aabb7511de0240abb000002   Mock                  [nothing]
5aabb7511de0240abb000004   Microsoft Azure ARM   [tenant_id client_id secret subscription_id]
5aabb7511de0240abb000005   Microsoft Azure       [tenant_id subscription_id]
```

Take also into account that you should have configured your credentials before, using the Web UI or `concerto settings cloud_accounts create`. We will choose `Microsoft Azure`, whose ID is `5aabb7511de0240abb000005`.

It's necessary to retrive the adequeate Cloud Account ID for `Microsoft Azure` Cloud Provider, in our case `5aabb7531de0240abb000024`:

```bash
$ concerto settings cloud_accounts list
ID                         CLOUD_PROVIDER_ID
5aabb7521de0240abb00001b   5aabb7511de0240abb000001
5aabb7521de0240abb00001c   5aabb7511de0240abb000002
5aabb7531de0240abb00001d   5aabb7511de0240abb000002
5aabb7531de0240abb00001e   5aabb7511de0240abb000002
5aabb7531de0240abb000020   5aabb7511de0240abb000003
5aabb7531de0240abb000022   5aabb7511de0240abb000004
5aabb7531de0240abb000024   5aabb7511de0240abb000005
5aba0656425b5d0c64000001   5aba04be425b5d0c16000000
5aba066c425b5d0c64000002   5aba04be425b5d0c16000000
```

Now that we have all the data that we need, commission the server:

```bash
$ concerto wizard apps deploy -id 5aabb75b1de0240abb00018b --location_id 5aabb7551de0240abb000060 --cloud_account_id 5aabb7531de0240abb000024 --hostname wpnode1
ID:                     5af98e0aea7a1a000c000000
NAME:                   wpnode1
FLAVOUR_REQUIREMENTS:
GENERIC_IMAGE_ID:
```

We have a new server template and a workspace with a commissioned server in IMCO.
<img src="./docs/images/commissioned-server.png" alt="Server Commissioned" width="500px" >

From the command line, get the new workspace, and then our commissioned server ID.

```bash
$ concerto cloud workspaces list
ID                         NAME                  DEFAULT        SSH_PROFILE_ID             FIREWALL_PROFILE_ID
5aabb7521de0240abb00000e   default               true           5aabb7521de0240abb00000d   5aabb7521de0240abb00000c
5aeae3fafbd05409ef000005   Wordpress_workspace   false          5aabb7521de0240abb00000d   5aeae3fafbd05409ef000003
```

```bash
$ concerto cloud workspaces list_workspace_servers --workspace_id 5aeae3fafbd05409ef000005
ID                         NAME           FQDN                                            STATE          PUBLIC_IP        WORKSPACE_ID               TEMPLATE_ID                SERVER_PLAN_ID             SSH_PROFILE_ID
5af98e0aea7a1a000c000000   wpnode1        s4bc576baee4d07b.centralus.cloudapp.azure.com   inactive       104.43.129.103   5aeae3fafbd05409ef000005   5aeae3fafbd05409ef000000   5aac0c05348f190b3e0011c2   5aabb7521de0240abb00000d
```

Our server's ID is `5af98e0aea7a1a000c000000`. We can now use `concerto cloud servers` subcommands to manage the server. Lets bring wordpress up:

```bash
$ concerto cloud servers boot --id 5af98e0aea7a1a000c000000
ID:                 5af98e0aea7a1a000c000000
NAME:               wpnode1
FQDN:               s4bc576baee4d07b.centralus.cloudapp.azure.com
STATE:              booting
PUBLIC_IP:          104.43.129.103
WORKSPACE_ID:       5aeae3fafbd05409ef000005
TEMPLATE_ID:        5aeae3fafbd05409ef000000
SERVER_PLAN_ID:     5aac0c05348f190b3e0011c2
CLOUD_ACCOUNT_ID:   5aabb7531de0240abb000024
SSH_PROFILE_ID:     5aabb7521de0240abb00000d
```

<img src="./docs/images/server-bootstraping.png" alt="Server Bootstraping" width="500px" >

<img src="./docs/images/server-operational.png" alt="Server Operational" width="500px" >

After a brief amount of time you will have your new `Wordpress` server up and running, ready to be configured.

<img src="./docs/images/wordpress.png" alt="Wordpress" width="500px" >

## Blueprint

IMCO blueprints are the compendium of:

- services, they map to IMCO's Web UI cookbooks. Use `concerto blueprint services list` to show all cookbooks available at your account.
- scripts, they provide a way to execute custom scripts after bootstraping, before a clean shutdown, or on demand.
- templates, an ordered combination of services and scripts.

### Blueprint Use Case

A template must be created with an OS target, a service list, and a list of custom attributes for those services.

#### Template OS

Blueprints are associated with an Operative System, and each cloud provider has a different way of identifying the OS that a machine is running.

IMCO takes care of the gap, and lets you select a cloud provider independent OS, and find out later which image is appropriate for the chosen cloud provider account and location. Hence blueprints are bound to OS, but cloud provider and location independent.

For our case we will be using Ubuntu 14.04. Let's find its IMCO ID

```bash
$ concerto cloud generic_images list
ID                         NAME
5aabb7551de0240abb000064   Ubuntu 14.04 Trusty Tahr x86_64
5aabb7551de0240abb000065   Ubuntu 16.04 Xenial Xerus x86_64
5aabb7551de0240abb000066   Windows 2012 R2 x86_64
5aabb7551de0240abb000067   Windows 2016 x86_64
5aabb7551de0240abb000068   Red Hat Enterprise Linux 7.3 x86_64
5aabb7551de0240abb000069   CentOS 7.4 x86_64
5aabb7551de0240abb00006a   Debian 9 x86_64
```

Take note of Ubuntu 14.04 ID, `5aabb7551de0240abb000064`.

#### Service List

We want to use IMCO's curated Joomla cookbook. Use `concerto blueprint services` to find the cookbooks to add.

```bash
$ concerto blueprint services list | awk 'NR==1 || /joomla/'
ID                         NAME                  DESCRIPTION                                    PUBLIC         LICENSE               RECIPES
5aabb871e4997809f700000e   joomla                Installs/Configures joomla environment         false          All rights reserved   [joomla@0.10.0 joomla::appserver@0.10.0 joomla::database@0.10.0]
```

Joomla curated cookbooks creates a local mysql database. We only have to tell our cookbook that we should override the `joomla.db.hostname` to `127.0.0.1`. Execute the following command to create the Joomla template.

```bash
$ concerto blueprint templates create --name joomla-tmplt --generic_image_id 5aabb7551de0240abb000064 --service_list '["joomla"]' --configuration_attributes '{"joomla":{"db":{"hostname":"127.0.0.1"}}}'
ID:                         5af9aab5ea7a1a000d00000e
NAME:                       joomla-tmplt
GENERIC IMAGE ID:           5aabb7551de0240abb000064
SERVICE LIST:               [joomla]
CONFIGURATION ATTRIBUTES:   {"joomla":{"db":{"hostname":"127.0.0.1"}}}
```

#### Instantiate a server

Now that we have our server blueprint defined, let's start one. Servers in IMCO need to know the workspace that define their runtime infrastructure environment, the server plan for the cloud provider, and the template used to build the instance.

As we did in the Wizard use case, we can find the missing data using these commands:

##### Find the workspace

```bash
$ concerto cloud workspaces list
ID                         NAME                  DEFAULT        SSH_PROFILE_ID             FIREWALL_PROFILE_ID
5aabb7521de0240abb00000e   default               true           5aabb7521de0240abb00000d   5aabb7521de0240abb00000c
5aeae3fafbd05409ef000005   Wordpress_workspace   false          5aabb7521de0240abb00000d   5aeae3fafbd05409ef000003
```

##### Find cloud provider server plan

```bash
$ concerto cloud cloud_providers list
ID                         NAME                  REQUIRED_CREDENTIALS
5aabb7511de0240abb000001   AWS                   [access_key_id secret_access_key]
5aabb7511de0240abb000002   Mock                  [nothing]
5aabb7511de0240abb000003   DigitalOcean          [personal_token]
5aabb7511de0240abb000004   Microsoft Azure ARM   [tenant_id client_id secret subscription_id]
5aabb7511de0240abb000005   Microsoft Azure       [tenant_id subscription_id]
5aba04be425b5d0c16000000   VCloud                [vdc_id routed]
```

We want to use `Microsoft Azure` with ID `5aabb7511de0240abb000005` and filtering by server_plan `Basic_A0`

```bash
$ concerto cloud server_plans list --cloud_provider_id 5aabb7511de0240abb000005 | awk 'NR==1 || /Basic_A0/'
ID                         NAME                  MEMORY         CPUS           STORAGE        LOCATION_ID                CLOUD_PROVIDER_ID
5aac0bff348f190b3e001030   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c02348f190b3e0010db   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c04348f190b3e001186   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c06348f190b3e001231   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c09348f190b3e0012dc   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c0b348f190b3e001387   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c0e348f190b3e001432   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c10348f190b3e0014dd   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c13348f190b3e001588   Basic_A0              768            1              20             5aabb7551de0240abb000061   5aabb7511de0240abb000005
5aac0c15348f190b3e001633   Basic_A0              768            1              20             5aabb7551de0240abb000061   5aabb7511de0240abb000005
5aac0c18348f190b3e0016de   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c1b348f190b3e001789   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c1d348f190b3e001834   Basic_A0              768            1              20             5aabb7551de0240abb000063   5aabb7511de0240abb000005
5aac0c20348f190b3e0018df   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c23348f190b3e00198a   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c26348f190b3e001a35   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c29348f190b3e001ae0   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c2c348f190b3e001b8b   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c2f348f190b3e001c36   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c32348f190b3e001ce1   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c35348f190b3e001d8c   Basic_A0              768            1              20             5aabb7551de0240abb000061   5aabb7511de0240abb000005
5aac0c38348f190b3e001e37   Basic_A0              768            1              20             5aabb7551de0240abb000061   5aabb7511de0240abb000005
5aac0c3c348f190b3e001ee2   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c3f348f190b3e001f8d   Basic_A0              768            1              20             5aabb7551de0240abb000060   5aabb7511de0240abb000005
5aac0c42348f190b3e002038   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
5aac0c45348f190b3e0020e3   Basic_A0              768            1              20             5aabb7551de0240abb000062   5aabb7511de0240abb000005
```

##### Find Template ID

We already know our template ID, but in case you want to make sure

```bash
$ concerto blueprint templates list
ID                         NAME                  GENERIC IMAGE ID
5ac734dbfc0d680143000014   ubuntu 14.04 x86_64   5aabb7551de0240abb000064
5ac734dbfc0d68014300001c   ubuntu16 x86_64       5aabb7551de0240abb000065
5ac734dcfc0d68014300003c   Windows 2012 x86_64   5aabb7551de0240abb000066
5ac734dcfc0d680143000040   CentOs 7.4 x86_64     5aabb7551de0240abb000069
5ad60467fd851c0b4b000067   Debian 9 x86_64       5aabb7551de0240abb00006a
5ae0a1a49fca7d02bb00001b   Windows 2016 x86_64   5aabb7551de0240abb000067
5aeae3fafbd05409ef000000   Wordpress_template    5aabb7551de0240abb000064
5af9aab5ea7a1a000d00000e   joomla-tmplt          5aabb7551de0240abb000064
```

##### Find Location ID

We already know our location ID, but in case you want to make sure

```bash
$ concerto wizard locations list
ID                         NAME
5aabb7551de0240abb000060   North America
5aabb7551de0240abb000061   Europe
5aabb7551de0240abb000062   Asia Pacific
5aabb7551de0240abb000063   South America
```

##### Find Cloud Account ID

It's necessary to retrive the adequeate Cloud Account ID for `Microsoft Azure` Cloud Provider, in our case `5aabb7511de0240abb000005`:

```bash
$ concerto settings cloud_accounts list
ID                         CLOUD_PROVIDER_ID
5aabb7521de0240abb00001b   5aabb7511de0240abb000001
5aabb7521de0240abb00001c   5aabb7511de0240abb000002
5aabb7531de0240abb00001d   5aabb7511de0240abb000002
5aabb7531de0240abb00001e   5aabb7511de0240abb000002
5aabb7531de0240abb000020   5aabb7511de0240abb000003
5aabb7531de0240abb000022   5aabb7511de0240abb000004
5aabb7531de0240abb000024   5aabb7511de0240abb000005
5aba0656425b5d0c64000001   5aba04be425b5d0c16000000
5aba066c425b5d0c64000002   5aba04be425b5d0c16000000
```

##### Create our Joomla Server

```bash
$ concerto cloud servers create --name joomla-node1 --workspace_id 5aabb7521de0240abb00000e --template_id 5af9aab5ea7a1a000d00000e --server_plan_id 5aac0c04348f190b3e001186 --cloud_account_id 5aabb7531de0240abb000024
ID:                 5af9acedea7a1a000d000012
NAME:               joomla-node1
FQDN:
STATE:              commissioning
PUBLIC_IP:
WORKSPACE_ID:       5aabb7521de0240abb00000e
TEMPLATE_ID:        5af9aab5ea7a1a000d00000e
SERVER_PLAN_ID:     5aac0c04348f190b3e001186
CLOUD_ACCOUNT_ID:   5aabb7531de0240abb000024
SSH_PROFILE_ID:     5aabb7521de0240abb00000d
```

And finally boot it

```bash
$ concerto cloud servers boot --id 5af9acedea7a1a000d000012
ID:                 5af9acedea7a1a000d000012
NAME:               joomla-node1
FQDN:
STATE:              booting
PUBLIC_IP:
WORKSPACE_ID:       5aabb7521de0240abb00000e
TEMPLATE_ID:        5af9aab5ea7a1a000d00000e
SERVER_PLAN_ID:     5aac0c04348f190b3e001186
CLOUD_ACCOUNT_ID:   5aabb7531de0240abb000024
SSH_PROFILE_ID:     5aabb7521de0240abb00000d
```

You can request for status and see how server is transitioning along service statuses (booting, bootstrapping, operational). Then, after a brief amount of time the final status is reached:

```bash
$ concerto cloud servers show --id 5af9acedea7a1a000d000012
ID:                 5af9acedea7a1a000d000012
NAME:               joomla-node1
FQDN:               sf452e764a403c4f.centralus.cloudapp.azure.com
STATE:              operational
PUBLIC_IP:          104.43.209.103
WORKSPACE_ID:       5aabb7521de0240abb00000e
TEMPLATE_ID:        5af9aab5ea7a1a000d00000e
SERVER_PLAN_ID:     5aac0c04348f190b3e001186
CLOUD_ACCOUNT_ID:   5aabb7531de0240abb000024
SSH_PROFILE_ID:     5aabb7521de0240abb00000d
```

## Firewall Management

IMCO CLI's `network` command lets you manage a network settings at the workspace scope.

As we have did before, execute this command with no futher commands to get usage information:

```bash
$ concerto network
NAME:
    - Manages network related commands for firewall profiles

USAGE:
    command [command options] [arguments...]

COMMANDS:
     firewall_profiles  Provides information about firewall profiles
```

As you can see, you can manage firewall from IMCO CLI.

### Firewall Update Case

Workspaces in IMCO are always associated with a firewall profile. By default ports 443 and 80 are open to fit most web environments, but if you are not using those ports but some others. We would need to close HTTP and HTTPS ports and open LDAP and LDAPS instead.

The first thing we will need is our workspace's related firewall identifier.

```bash
$ concerto cloud workspaces list
ID                         NAME                  DEFAULT        SSH_PROFILE_ID             FIREWALL_PROFILE_ID
5aabb7521de0240abb00000e   default               true           5aabb7521de0240abb00000d   5aabb7521de0240abb00000c
5aeae3fafbd05409ef000005   Wordpress_workspace   false          5aabb7521de0240abb00000d   5aeae3fafbd05409ef000003
5af9b2a042d90d09f000000b   My New Workspace      false          5aabb7521de0240abb00000d   5af9b28c42d90d09f0000008
```

We have our LDAP servers running under `My New Workspace`. If you are unsure about in which workspace are your servers running, list the servers in the workspace

```bash
concerto cloud workspaces list_workspace_servers --workspace_id 5af9b2a042d90d09f000000b
ID                         NAME           FQDN                                            STATE          PUBLIC_IP        WORKSPACE_ID               TEMPLATE_ID                SERVER_PLAN_ID             SSH_PROFILE_ID
5af9b31b42d90d09f000000d   openldap-1                                                     inactive                        5af9b2a042d90d09f000000b   5ad60467fd851c0b4b000067   5aabb76be499780a00000399   5aabb7521de0240abb00000d
5af9b33242d90d09f0000010   openldap-2     s485036e9e25cd6e.centralus.cloudapp.azure.com   operational    104.43.132.132   5af9b2a042d90d09f000000b   5ad60467fd851c0b4b000067   5aabb76ae499780a000002ee   5aabb7521de0240abb00000d
```

Now that we have the firewall profile ID, list it's contents

```bash
$ concerto network firewall_profiles show --id 5af9b28c42d90d09f0000008
ID:            5af5894bfb170309f0000022
NAME:          My New Firewall Profile
DESCRIPTION:
DEFAULT:       false
RULES:         [{Protocol:tcp MinPort:22 MaxPort:22 CidrIp:any} {Protocol:tcp MinPort:5985 MaxPort:5985 CidrIp:any} {Protocol:tcp MinPort:3389 MaxPort:3389 CidrIp:any} {Protocol:tcp MinPort:10050 MaxPort:10050 CidrIp:any} {Protocol:tcp MinPort:443 MaxPort:443 CidrIp:any} {Protocol:tcp MinPort:80 MaxPort:80 CidrIp:any}]
```

The first four values are ports that IMCO may use to keep the desired state of the machine, and that will always be accessed using certificates.

When updating, we tell IMCO a new set of rules. Execute the following command to open 389 and 686 to anyone.

```bash
$ concerto network firewall_profiles update --id 5af9b28c42d90d09f0000008 --rules '[{"ip_protocol":"tcp", "min_port":389, "max_port":389, "source":"0.0.0.0/0"}, {"ip_protocol":"tcp", "min_port":636, "max_port":636, "source":"0.0.0.0/0"}]'
ID:            5af9b28c42d90d09f0000008
NAME:          My New Firewall Profile
DESCRIPTION:
DEFAULT:       false
RULES:         [{Protocol:tcp MinPort:22 MaxPort:22 CidrIp:any} {Protocol:tcp MinPort:5985 MaxPort:5985 CidrIp:any} {Protocol:tcp MinPort:3389 MaxPort:3389 CidrIp:any} {Protocol:tcp MinPort:10050 MaxPort:10050 CidrIp:any} {Protocol:tcp MinPort:389 MaxPort:389 CidrIp:any} {Protocol:tcp MinPort:636 MaxPort:636 CidrIp:any}]
```

Firewall update returns the complete set of rules. As you can see, now LDAP and LDAPS ports are open.

## Blueprint Update

We have already used [blueprints](#blueprint) before. So you might already know that we can delete and update blueprints.

### Blueprint Update Case

Let's pretend there is an existing Joomla blueprint, and that we want to update the previous password to a safer one.

This is the Joomla blueprint that we created in a previous use case.

```bash
$ concerto blueprint templates show --id 5af9aab5ea7a1a000d00000e
ID:                         5af9aab5ea7a1a000d00000e
NAME:                       joomla-tmplt
GENERIC IMAGE ID:           5aabb7551de0240abb000064
SERVICE LIST:               [joomla]
CONFIGURATION ATTRIBUTES:   {"joomla":{"db":{"hostname":"127.0.0.1"}}}
```

Beware of adding previous services or configuration attributes. Update will replace existing items with the ones provided. If we don't want to lose the `joomla.db.hostname` attribute, add it to our configuretion attributes parameter:

```bash
$ concerto blueprint templates update --id 5af9aab5ea7a1a000d00000e --configuration_attributes '{"joomla":{"db":{"hostname":"127.0.0.1", "password":"$afeP4sSw0rd"}}}'
ID:                         5af9aab5ea7a1a000d00000e
NAME:                       joomla-tmplt
GENERIC IMAGE ID:           5aabb7551de0240abb000064
SERVICE LIST:               [joomla]
CONFIGURATION ATTRIBUTES:   {"joomla":{"db":{"hostname":"127.0.0.1","password":"$afeP4sSw0rd"}}}
```

As you can see, non specified parameters, like name and service list, remain unchanged. Let's now change the service list, adding a two cookbooks.

```bash
$ concerto blueprint templates update --id 5af9aab5ea7a1a000d00000e  --service_list '["joomla","python@1.4.6","polipo"]'
ID:                         5af9aab5ea7a1a000d00000e
NAME:                       joomla-tmplt
GENERIC IMAGE ID:           5aabb7551de0240abb000064
SERVICE LIST:               [joomla python@1.4.6 polipo]
CONFIGURATION ATTRIBUTES:   {"joomla":{"db":{"hostname":"127.0.0.1","password":"$afeP4sSw0rd"}}}
```

Of course, we can change service list and configuration attributes in one command.

```bash
$ concerto blueprint templates update --id 5af9aab5ea7a1a000d00000e --configuration_attributes '{"joomla":{"db":{"hostname":"127.0.0.1", "password":"$afeP4sSw0rd"}}}' --service_list '["joomla","python@1.4.6","polipo"]'
ID:                         5af9aab5ea7a1a000d00000e
NAME:                       joomla-tmplt
GENERIC IMAGE ID:           5aabb7551de0240abb000064
SERVICE LIST:               [joomla python@1.4.6 polipo]
CONFIGURATION ATTRIBUTES:   {"joomla":{"db":{"hostname":"127.0.0.1","password":"$afeP4sSw0rd"}}}
```

## Contribute

To contribute

- Find and open issue, or report a new one. Include proper information about the environment, at least: operating system, CLI version, steps to reproduce the issue and related issues. Avoid writing multi-issue reports, and make sure that the issue is unique.
- Fork the repository to your account
- Commit scoped chunks, adding concise and clear comments
- Remember to add tests to your contributed code
- Push changes to the forked repository
- Submit the PR to IMCO CLI
- Let the maintainers give you the LGTM.

Please, use gofmt, golint, go vet, and follow [go style](https://github.com/golang/go/wiki/CodeReviewComments) advices

[cli_build]: https://drone.io/github.com/ingrammicro/concerto/latest
[cli_linux]: https://github.com/ingrammicro/concerto/raw/master/binaries/concerto.x64.linux
[cli_darwin]: https://github.com/ingrammicro/concerto/raw/master/binaries/concerto.x64.darwin
[cli_windows]: https://github.com/ingrammicro/concerto/raw/master/binaries/concerto.x64.windows.exe
