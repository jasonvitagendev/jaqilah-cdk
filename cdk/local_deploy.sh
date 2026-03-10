export CDK_ACCOUNT=125084610626
export CDK_REGION=ap-southeast-1

#export DOMAIN_NAME=beachy.vip
#export PROJECT_NAME=beachy
#export HOSTED_ZONE_ID=Z00143171AP5N24Z3VX72
#export EMAIL_ADDRESS=help@beachy.vip
#export EMAIL_NAME="Beachy Help"
#export AUTH_DOMAIN_NAME=auth.beachy.vip
#export API_DOMAIN_NAME=api.beachy.vip
#export COGNITO_USER_POOL_CLIENT_ID=46c4ok11j8hc20gh88h36bdceu
#export COGNITO_USER_POOL_CLIENT_SECRET=1em7bbraeimffkml7lvo7lo0mt17hnbmi20gl2jqhlttum56ohnh
#export COGNITO_USER_POOL_ID=ap-southeast-1_VlCX1XHfr
#export ENV=prod
#export HOSTED_ZONE_ID_WEBJIRAN=Z09539432EWHBJ25N64QC
#cdk destroy --all
#cdk bootstrap --force
#cdk list

#cdk deploy BeachyStaticSiteStack
#cdk deploy BeachyDBStack
#cdk deploy BeachyApiEndpointsStack
#cdk bootstrap

cdk deploy -vv --require-approval never --all
