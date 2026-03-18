logical_product_family  = "lp"
logical_product_service = "lps"
class_env               = "dev"
instance_env            = "01"
instance_resource       = 1

resource_names_map = {
  eventbus = {
    name       = "eventbus"
    max_length = 64
  }
}

description = "Example EventBridge event bus for testing"

tags = {
  Environment = "test"
  Module      = "cloudwatch-event-bus"
}
