#Set your public SSH key here
variable "ssh_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDcwZAOfqf6E6p8IkrurF2vR3NccPbMlXFPaFe2+Eh/8QnQCJVTL6PKduXjXynuLziC9cubXIDzQA+4OpFYUV2u0fAkXLOXRIwgEmOrnsGAqJTqIsMC3XwGRhR9M84c4XPAX5sYpOsvZX/qwFE95GAdExCUkS3H39rpmSCnZG9AY4nPsVRlIIDP+/6YSy9KWp2YVYe5bDaMKRtwKSq3EOUhl3Mm8Ykzd35Z0Cysgm2hR2poN+EB7GD67fyi+6ohpdJHVhinHi7cQI4DUp+37nVZG4ofYFL9yRdULlHcFa9MocESvFVlVW0FCvwFKXDty6askpg9yf4FnM0OSbhgqXzD austin@EARTH"
}
#Establish which Proxmox host you'd like to spin a VM up on
variable "proxmox_host" {
    default = "pve"
}
#Specify which template name you'd like to use
variable "template_name" {
    default = "ubuntu-2204-cloudinit-template"
}
#Establish which nic you would like to utilize
variable "nic_name" {
    default = "vmbr0"
}
#Provide the url of the host you would like the API to communicate on.
#It is safe to default to setting this as the URL for what you used
#as your `proxmox_host`, although they can be different
variable "api_url" {
    default = "https://192.168.88.33:8006/api2/json"
}

#Blank var for use by terraform.tfvars
variable "token_secret" {
    default = "29d0a9ea-515c-49f3-a741-3af01498bd3a"
}
#Blank var for use by terraform.tfvars
variable "token_id" {
    default = "terraform@pve!terraform-token"
}