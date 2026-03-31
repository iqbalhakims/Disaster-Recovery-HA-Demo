locals {
  cluster_tag = "k8s:${digitalocean_kubernetes_cluster.prod.id}"
}

resource "digitalocean_monitor_alert" "cpu_high" {
  description = "DOKS node CPU > 80%"
  type        = "v1/insights/droplet/cpu"
  compare     = "GreaterThan"
  value       = 80
  window      = "5m"
  enabled     = true
  tags        = [local.cluster_tag]

  alerts {
    email = var.alert_email
  }
}

resource "digitalocean_monitor_alert" "memory_high" {
  description = "DOKS node memory > 80%"
  type        = "v1/insights/droplet/memory_utilization_percent"
  compare     = "GreaterThan"
  value       = 80
  window      = "5m"
  enabled     = true
  tags        = [local.cluster_tag]

  alerts {
    email = var.alert_email
  }
}

resource "digitalocean_monitor_alert" "disk_high" {
  description = "DOKS node disk > 85%"
  type        = "v1/insights/droplet/disk_utilization_percent"
  compare     = "GreaterThan"
  value       = 85
  window      = "5m"
  enabled     = true
  tags        = [local.cluster_tag]

  alerts {
    email = var.alert_email
  }
}
