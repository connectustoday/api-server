# Reminder
When you are adding new fields to a struct, remember to:
* Add the field to interfaces-internal
* Add the field to interfaces-api (if applicable)
* Add the field to interfaces-conv conversions (if applicable)
* Check if registration requires the field