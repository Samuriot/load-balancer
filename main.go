package main


import (
	"fmt"
	"net"
	"regexp"
)

// from a high level overview, a load balancer is basically a router
// it routes traffic across servers, depending on the algorithm chosen
// round robin, weighted round robin, etc

// components we need so far:
// Array to store servers
// counter to keep track of which server to send to next

type LoadBalancer interface {
	Init() bool
	ProcessPacket() bool
}

type Poller interface {
	Poll() ([byte], error)
}

type UDP_Poller struct {
	conn *net.UDPConn
}

type L4_LoadBalancer struct {
	backend_servers []net.IP
}

func (lb *L4_LoadBalancer) Init() bool {
	tcp_sock, tcp_err := net.Listen("tcp", ":8080");
	if tcp_err != nil {
		fmt.Println(tcp_err);
		return;
	}

	udp_sock, udp_err := net.Listen("udp", ":8080");
	if udp_err != nil {
		fmt.Println(err);
		return;
	}

	go Poll(&tcp_sock);
	go Poll(&udp_sock);
}

func (lb * L4_LoadBalancer) Poll() bool {
	tcp_sock, tcp_err := net.Listen("tcp", ":8080");
	if tcp_err != nil {
		fmt.Println(tcp_err);
		return;
	}

	udp_sock, udp_err := net.Listen("udp", ":8080");
	if udp_err != nil {
		fmt.Println(err);
		return;
	}

	for { // golang doesn't have while true...
		conn, err := sock.Accept();
		if err != nil {
			fmt.Println(err);
			continue;
		}

		go ProcessPacket(conn);
	}
}

func check_ip_pattern(ip_addr string) bool {
	re := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	if (re.MatchString(ip_addr)) {
		return true;
	}

	return false;
}

func add_server(ip_addr string, ip_list *[]net.IP) bool {
	if(ip_addr == "") {
		return false;
	}

	if (!check_ip_pattern(ip_addr)) {
		return false;
	}

	ip_obj := net.ParseIP(ip_addr);
	*ip_list = append(*ip_list, ip_obj);
	return true;
}

func main() {
	ip_list := []net.IP{};
	ip_test := "192.168.1.1";
	success := check_ip_pattern("192.168.1.1");
	if (success) {
		fmt.Printf("This is an IP\n");
	}
	add_server(ip_test, &ip_list);
	for idx, ip := range ip_list {
		fmt.Println(idx, ip);
	}
}
