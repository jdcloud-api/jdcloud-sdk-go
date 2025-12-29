# 更新历史

API版本：2.7.45

| 发布时间       | 版本号    | 更新     | 说明                                                                                                                                                                                                    |
|------------|--------| -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 2025-12-02 | 2.7.46 | 新增接口 | 新增接口 modifyInstanceType，用于修改实例架构类型（标准版架构变更为集群架构） |
| 2025-11-18 | 2.7.45 | 新增接口 | 更新describeBillingInstances接口 |
| 2025-11-14 | 2.7.44 | 更新接口 | 更新describeInstanceTLS接口 |
| 2025-10-28 | 2.7.43 | 更新接口 | 更新describeInstanceTLS、modifyInstanceTLS接口 |
| 2025-10-23 | 2.7.42 | 新增接口 | 新增配置模板管理接口：createConfigTemplate(创建)、describeConfigTemplates(列表)、describeConfigTemplate(详情)、modifyTemplate(修改)、deleteConfigTemplate(删除)；新增配置模板相关模型定义 |
| 2025-10-15 | 2.7.41 | 新增接口 | 新增接口 listTaskTypes，用于获取任务类型列表；新增taskType模型定义 |
| 2025-06-24 | 2.7.40 | 新增字段 | 增加typeCn字段 |
| 2024-12-19 | 2.7.39 | 新增接口 | 新增接口 modifyBlockStatus |
| 2024-12-16 | 2.7.38 | 新增接口 | 新增接口 describeChartReleases、createChartRelease、modifyChartRelease |
| 2024-12-10 | 2.7.37 | 新增接口 | 新增接口 createImageRelease、modifyImageRelease |
| 2024-12-10 | 2.7.36 | 新增接口 | 新增接口 describeImageReleases |
| 2024-12-09 | 2.7.35 | 新增接口 | 新增接口 modifyInstanceNodeGroupVersion、modifyInstanceVersion |
| 2024-11-22 | 2.7.34 | 更新接口 | 增加接口: haDiagnosis、modifySentinel、describeAvailableSentinelList、describeSentinelAvailableZones、modifySentinelAvailableZones、modifyRedisAZSpecifyType、sentinelNode |
| 2024-11-05 | 2.7.33 | 更新接口 | 修改接口 migrateProxyAvailableZones、migrateRedisAvailableZones |
| 2024-10-14 | 2.7.32 | 更新接口 | 修改接口 describeProxyNodeList, 增加az、address入参，增加azs出参 |
| 2024-09-27 | 2.7.31 | 新增接口 | 增加接口 describeProxyNodeList, AzId/AzIdSpec 新增azsForProxy参数 |
| 2024-08-26 | 2.7.30 | 新增接口 | 增加接口 checkDeletable |
| 2024-07-30 | 2.7.29 | 新增接口 | 增加接口 checkPasswordValid |
| 2024-07-18 | 2.7.28 | 新增接口 | 增加接口 groupWhiteList 增、删、改、查接口 |
| 2024-06-11 | 2.7.27 | 更新接口 | 更新接口 restartInstance, 支持重启代理节点 |
| 2024-06-04 | 2.7.26 | 更新接口 | 更新接口 describeUpgradeVersion、modifyInstanceMinorVersion、cacheInstance, 增加参数支持查询和升级代理镜像 |
| 2024-04-22 | 2.7.25 | 新增接口 | 新增接口 modifyPublicAddress |
| 2024-04-17 | 2.7.24 | 新增接口 | 新增接口 detailNode |
| 2024-03-28 | 2.7.23 | 新增接口 | 新增接口 importData |
| 2024-03-04 | 2.7.22 | 更新接口 | 更新接口 createCacheInstance, 增加入参 |
| 2023-12-29 | 2.7.21 | 新增接口 | 新增接口 describeInstanceTLS、modifyInstanceTLS |
| 2023-12-27 | 2.7.20 | 新增接口 | 新增修改实例自定义域名接口modifyInstanceCustomDomain, 修改接口describeCacheInstance, RespExtension新增domainSuffix属性 |
| 2023-12-25 | 2.7.19 | 新增接口 | 新增接口：cacheInstanceRecycle查询回收站实例列表、cacheInstanceRecycle/{cacheInstanceId}强制销毁回收站实例 |
| 2023-11-17 | 2.7.18 | 新增接口 | 新增接口modifyMaintenanceTime、modifyTaskRunTime、cancelTask，createCacheInstance接口返回字段增加运维时间字段 |
| 2023-05-23 | 2.7.17 | 新增接口 | 新增接口switchInstanceHA, 增加describeNodeList接口字段 |
| 2023-05-15 | 2.7.16 | 更新接口参数 | modifyCacheInstanceClass接口, 新增参数 replicaNumber |
| 2023-04-25 | 2.7.15 | 新增接口 | 新增接口 describeUpgradeVersion、modifyInstanceMinorVersion |
| 2023-03-16 | 2.7.14 | 更新接口描述 | describeSlowLog更新字段描述和案例 |
| 2023-03-14 | 2.7.13 | 更新接口 | describeSlowLog响应增加clientAddr和clientName |
| 2023-03-13 | 2.7.12 | 更新接口 | 修改describeSlowLog请求格式，补充参数 |
| 2023-03-07 | 2.7.11 | 更新接口 | 修改describeSlowLog请求格式，删除describeHistorySlowLog接口 |
| 2023-03-06 | 2.7.10 | 更新接口 | 修改describeHistorySlowLog响应格式 |
| 2023-03-04 | 2.7.9  | 新增接口 | 新增内部接口getHistorySlowLogConfig、describeHistorySlowLog |
| 2023-02-23 | 2.7.8  | 新增接口 | 新增内部接口describeListFilterAndSort, 修改接口describeClusterInfo |
| 2023-02-16 | 2.7.7  | 新增接口 | 新增备份详情接口describeBackupInfo,增加describeBackups接口参数,创建实例接口增加备份数据克隆参数 |
| 2023-02-07 | 2.7.6  | 新增接口 | 新增纳管集群接口externalInstance、describeExternalInstanceVersions、checkExternalInstanceConnectivity |
| 2023-01-15 | 2.7.5  | 更新接口 | 修改executeMultiCommand接口，参加入参role |
| 2022-12-30 | 2.7.4  | 更新接口 | 修改executeCommandHistory接口的x-jdcloud-internal: true |
| 2022-12-30 | 2.7.3  | 更新接口 | 修改restartInstance接口的x-jdcloud-internal: true |
| 2022-12-29 | 2.7.2  | 新增接口 | 新增接口executeMultiCommand(内部)，executeCommandHistory |
| 2022-12-27 | 2.7.1  | 新增接口 | 新增接口restartInstance |
| 2022-12-22 | 2.7.0  | 更新接口 | 更新接口createCacheInstance、describeCacheInstance、describeCacheInstances |
| 2022-12-19 | 2.6.32 | 新增接口 | 新增接口restartProxy、restartProxy |
| 2022-12-16 | 2.6.31 | 更新接口 | 修改接口createBigKeyAnalysis、createCacheAnalysis的请求参数 |
| 2022-12-01 | 2.6.30 | 新增接口 | 新增接口describeResizeModeIpTimeInfo, modifyCacheInstanceClass接口增加入参 |
| 2022-11-15 | 2.6.29 | 新增接口 | 内部接口调整 |
| 2022-11-07 | 2.6.28 | 更新接口 | 修改接口describeHotKeyDetail2的返回字段 |
| 2022-10-11 | 2.6.27 | 更新接口 | 修改接口describeNodeList的返回字段 |
| 2022-09-26 | 2.6.26 | 更新接口 | 修改接口describeCacheInstances、describeCacheInstance的返回字段 |
| 2022-09-20 | 2.6.25 | 新增接口 | 内部接口调整 |
| 2022-09-05 | 2.6.24 | 新增接口 | 新增接口describeHotKeySummary、describeHotKeyDetail |
| 2022-08-19 | 2.6.23 | 更新接口 | 修改接口describeClientPerfData、filteredClientPerfData的返回字段 |
| 2022-08-16 | 2.6.22 | 新增接口 | 新增接口modifyAccounts |
| 2022-08-15 | 2.6.21 | 新增接口 | 新增接口checkInstances |
| 2022-08-05 | 2.6.20 | 新增接口 | 新增接口createBigKeyAnalysis2、describeBigKeyList2、describeBigKeyDetail2、modifyBigKeyAnalysisTime2、describeBigKeyAnalysisTime2、modifyAnalysisThreshold2、describeAnalysisThreshold2、describeHotKeyResult2、describeHotKeyDetail2 |
| 2022-08-01 | 2.6.19 | 新增接口 | 新增接口describeAvailableResource2 |
| 2022-07-20 | 2.6.18 | 新增接口 | 更新接口modifyAnalysisThreshold、describeBigKeyList |
| 2022-07-06 | 2.6.17 | 新增接口 | 新增接口bigKey、bigKeyDetail、bigKeyAutoAnalysisTime、stopCacheAnalysis、cacheAnalysisThreshold |
| 2022-05-31 | 2.6.16 | 新增接口 | 新增接口describeNodeList，更新接口executeCommand |
| 2022-05-23 | 2.6.15 | 接口更新 | 更新接口createCacheInstance、describeCacheInstance、describeCacheInstances、describeAvailableResource    |
| 2022-05-20 | 2.6.14 | 新增接口 | 新增接口filteredClientPerfData    |
| 2022-05-17 | 2.6.13 | 接口更新 | 更新接口describeInstanceConfig        |
| 2022-05-12 | 2.6.12 | 接口更新 | 修改接口describeClearData的方法为get        |
| 2022-05-12 | 2.6.11 | 新增接口 | 新增接口describeClientPerfData        |
| 2022-05-06 | 2.6.10 | 新增接口 | 新增接口describeConfigCenterTokenAndCipher、describeClientSumUseR2MJavaClient、describeClientDetailUseR2MJavaClient        |
| 2022-04-08 | 2.6.9  | 新增接口 | 新增接口startClearData、stopClearData、describeClearData                                                                                                                                                |
| 2022-01-10 | 2.6.8  | 接口更新 | 更新接口describeInstanceConfig                                                                                                                                                                          |
| 2021-12-27 | 2.6.7  | 新增接口 | 新增接口describeUnSupportedFunction，支持白名单功能                                                                                                                                                     |
| 2021-12-07 | 2.6.6  | 新增接口 | 新增接口createAccount、describeAccounts、modifyAccount、deleteAccount，支持账号管理功能                                                                                                                 |
| 2021-12-07 | 2.6.5  | 接口更新 | 更新接口modifyBackupPolicy、describeBackupPolicy，增加autoBackup参数                                                                                                                                    |
| 2021-11-18 | 2.6.4  | 新增接口 | 新增接口setDisableCommands、getDisableCommands，支持设置、查询禁用命令                                                                                                                                  |
| 2021-11-08 | 2.6.3  | 接口更新 | 更新接口createCacheInstance、describeCacheInstance、describeCacheInstances、availableResource                                                                                                           |
| 2021-09-14 | 2.6.2  | 接口更新 | 更新接口createCacheInstance参数                                                                                                                                                                         |
| 2021-09-13 | 2.6.1  | 新增接口 | 内部接口调整                                                                                                                                                                                            |
| 2021-08-03 | 2.6.0  | 新增接口 | 内部接口调整                                                                                                                                                                                            |
| 2021-03-25 | 2.5.0  | 接口更新 | 更新接口describeClusterInfo； 新增接口describeAvailableRegion、describeAvailableResource                                                                                                                |
| 2021-03-02 | 2.4.0  | 新增接口 | 更新接口createCacheInstance、describeCacheInstance、describeCacheInstances，新增接口setExposeType，支持设置外部访问方式，新增接口describeExposeType，查询支持的外部访问方式列表                         |
| 2021-1-25  | 2.3.0  | 新增接口 | 新增describeTaskProgressList接口，查询实例的任务进度列表                                                                                                                                                |
| 2020-12-14 | 2.2.3  | 接口更新 | createCacheInstance 增加标签                                                                                                                                                                            |
| 2020-09-23 | 2.2.2  | 增加接口 | 内部接口调整                                                                                                                                                                                            |
| 2020-08-17 | 2.2.1  | 接口更新 | 内部接口调整                                                                                                                                                                                            |
| 2020-07-29 | 2.2.0  | 增加接口 | 新增接口describeClientList、describeClientIpDetail，内部接口调整                                                                                                                                        |
| 2020-05-13 | 2.1.2  | 接口更新 | 更新describeSpecConfig接口返回结果                                                                                                                                                                      |
| 2020-03-27 | 2.1.1  | 接口更新 | 更新describeSpecConfig接口返回结果                                                                                                                                                                      |
| 2020-03-17 | 2.1.0  | 接口更新 | createCacheInstance、modifyCacheInstanceClass接口新增shardNumber参数；新增describeSpecConfig接口                                                                                                        |
| 2020-03-16 | 2.0.0  | 增加接口 | 发布IP白名单接口：describeIpWhiteList、modifyIpWhiteList；新增大key热key分析接口：describeAnalysisTime、describeCacheAnalysisList、describeCacheAnalysisResult、modifyAnalysisTime、createCacheAnalysis |
| 2019-12-02 | 1.9.0  | 增加接口 | 内部接口调整                                                                                                                                                                                            |
| 2019-10-28 | 1.8.0  | 接口更新 | 修改createCacheInstance、modifyCacheInstanceClass接口返回参数，更新文档，内部接口调整                                                                                                                   |
| 2019-10-15 | 1.7.0  | 接口更新 | 发布接口：describeSlowLog                                                                                                                                                                               |
| 2019-08-27 | 1.6.0  | 接口更新 | 内部接口调整                                                                                                                                                                                            |
| 2019-08-22 | 1.5.0  | 接口更新 | 发布接口：describeInstanceConfig、modifyInstanceConfig、describeBackupPolicy、modifyBackupPolicy、createBackup、describeBackups、describeDownloadUrl、restoreInstance                                   |
| 2019-08-14 | 1.4.0  | 增加接口 | 内部接口调整                                                                                                                                                                                            |
| 2019-07-23 | 1.3.0  | 接口更新 | 发布接口：describeClusterInfo                                                                                                                                                                           |
| 2019-05-27 | 1.2.0  | 增加接口 | 内部接口调整                                                                                                                                                                                            |
| 2019-05-07 | 1.1.0  | 接口更新 | 内部接口调整                                                                                                                                                                                            |
| 2019-04-15 | 1.0.5  | 增加接口 | 内部接口调整                                                                                                                                                                                            |
| 2019-01-14 | 1.0.4  | 增加接口 | 内部接口调整                                                                                                                                                                                            |
| 2018-12-29 | 1.0.3  | 接口更新 | 查询规格，返回参数的instanceClass类型新增instanceType字段                                                                                                                                               |
| 2018-12-20 | 1.0.2  | 接口更新 | 支持创建redis 4.0版本的缓存引擎，openAPI功能不变，创建、查询规格时新增redisVersion请求参数                                                                                                              |
| 2018-08-08 | 1.0.1  | 接口更新 | 列表详情接口返回实例版本instanceVersion，创建和重置密码接口密码为空即为免密                                                                                                                             |
| 2018-03-31 | 1.0.0  | 增加接口 | 初始版本，缓存redis基本操作接口                                                                                                                                                                         |
