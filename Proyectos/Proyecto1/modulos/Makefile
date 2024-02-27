obj-m += /ram/ram_so1_1s2024.o /cpu/cpu_so1_1s2024.o

all:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules

clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean

mount:
	sudo $(MAKE) insmod

unmount:
	sudo $(MAKE) rmmod

insmod:
	insmod ram/ram_so1_1s2024.ko
	insmod cpu/cpu_so1_1s2024.ko

rmmod:
	rmmod ram/ram_so1_1s2024.ko
	rmmod cpu/cpu_so1_1s2024.ko