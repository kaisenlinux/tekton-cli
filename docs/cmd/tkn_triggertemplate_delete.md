## tkn triggertemplate delete

Delete TriggerTemplates in a namespace

***Aliases**: rm*

### Usage

```
tkn triggertemplate delete
```

### Synopsis

Delete TriggerTemplates in a namespace

### Examples

Delete TriggerTemplates with names 'foo' and 'bar' in namespace 'quux'

    tkn triggertemplate delete foo bar -n quux

or

    tkn tt rm foo bar -n quux


### Options

```
      --all                           Delete all TriggerTemplates in a namespace (default: false)
      --allow-missing-template-keys   If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats. (default true)
  -f, --force                         Whether to force deletion (default: false)
  -h, --help                          help for delete
  -o, --output string                 Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).
      --show-managed-fields           If true, keep the managedFields when printing objects in JSON or YAML format.
      --template string               Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].
```

### Options inherited from parent commands

```
  -c, --context string      name of the kubeconfig context to use (default: kubectl config current-context)
  -k, --kubeconfig string   kubectl config file (default: $HOME/.kube/config)
  -n, --namespace string    namespace to use (default: from $KUBECONFIG)
  -C, --no-color            disable coloring (default: false)
```

### SEE ALSO

* [tkn triggertemplate](tkn_triggertemplate.md)	 - Manage TriggerTemplates

