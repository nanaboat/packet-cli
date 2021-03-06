## packet volume create

Creates a volume

### Synopsis

Example:

  packet volume create --size [size_in_GB] --plan [plan_UUID] --project-id [project_UUID] --facility [facility_code]
  
  

```
packet volume create [flags]
```

### Options

```
  -b, --billing-cycle string   Billing cycle (default "hourly")
  -d, --description string     Description of the volume
  -f, --facility string        Code of the facility where the volume will be created
  -h, --help                   help for create
  -l, --locked                 Set the volume to be locked
  -P, --plan string            Name of the plan
  -p, --project-id string      UUID of the project
  -s, --size int               Size in GB]
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet volume](packet_volume.md)	 - Volume operations

