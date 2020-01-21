package scanner

// UNKNOWN default for unknown port
const UNKNOWN = "Unknown"

// KNOWNPORTS are sample known ports
var KNOWNPORTS = map[int]string{
	27017: "mongodb [ http://www.mongodb.org/ ]",
	28017: "mongodb web admin [ http://www.mongodb.org/ ]",
	21:    "ftp",
	22:    "SSH",
	23:    "telnet",
	25:    "SMTP",
	66:    "Oracle SQL*NET?",
	69:    "tftp",
	80:    "http",
	88:    "kerberos",
	109:   "pop2",
	110:   "pop3",
	123:   "ntp",
	137:   "netbios",
	139:   "netbios",
	443:   "https",
	445:   "Samba",
	631:   "cups",
	5800:  "VNC remote desktop",
	194:   "IRC",
	118:   "SQL service?",
	150:   "SQL-net?",
	1433:  "Microsoft SQL server",
	1434:  "Microsoft SQL monitor",
	3306:  "MySQL",
	3396:  "Novell NDPS Printer Agent",
	3535:  "SMTP (alternate)",
	554:   "RTSP",
	9160:  "Cassandra [ http://cassandra.apache.org/ ]",
	8000:  "Nodejs",
	9200:  "Elasticsearch",
	5601:  "Kibana",
}

func predictPort(port int) string {
	if result, ok := KNOWNPORTS[port]; ok {
		return result
	}
	return UNKNOWN
}
