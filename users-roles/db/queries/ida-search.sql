select 
	id
from
	identity_attribute
where
	1=1
	{{ with .AssignedToParticipant -}} and assignedToParticipant = {{ param . }}{{ end }}
	{{ with .Code -}} and code = {{ param . }}{{ end }}
	{{ with .Enabled -}} and enabled = {{ param . }}{{ end }}
	{{ with .Id -}} and id = {{ param . }}{{ end }}
	{{ with .Name -}} and name = {{ param . }}{{ end }}
	{{ with .ParticipantTypeIn -}} and participantTypeIn = {{ param . }}{{ end }}
	{{ with .ParticipantTypeNotIn -}} and participantTypeNotIn = {{ param . }}{{ end }}
	{{ with .UpdateTimestampFrom -}} and updateTimestampFrom = {{ param . }}{{ end }}
	{{ with .UpdateTimestampTo -}} and updateTimestampTo = {{ param . }}{{ end }}
