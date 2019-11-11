### Control and conformance through Obligatory Passage Points

_(cc-by-sa 4.0)_

---
#### Fixing some small bug

@ul
* Hard when rotten code
* Rot is quite common!
@ulend

+++
Though common, _wildly irresponsible_
+++
Simply refactor and add tests

@ul
* Except if API is fundamentally flawed
* Radical refactor
@ulend
+++
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

```python
def execute_arbitrary_sql(self, sql_string):
    …  # please no heavy joins... plz
```
+++
#### A shift in control (from user to author)

@ul
* Higher abstraction level
* Conformance to norms simply by use of API
@ulend
+++
#### The _Variation Behind Interface_-principle
@ul
* Protect code integrity
* Things change fast in this biz
@ulend
+++
#### Disruptive shifts of power

Justified by principle alone?
@ul
* Will cause conflict
* Might need more juice in your argument
@ulend
---
#### An argument for radical refactoring

@ul
* Studied Python's source code
* Narrative: openness, collaboration, participation, access
* Practice: closedness, regulation
* Control and conformance through code
* … of both technical and social aspects
* Emergent
@ulend
---
#### Actor–network theory
1\. _Everything_ is part of the network.

@ul
* No external forces forces
* Fibrous, messy looking
@ulend

+++
#### Actor–network theory
2\. Insistence on human and non-human _actants_
+++
#### Actor–network theory
2\. Insistence on human and non-human _actants_
```python
def list_actants():
    return [jenkins, some_module, you, me, product_owner, team_lead]
```
---
#### The Obligatory Passage Point (OPP)

![OPP](https://raw.githubusercontent.com/jondequinor/talks/master/opp/img/generic_opp.jpg)

---
#### The Obligatory Passage Point (OPP)

![OPP](https://raw.githubusercontent.com/jondequinor/talks/master/opp/img/generic_opp.jpg)

Actant needs to go through in order to score goal
---

---
## Takk
---
+++
EOF
