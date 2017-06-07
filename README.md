# It monitors the specified log and processes it.
If you monitor a specified log and detect a specific character, the EC2 tag is updated.

# Write the setting in TOML.
Set TagStartValue to TagName when finding the character specified in SearchStart from TailFile.  
Set TagEndValue to TagName when finding the character specified in SearchEnd from TailFile.  

## Example for TOML.
```
TailFile = "./example.log" //Path of file to tail.
PostionFile = "./monitor.pos" //Path of position file.
SearchStart = "aaaaa" //First character to search.
SearchEnd = "bbbbb" //End character to search.
TagName = "cw_alert" //Tag name to be set for EC2.
TagStartValue = "off" //The value of the tag to be set to EC 2 when the character of SearchStart is found.
TagEndValue = "on" //The value of the tag to be set to EC 2 when the character of SearchEnd is found.
Delay = 3000 //Delay time until tag setting.
```
