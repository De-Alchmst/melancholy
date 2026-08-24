build/cart.wasm: src/*.odin
	@-mkdir -p build
	odin build src -o:size -out:build/cart.wasm -target:freestanding_wasm32 -no-entry-point -extra-linker-flags:"--import-memory -zstack-size=14752 --initial-memory=65536 --max-memory=65536 --stack-first --lto-O3 --gc-sections --strip-all"
	# add wasmopt if needed
	# wasm-opt -Oz --enable-mutable-globals --enable-nontrapping-float-to-int --enable-bulk-memory-opt ./build/cart.wasm -o ./build/cart.wasm
