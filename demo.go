package main

import (
	"log"
	vpplink "vppManager/api"
	"vppManager/api/types"
	"vppManager/net"
)

func createVppInt(){
	link, _ := vpplink.NewVppLink("", nil)
	err := link.Reconnect()
	if err != nil{
		log.Print("connect : ", err)
	}

	nics, err := net.GetIntfacesInfo()
	if err != nil{
		log.Print("get local nics' info Error : ", err)
	}

	for i := 0; i < len(nics); i++{

		hostInt := types.GenericVppInterface{
			//Name: nics[i].Name,
			HostInterfaceName: nics[i].Name,
			HardwareAddr: &nics[i].Mac,
		}

		intId, err := link.CreateAfPacket(&types.AfPacketInterface{GenericVppInterface:hostInt})
		if err != nil{
			log.Print("create interface : ", err)
			return
		}

		err = link.InterfaceAdminUp(intId)
		if err != nil{
			log.Print("enable vpp interface Error : ", err)
			return
		}

		err = link.AddInterfaceAddress(intId, nics[i].Ipv4)
		if err != nil{
			log.Print("set vpp interface ip address Error : ", err)
			return
		}
		log.Print("success : ", intId)
	}

}

func main(){
	createVppInt()
}
