RFC Output Generator

The RFC output generator takes OpenAPI spec document and ultimately
translates it to a Word document in order to simplify loading it into
Google Docs.

The `widderdoc.sh` script accomplishes this by using widdershins to
convert the OpenAPI specification document to MarkDown. Next it uses
`pandoc` to convert that MarkDown to `docx` format. The script expects
the spec to be named openapi.yaml in a folder named spec one folder up
from the folder containing the script.
