### Control and conformance through Obligatory Passage Points

_(cc-by-sa 4.0)_

---
#### Fixing some small bug

@ul
* Hard when rotten code
* Rot is quite common!
@ulend

---
Though common, _wildly irresponsible_
---
Simply refactor and add tests

@ul
* Except if API is fundamentally flawed
* Radical refactor
@ulend
---
The Programmer's fifth oath:

`I will fearlessly and relentlessly improve my creations at every opportunity. I will never degrade them.`
---
#### Lack of control:

```python
def here_is_the_database_connection_please_do_whatever_you_want(self):
    return self.db_con
```
+++
#### Lack of conformance:

(At least beyond social contract)

```python
def execute_arbitrary_sql(sql_string):
    …  # please no heavy joins
```
---
#### A shift in control (from user to author)


@ul
* Except if API is fundamentally flawed
* Radical refactor
@ulend
---

## Takk
---
+++
EOF
