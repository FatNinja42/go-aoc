YEAR:=y2015
TEMPLATE_PREFIX:=day
TEMPLATE_POSTFIX:=x
TEMPLATE_FOLDER:=template

cpath := $(shell pwd)/$(YEAR)/$(TEMPLATE_PREFIX)
template_path=$(shell pwd)/$(TEMPLATE_FOLDER)/$(TEMPLATE_PREFIX)

new:
	@if [ ! -d  $(cpath)$(day) ]; then\
		cp -R $(template_path)$(TEMPLATE_POSTFIX) $(cpath)$(day) ;\
		find  $(cpath)$(day) -name "*.go" | xargs sed -i '' 's/dayx/day$(day)/g';\
  	else\
  		echo "Day $(day)" already exists.;\
  	fi