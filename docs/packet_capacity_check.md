## packet capacity check

Validates if a deploy can be fulfilled.

### Synopsis

Example:

packet capacity check -f [facility] -p [plan] -q [quantity]

	

```
packet capacity check [flags]
```

### Options

```
  -f, --facility string   Code of the facility
  -h, --help              help for check
  -p, --plan string       Name of the plan
  -q, --quantity int      Number of devices wanted
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

* [packet capacity](packet_capacity.md)	 - Capacities operations

