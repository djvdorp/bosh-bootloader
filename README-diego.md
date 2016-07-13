# Diego installation
---

## Steps

### Starting point

This manual assumes that you have installed [CloudFoundry](README-cloudfoundry.md) first.

### Additional dependencies

- `bosh upload release https://bosh.io/d/github.com/cloudfoundry-incubator/etcd-release`
- `bosh upload release https://bosh.io/d/github.com/cloudfoundry-incubator/garden-runc-release`
- `bosh upload release https://bosh.io/d/github.com/cloudfoundry-incubator/garden-linux-release`
- `bosh upload release https://bosh.io/d/github.com/cloudfoundry/cflinuxfs2-rootfs-release`
- `bosh upload release https://bosh.io/d/github.com/cloudfoundry/diego-release?v=0.1476.0` (`v0.1476.0` is recommended for CF `v238`)

### Actual Diego installation

**Note:** this information is based on [this issue](https://github.com/cloudfoundry/diego-release/issues/171) and [this document](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md).

1. [Configure Additional Infrastructure](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md#configure-additional-infrastructure): this is a matter of creating 3 [Diego Subnets](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md#diego-subnets), since `bbl` has already created a [Load Balancer](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md#load-balancer) for you.
2. In order to generate a Diego deployment manifest you need to at least specify [3 values when generating a manifest](https://github.com/cloudfoundry/diego-release/blob/develop/scripts/generate-deployment-manifest#L16-L18). You will need to generate a host key for the ssh-proxy. [This documentation](https://github.com/cloudfoundry/diego-release/tree/develop/examples/aws#generating-ssh-proxy-host-key-and-fingerprint) covers how to do so.
3. The (compiled) CloudFoundry manifest `cf-deployment.yml` should be fine here, but you do need to add the [following properties](https://github.com/cloudfoundry/diego-release/tree/develop/examples/minimal-aws#modify-the-cf-manifest-with-diego-properties) specified in the first code example of this documentation in order to get `cf ssh` to work. Doing this the proper way, you will modify your `cf-stub.yml` and re-run the `generate_deployment_manifest` command again. You will need to add the `ssh-proxy-host-key-fingerprint` generated above. Everything else is [explained here](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md#configure-cf-manifest-for-diego).
4. For the `iaas-settings.yml`, [this template](https://github.com/cloudfoundry/diego-release/blob/develop/examples/aws/templates/diego/iaas-settings-internal.yml) is a good example. You will need to replace the various spiff merges with corresponding identifiers from the aws resources that you will need to create. (Specifically you will need to create 3 subnets for Diego's 3 availability zones, and 1 ELB for the ssh-proxy; this is what you did in step 1)
5. For the `property-overrides.yml`, [this template](https://github.com/cloudfoundry/diego-release/blob/develop/examples/aws/templates/diego/property-overrides.yml) is a good example. Once again you will need to supply values from generating the `ssh-proxy-host-key.pem`. The certificates that you need to generate will be covered in [this documentation](https://github.com/cloudfoundry/diego-release/blob/develop/examples/aws/README.md#configuring-security).
6. [Generate the Diego Deployment Manifest](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md#generate-diego-deployment-manifest).
7. [Upload Additional Releases](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md#upload-additional-releases), this was listed in the *Additional dependencies* above.
8. [Upload Diego Release](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md#create-and-upload-diego-release), this was also listed in the *Additional dependencies* above.
9. Run `bbl lbs` to get the Amazon ELB endpoint for the `CFSSHPro`, then configure a new `CNAME` record for `ssh.<domain>` to this ELB.
10. Add an additional security group to the (compiled) Manifests' `resource_pool` `cloud_properties` for your `access_z*` VMs, this additional security group allows the ingress from our ELB that's required and is created by the previous `bbl unsupported-create-lbs --type cf` command earlier.
11. Then do a `bosh deploy` for your CloudFoundry deployment and wait patiently. **Keep in mind that at this point your CloudFoundry installation is not functional anymore!**
12. [Deploy Diego](https://github.com/cloudfoundry/diego-release/blob/develop/docs/deploy-alongside-existing-cf.md#deploy-diego) and wait patiently. After this, your CloudFoundry deployment should be fully functional again and have Diego enabled globally by default!
