package main

import "fmt"
import "./driver"
	import "time"



func main() {
	driver.Driver_init()
	fmt.Println("test",driver.Get_stop_signal())
	driver.Set_motor_direction(1)
	for i:=0; i<4; i++ {
		time.Sleep(500*time.Millisecond)
		if i<3 {
			driver.Set_button_lamp(1,i,1)
			driver.Set_button_lamp(0,i,1)
		}
		driver.Set_button_lamp(2,i,1)
	}
	driver.Set_floor_indicator(2)
	pressed:=[]int{0,0,0,0}
	for{
		for i:=0; i<4; i++ {
			if driver.Get_button_signal(2,i)==0 {
				pressed[i]=0;
			}
			if driver.Get_button_signal(2,i)==1 && pressed[i]==0 {
				pressed[i]=1;
				fmt.Print("Button 2, ",i," is pressed!\r")
				if i==3 {
					driver.Set_door_lamp(1)
					driver.Set_stop_lamp(1)
				}
				if i==2 {
					driver.Set_door_lamp(0)
					driver.Set_stop_lamp(0)
				}
			}
		}
		if driver.Get_floor_sensor_signal()==3 {
			driver.Set_motor_direction(-1)
		}else if driver.Get_floor_sensor_signal()==0 {
			driver.Set_motor_direction(1)
		}
	}
}
