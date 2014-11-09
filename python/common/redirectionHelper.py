import re
import config as CONFIG

paramre = re.compile('\$\{([0-9]+:[-+!][^}]+)\}')

def apply(target,path):
  if path is None:
    extra = ""
  elif not path.startswith("/"):
    extra = "/" + path
  else:
    extra = path
  values = extra.split("/")
  params = {}
  params["0"] = extra
  for i in range(1,len(values)):
    params[str(i)] = values[i]
    i+=1
  print "extra:"
  print extra
  print "values:"
  print values
  print "params:"
  print params

  plist = paramre.findall(target)
  print plist
 
  if not plist:
    print "no dynamic param"
    return target
 
  for expr in plist:
    if ":" in expr:
      key = expr[:expr.find(":")]
      defvalue = expr[expr.find(":"):][1:]

      if defvalue.startswith("+"):
        altvalue = True
        defvalue = defvalue[1:]
      elif defvalue.startswith("-"):
        altvalue = False
        defvalue = defvalue[1:]
      elif defvalue.startswith("!"):
        onlyifnotset = True
        altvalue = False
        defvalue = defvalue[1:]
      else:
        altvalue = False
    else:
      key = expr

      #
      # Support the following syntaxes
      #
      # ${X:-DEFAULT}  If parameter X is not set, use 'DEFAULT', otherwise substitute parameter X
      # ${X:!DEFAULT}  If parameter X is not set, use 'DEFAULT', otherwise use the empty string ''
      # ${X:+DEFAULT}  If parameter X is set, use 'DEFAULT' where $X was replaced by value of X
      # ${X} is equivalent to ${X:-}
      #
    if params.has_key(key):
      # Positional parameter was set, replace occurences of ${X[:...]} with its value.
      if defvalue is not None and altvalue:
        # First substitute occurences of $X in the default value
        defvalue = defvalue.replace("$" + key, params[key])
        target = target.replace("${" + expr + "}", defvalue);
      else:
        if not onlyifnotset:
          target = target.replace("${" + expr + "}", params[key])
        else:
          target = target.replace("${" + expr + "}", "")
    else:
      if defvalue is not None and not altvalue:
        target = target.replace("${" + expr + "}", defvalue)    
      else:
        target = target.replace("${" + expr + "}", "")    
  return target 
