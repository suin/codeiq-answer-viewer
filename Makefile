VIEW="codeiq/view.go"

dev:
	echo "// This file is generated by Makefile" > $(VIEW)
	echo "package codeiq" >> $(VIEW)

	echo "const Template = \`" >> $(VIEW)
	cat assets/template.html >>  $(VIEW)
	echo "\`" >> $(VIEW)

	echo "const Script = \`"  >> $(VIEW)
	coffee -c -p assets/script.coffee >> $(VIEW)
	echo "\`" >> $(VIEW)

	go fmt ./...
	go run server.go