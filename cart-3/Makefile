# Build dependencies
WASM_OPT = wasm-opt

# Compilation flags
FLAGS = -o:size -target:freestanding_wasm32 -no-entry-point -extra-linker-flags:"--import-memory -zstack-size=14752 --initial-memory=65536 --max-memory=65536 --stack-first --lto-O3 --gc-sections --strip-all"

# wasm-opt flags
WASM_OPT_FLAGS = -Oz --enable-mutable-globals --enable-nontrapping-float-to-int --enable-bulk-memory-opt

all:
	@mkdir -p build
	odin build src $(FLAGS) -out:build/cart.wasm
ifneq ($(DEBUG), 1)
ifeq (, $(shell command -v $(WASM_OPT)))
	@echo Tip: $(WASM_OPT) was not found. Install it from binaryen for smaller builds!
else
	$(WASM_OPT) $(WASM_OPT_FLAGS) build/cart.wasm -o build/cart.wasm
endif
endif

.PHONY: clean
clean:
	rm -rf build

.PHONY: release
release: all
	w4 bundle --html    build/cart3.html  --title cart0 build/cart.wasm
	w4 bundle --windows build/cart3.exe   --title cart0 build/cart.wasm
	w4 bundle --mac     build/cart3.mac   --title cart0 build/cart.wasm
	w4 bundle --linux   build/cart3.linux --title cart0 build/cart.wasm

