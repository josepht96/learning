variable "example_variable" {
  description = "some desc"
  type        = string
  default     = "hello"
}

variable "list_example" {
  description = "blah"
  type        = list(any)
    default = [1, 2, 3]
}
