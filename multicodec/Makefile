
build:
	go build

install:
	go install

examples: build
	cat examples/one.json | ./multicodec -r /cbor >examples/one.cbor
	cat examples/many.json | ./multicodec -r /cbor >examples/many.cbor
	cat examples/one.json >examples/four.mixed
	cat examples/one.cbor >>examples/four.mixed
	cat examples/one.json >>examples/four.mixed
	cat examples/one.cbor >>examples/four.mixed
	cat examples/four.mixed | ./multicodec -r /json >examples/four.unmixed.json
	cat examples/four.mixed | ./multicodec -r /cbor >examples/four.unmixed.cbor
	cat examples/four.mixed | ./multicodec --headers >examples/four.headers
	cat examples/four.mixed | ./multicodec --paths >examples/four.paths
