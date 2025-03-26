select 
	id
from
	identity_attribute
where
	1=1
	{{ with .Id -}} and id = {{ param . }}{{ end }}
