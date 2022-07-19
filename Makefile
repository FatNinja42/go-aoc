YEAR:=y2015
DAY_PREFIX:=day
TEMPLATE_POSTFIX:=x

cpath := $(shell pwd)/$(YEAR)/$(DAY_PREFIX)

new:
	@if [ ! -d  $(cpath)$(day) ]; then\
		cp -R $(cpath)$(TEMPLATE_POSTFIX) $(cpath)$(day) ;\
		find  $(cpath)$(day) -name "*.go" | xargs sed -i '' 's/dayx/day$(day)/g';\
  	else\
  		echo "Day $(day)" already exists.;\
  	fi
