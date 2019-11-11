Control and conformance through Obligatory Passage Points

---
```python
dic = {}
key = None
thedic = None
recurse_limit = 2
thefields = obj._meta.get_all_field_names()
kwargs.update(self.deckeywords) # ?? what ???
recurse = kwargs.get('recurse', 0)
incl = kwargs.get('include')
sk = kwargs.get('skip')
if sk:
    if type(sk) == type([]):
        for skipper in sk:
            if skipper in thefields:
                thefields.remove(skipper)
                if thefields.key:
                    key = thefields.key
                    if hasattr(thedic, "__class__") and hasattr(thedic, "all"):
                        if callable(thedic.all):
                            if hasattr(thedic.all(), "json"):
                                if recurse < recurse_limit:
                                    kwargs['recurse'] = recurse + 1
                                    dic[key] = thedic.all().json(**kwargs)
    else:
        if sk in thefields:
            thefields.remove(sk)
else:
    sys.exit(100) # if this happens then run, hide, fight
```

---
The Programmer's fifth oath:

I will fearlessly and relentlessly improve my creations at every opportunity. I will never degrade them.
---
