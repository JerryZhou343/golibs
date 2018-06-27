package common

var(
    BUILD_VERSION ="xxx.xxx.xx"
    BUILD_DATE_TIME = "xxxx-xx-xx xx:xx:xx"
    ApplicationName = "Application"
    OrganizationName = "xxxx"
)



func GetVersionInfo()string{
    return "\n===========================================\n"+
        "Copyright(C) " + OrganizationName + "Ltd. all right reserved.\n" +
        "Application: " + ApplicationName +  "\n" +
        "COMPLIE VERSION: " + BUILD_VERSION + "\n" +
        "COMPLIE TIME: " + BUILD_DATE_TIME + "\n"+
        "============================================\n"
}