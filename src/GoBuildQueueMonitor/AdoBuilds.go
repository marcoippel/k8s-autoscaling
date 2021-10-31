package main

type AdoBuilds struct {
	Fps struct {
		Content []struct {
			ClientID         string   `json:"clientId"`
			ContentLength    int64    `json:"contentLength"`
			ContentType      string   `json:"contentType"`
			ContributionID   string   `json:"contributionId"`
			Integrity        string   `json:"integrity"`
			ModuleNamespaces []string `json:"moduleNamespaces"`
			Priority         int64    `json:"priority"`
			URL              string   `json:"url"`
		} `json:"content"`
		DataProviders struct {
			ClientProviders interface{} `json:"clientProviders"`
			Data            struct {
				Ms_vss_admin_web_collection_admin_hub_contribution_data_provider struct {
					Contributions []interface{} `json:"contributions"`
				} `json:"ms.vss-admin-web.collection-admin-hub-contribution-data-provider"`
				Ms_vss_build_web_admin_agent_definitions_data_provider []interface{} `json:"ms.vss-build-web.admin-agent-definitions-data-provider"`
				Ms_vss_build_web_agent_clouds_list_data_provider       struct {
					TaskAgentClouds []struct {
						AcquireAgentEndpoint          string `json:"acquireAgentEndpoint"`
						AgentCloudID                  int64  `json:"agentCloudId"`
						GetAgentDefinitionEndpoint    string `json:"getAgentDefinitionEndpoint"`
						GetAgentRequestStatusEndpoint string `json:"getAgentRequestStatusEndpoint"`
						ID                            string `json:"id"`
						Internal                      bool   `json:"internal"`
						Name                          string `json:"name"`
						ReleaseAgentEndpoint          string `json:"releaseAgentEndpoint"`
						Type                          string `json:"type"`
					} `json:"taskAgentClouds"`
				} `json:"ms.vss-build-web.agent-clouds-list-data-provider"`
				Ms_vss_build_web_agent_jobs_data_provider struct {
					Jobs []struct {
						AssignTime string `json:"assignTime"`
						Data       struct {
							IsScheduledKey string `json:"IsScheduledKey"`
							ParallelismTag string `json:"ParallelismTag"`
						} `json:"data"`
						Definition struct {
							Links struct {
								Self struct {
									Href string `json:"href"`
								} `json:"self"`
								Web struct {
									Href string `json:"href"`
								} `json:"web"`
							} `json:"_links"`
							ID   int64  `json:"id"`
							Name string `json:"name"`
						} `json:"definition"`
						Demands                []string `json:"demands"`
						FinishTime             string   `json:"finishTime"`
						HostID                 string   `json:"hostId"`
						JobID                  string   `json:"jobId"`
						MatchesAllAgentsInPool bool     `json:"matchesAllAgentsInPool"`
						OrchestrationID        string   `json:"orchestrationId"`
						Owner                  struct {
							Links struct {
								Self struct {
									Href string `json:"href"`
								} `json:"self"`
								Web struct {
									Href string `json:"href"`
								} `json:"web"`
							} `json:"_links"`
							ID   int64  `json:"id"`
							Name string `json:"name"`
						} `json:"owner"`
						PlanID        string `json:"planId"`
						PlanType      string `json:"planType"`
						PoolID        int64  `json:"poolId"`
						Priority      int64  `json:"priority"`
						QueueTime     string `json:"queueTime"`
						ReceiveTime   string `json:"receiveTime"`
						RequestID     int64  `json:"requestId"`
						ReservedAgent struct {
							Links struct {
								Self struct {
									Href string `json:"href"`
								} `json:"self"`
								Web struct {
									Href string `json:"href"`
								} `json:"web"`
							} `json:"_links"`
							AccessPoint       string `json:"accessPoint"`
							Enabled           bool   `json:"enabled"`
							ID                int64  `json:"id"`
							Name              string `json:"name"`
							OsDescription     string `json:"osDescription"`
							ProvisioningState string `json:"provisioningState"`
							Status            int64  `json:"status"`
							Version           string `json:"version"`
						} `json:"reservedAgent"`
						Result       int64  `json:"result"`
						ScopeID      string `json:"scopeId"`
						ServiceOwner string `json:"serviceOwner"`
					} `json:"jobs"`
				} `json:"ms.vss-build-web.agent-jobs-data-provider"`
				Ms_vss_build_web_agent_pool_data_provider struct {
					Agents []struct {
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							Web struct {
								Href string `json:"href"`
							} `json:"web"`
						} `json:"_links"`
						AccessPoint   string `json:"accessPoint"`
						Authorization struct {
							ClientID  string `json:"clientId"`
							PublicKey struct {
								Exponent string `json:"exponent"`
								Modulus  string `json:"modulus"`
							} `json:"publicKey"`
						} `json:"authorization"`
						CreatedOn            string `json:"createdOn"`
						Enabled              bool   `json:"enabled"`
						ID                   int64  `json:"id"`
						LastCompletedRequest struct {
							AssignTime string `json:"assignTime"`
							Data       struct {
								IsScheduledKey string `json:"IsScheduledKey"`
								ParallelismTag string `json:"ParallelismTag"`
							} `json:"data"`
							Definition struct {
								Links struct {
									Self struct {
										Href string `json:"href"`
									} `json:"self"`
									Web struct {
										Href string `json:"href"`
									} `json:"web"`
								} `json:"_links"`
								ID   int64  `json:"id"`
								Name string `json:"name"`
							} `json:"definition"`
							Demands                []string `json:"demands"`
							FinishTime             string   `json:"finishTime"`
							HostID                 string   `json:"hostId"`
							JobID                  string   `json:"jobId"`
							MatchesAllAgentsInPool bool     `json:"matchesAllAgentsInPool"`
							OrchestrationID        string   `json:"orchestrationId"`
							Owner                  struct {
								Links struct {
									Self struct {
										Href string `json:"href"`
									} `json:"self"`
									Web struct {
										Href string `json:"href"`
									} `json:"web"`
								} `json:"_links"`
								ID   int64  `json:"id"`
								Name string `json:"name"`
							} `json:"owner"`
							PlanID        string `json:"planId"`
							PlanType      string `json:"planType"`
							PoolID        int64  `json:"poolId"`
							Priority      int64  `json:"priority"`
							QueueTime     string `json:"queueTime"`
							ReceiveTime   string `json:"receiveTime"`
							RequestID     int64  `json:"requestId"`
							ReservedAgent struct {
								AccessPoint       string `json:"accessPoint"`
								Enabled           bool   `json:"enabled"`
								ID                int64  `json:"id"`
								Name              string `json:"name"`
								OsDescription     string `json:"osDescription"`
								ProvisioningState string `json:"provisioningState"`
								Status            int64  `json:"status"`
								Version           string `json:"version"`
							} `json:"reservedAgent"`
							Result       int64  `json:"result"`
							ScopeID      string `json:"scopeId"`
							ServiceOwner string `json:"serviceOwner"`
						} `json:"lastCompletedRequest"`
						MaxParallelism     int64  `json:"maxParallelism"`
						Name               string `json:"name"`
						OsDescription      string `json:"osDescription"`
						ProvisioningState  string `json:"provisioningState"`
						Status             int64  `json:"status"`
						StatusChangedOn    string `json:"statusChangedOn"`
						SystemCapabilities struct {
							AgentAllowRunasroot                       string `json:"AGENT_ALLOW_RUNASROOT"`
							AzpAgentName                              string `json:"AZP_AGENT_NAME"`
							AzpPool                                   string `json:"AZP_POOL"`
							AzpURL                                    string `json:"AZP_URL"`
							Agent_ComputerName                        string `json:"Agent.ComputerName"`
							Agent_HomeDirectory                       string `json:"Agent.HomeDirectory"`
							Agent_Name                                string `json:"Agent.Name"`
							Agent_OS                                  string `json:"Agent.OS"`
							Agent_OSArchitecture                      string `json:"Agent.OSArchitecture"`
							Agent_Version                             string `json:"Agent.Version"`
							BuildqueuemonitoringPort                  string `json:"BUILDQUEUEMONITORING_PORT"`
							BUILDQUEUEMONITORINGPORT80TCP             string `json:"BUILDQUEUEMONITORING_PORT_80_TCP"`
							BUILDQUEUEMONITORINGPORT80TCPADDR         string `json:"BUILDQUEUEMONITORING_PORT_80_TCP_ADDR"`
							BUILDQUEUEMONITORINGPORT80TCPPORT         string `json:"BUILDQUEUEMONITORING_PORT_80_TCP_PORT"`
							BUILDQUEUEMONITORINGPORT80TCPPROTO        string `json:"BUILDQUEUEMONITORING_PORT_80_TCP_PROTO"`
							BuildqueuemonitoringServiceHost           string `json:"BUILDQUEUEMONITORING_SERVICE_HOST"`
							BuildqueuemonitoringServicePort           string `json:"BUILDQUEUEMONITORING_SERVICE_PORT"`
							BuildqueuemonitoringServicePortHTTP       string `json:"BUILDQUEUEMONITORING_SERVICE_PORT_HTTP"`
							CadvisorPort                              string `json:"CADVISOR_PORT"`
							CADVISORPORT8080TCP                       string `json:"CADVISOR_PORT_8080_TCP"`
							CADVISORPORT8080TCPADDR                   string `json:"CADVISOR_PORT_8080_TCP_ADDR"`
							CADVISORPORT8080TCPPORT                   string `json:"CADVISOR_PORT_8080_TCP_PORT"`
							CADVISORPORT8080TCPPROTO                  string `json:"CADVISOR_PORT_8080_TCP_PROTO"`
							CadvisorServiceHost                       string `json:"CADVISOR_SERVICE_HOST"`
							CadvisorServicePort                       string `json:"CADVISOR_SERVICE_PORT"`
							CadvisorServicePortHTTP                   string `json:"CADVISOR_SERVICE_PORT_HTTP"`
							DebianFrontend                            string `json:"DEBIAN_FRONTEND"`
							DotnetSystemGlobalizationInvariant        string `json:"DOTNET_SYSTEM_GLOBALIZATION_INVARIANT"`
							Home                                      string `json:"HOME"`
							Hostname                                  string `json:"HOSTNAME"`
							InteractiveSession                        string `json:"InteractiveSession"`
							KubernetesPort                            string `json:"KUBERNETES_PORT"`
							KUBERNETESPORT443TCP                      string `json:"KUBERNETES_PORT_443_TCP"`
							KUBERNETESPORT443TCPADDR                  string `json:"KUBERNETES_PORT_443_TCP_ADDR"`
							KUBERNETESPORT443TCPPORT                  string `json:"KUBERNETES_PORT_443_TCP_PORT"`
							KUBERNETESPORT443TCPPROTO                 string `json:"KUBERNETES_PORT_443_TCP_PROTO"`
							KubernetesServiceHost                     string `json:"KUBERNETES_SERVICE_HOST"`
							KubernetesServicePort                     string `json:"KUBERNETES_SERVICE_PORT"`
							KubernetesServicePortHTTPS                string `json:"KUBERNETES_SERVICE_PORT_HTTPS"`
							Lang                                      string `json:"LANG"`
							LcAll                                     string `json:"LC_ALL"`
							MetricsServerPort                         string `json:"METRICS_SERVER_PORT"`
							METRICSSERVERPORT443TCP                   string `json:"METRICS_SERVER_PORT_443_TCP"`
							METRICSSERVERPORT443TCPADDR               string `json:"METRICS_SERVER_PORT_443_TCP_ADDR"`
							METRICSSERVERPORT443TCPPORT               string `json:"METRICS_SERVER_PORT_443_TCP_PORT"`
							METRICSSERVERPORT443TCPPROTO              string `json:"METRICS_SERVER_PORT_443_TCP_PROTO"`
							MetricsServerServiceHost                  string `json:"METRICS_SERVER_SERVICE_HOST"`
							MetricsServerServicePort                  string `json:"METRICS_SERVER_SERVICE_PORT"`
							MetricsServerServicePortHTTPS             string `json:"METRICS_SERVER_SERVICE_PORT_HTTPS"`
							MyReleasePrometheusAdapterPort            string `json:"MY_RELEASE_PROMETHEUS_ADAPTER_PORT"`
							MYRELEASEPROMETHEUSADAPTERPORT443TCP      string `json:"MY_RELEASE_PROMETHEUS_ADAPTER_PORT_443_TCP"`
							MYRELEASEPROMETHEUSADAPTERPORT443TCPADDR  string `json:"MY_RELEASE_PROMETHEUS_ADAPTER_PORT_443_TCP_ADDR"`
							MYRELEASEPROMETHEUSADAPTERPORT443TCPPORT  string `json:"MY_RELEASE_PROMETHEUS_ADAPTER_PORT_443_TCP_PORT"`
							MYRELEASEPROMETHEUSADAPTERPORT443TCPPROTO string `json:"MY_RELEASE_PROMETHEUS_ADAPTER_PORT_443_TCP_PROTO"`
							MyReleasePrometheusAdapterServiceHost     string `json:"MY_RELEASE_PROMETHEUS_ADAPTER_SERVICE_HOST"`
							MyReleasePrometheusAdapterServicePort     string `json:"MY_RELEASE_PROMETHEUS_ADAPTER_SERVICE_PORT"`
							Oldpwd                                    string `json:"OLDPWD"`
							Path                                      string `json:"PATH"`
							PowershellDistributionChannel             string `json:"POWERSHELL_DISTRIBUTION_CHANNEL"`
							PrometheusAdapterPort                     string `json:"PROMETHEUS_ADAPTER_PORT"`
							PROMETHEUSADAPTERPORT443TCP               string `json:"PROMETHEUS_ADAPTER_PORT_443_TCP"`
							PROMETHEUSADAPTERPORT443TCPADDR           string `json:"PROMETHEUS_ADAPTER_PORT_443_TCP_ADDR"`
							PROMETHEUSADAPTERPORT443TCPPORT           string `json:"PROMETHEUS_ADAPTER_PORT_443_TCP_PORT"`
							PROMETHEUSADAPTERPORT443TCPPROTO          string `json:"PROMETHEUS_ADAPTER_PORT_443_TCP_PROTO"`
							PrometheusAdapterServiceHost              string `json:"PROMETHEUS_ADAPTER_SERVICE_HOST"`
							PrometheusAdapterServicePort              string `json:"PROMETHEUS_ADAPTER_SERVICE_PORT"`
							PrometheusPort                            string `json:"PROMETHEUS_PORT"`
							PROMETHEUSPORT9090TCP                     string `json:"PROMETHEUS_PORT_9090_TCP"`
							PROMETHEUSPORT9090TCPADDR                 string `json:"PROMETHEUS_PORT_9090_TCP_ADDR"`
							PROMETHEUSPORT9090TCPPORT                 string `json:"PROMETHEUS_PORT_9090_TCP_PORT"`
							PROMETHEUSPORT9090TCPPROTO                string `json:"PROMETHEUS_PORT_9090_TCP_PROTO"`
							PrometheusServerPort                      string `json:"PROMETHEUS_SERVER_PORT"`
							PROMETHEUSSERVERPORT80TCP                 string `json:"PROMETHEUS_SERVER_PORT_80_TCP"`
							PROMETHEUSSERVERPORT80TCPADDR             string `json:"PROMETHEUS_SERVER_PORT_80_TCP_ADDR"`
							PROMETHEUSSERVERPORT80TCPPORT             string `json:"PROMETHEUS_SERVER_PORT_80_TCP_PORT"`
							PROMETHEUSSERVERPORT80TCPPROTO            string `json:"PROMETHEUS_SERVER_PORT_80_TCP_PROTO"`
							PrometheusServerServiceHost               string `json:"PROMETHEUS_SERVER_SERVICE_HOST"`
							PrometheusServerServicePort               string `json:"PROMETHEUS_SERVER_SERVICE_PORT"`
							PrometheusServerServicePortHTTP           string `json:"PROMETHEUS_SERVER_SERVICE_PORT_HTTP"`
							PrometheusServiceHost                     string `json:"PROMETHEUS_SERVICE_HOST"`
							PrometheusServicePort                     string `json:"PROMETHEUS_SERVICE_PORT"`
							PrometheusServicePortHTTP                 string `json:"PROMETHEUS_SERVICE_PORT_HTTP"`
							PSModuleAnalysisCachePath                 string `json:"PSModuleAnalysisCachePath"`
							PsInstallFolder                           string `json:"PS_INSTALL_FOLDER"`
							Pwd                                       string `json:"PWD"`
							VsoAgentIgnore                            string `json:"VSO_AGENT_IGNORE"`
							Curl                                      string `json:"curl"`
							Git                                       string `json:"git"`
							Python3                                   string `json:"python3"`
							Sh                                        string `json:"sh"`
						} `json:"systemCapabilities"`
						Version string `json:"version"`
					} `json:"agents"`
					SelectedAgentPool struct {
						AgentCloudID  interface{} `json:"agentCloudId"`
						AutoProvision bool        `json:"autoProvision"`
						AutoSize      interface{} `json:"autoSize"`
						AutoUpdate    bool        `json:"autoUpdate"`
						CreatedBy     struct {
							Links struct {
								Avatar struct {
									Href string `json:"href"`
								} `json:"avatar"`
							} `json:"_links"`
							Descriptor  string `json:"descriptor"`
							DisplayName string `json:"displayName"`
							ID          string `json:"id"`
							ImageURL    string `json:"imageUrl"`
							UniqueName  string `json:"uniqueName"`
							URL         string `json:"url"`
						} `json:"createdBy"`
						CreatedOn string `json:"createdOn"`
						ID        int64  `json:"id"`
						IsHosted  bool   `json:"isHosted"`
						IsLegacy  bool   `json:"isLegacy"`
						Name      string `json:"name"`
						Options   int64  `json:"options"`
						Owner     struct {
							Links struct {
								Avatar struct {
									Href string `json:"href"`
								} `json:"avatar"`
							} `json:"_links"`
							Descriptor  string `json:"descriptor"`
							DisplayName string `json:"displayName"`
							ID          string `json:"id"`
							ImageURL    string `json:"imageUrl"`
							UniqueName  string `json:"uniqueName"`
							URL         string `json:"url"`
						} `json:"owner"`
						PoolType   int64       `json:"poolType"`
						Scope      string      `json:"scope"`
						Size       int64       `json:"size"`
						TargetSize interface{} `json:"targetSize"`
					} `json:"selectedAgentPool"`
					ValidTimeZones []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"validTimeZones"`
				} `json:"ms.vss-build-web.agent-pool-data-provider"`
				Ms_vss_build_web_agent_pool_security_data_provider interface{} `json:"ms.vss-build-web.agent-pool-security-data-provider"`
				Ms_vss_build_web_elastic_pool_data_provider        interface{} `json:"ms.vss-build-web.elastic-pool-data-provider"`
				Ms_vss_tfs_web_global_messages_data_provider       interface{} `json:"ms.vss-tfs-web.global-messages-data-provider"`
				Ms_vss_tfs_web_header_action_data                  struct {
					MarketplaceURL      string `json:"marketplaceUrl"`
					SettingsURL         string `json:"settingsUrl"`
					ShowProjectSelector bool   `json:"showProjectSelector"`
					SuiteHomeURL        string `json:"suiteHomeUrl"`
				} `json:"ms.vss-tfs-web.header-action-data"`
				Ms_vss_tfs_web_navigation_context_data_provider struct {
					RightMenu struct {
						AdminSettings struct {
							Available bool `json:"available"`
						} `json:"adminSettings"`
						RightMenuBar struct {
							Actions struct {
								Alerts struct {
									Icon       string `json:"icon"`
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"alerts"`
								Community struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"community"`
								Divider struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
								} `json:"divider"`
								Extensions struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"extensions"`
								Gallery struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"gallery"`
								GettingStarted struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"gettingStarted"`
								Help struct {
									Icon       string `json:"icon"`
									ID         string `json:"id"`
									SubMenuID  string `json:"subMenuId"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									Title      string `json:"title"`
								} `json:"help"`
								ManageFeatures struct {
									Command struct {
										Dependencies []string `json:"dependencies"`
										MethodName   string   `json:"methodName"`
										ServiceName  string   `json:"serviceName"`
									} `json:"command"`
									CommandID  string `json:"commandId"`
									Icon       string `json:"icon"`
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									Title      string `json:"title"`
								} `json:"manageFeatures"`
								Msdn struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"msdn"`
								Privacy struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"privacy"`
								Profile struct {
									Icon       string `json:"icon"`
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									Title      string `json:"title"`
									URL        string `json:"url"`
								} `json:"profile"`
								Security struct {
									Icon       string `json:"icon"`
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"security"`
								SignOut struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"signOut"`
								Support struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"support"`
								Themes struct {
									Command struct {
										Dependencies []string `json:"dependencies"`
										MethodName   string   `json:"methodName"`
										ServiceName  string   `json:"serviceName"`
									} `json:"command"`
									FabricIconName string `json:"fabricIconName"`
									ID             string `json:"id"`
									TargetSelf     bool   `json:"targetSelf"`
									Text           string `json:"text"`
									Title          string `json:"title"`
								} `json:"themes"`
								Usage struct {
									Icon       string `json:"icon"`
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"usage"`
								User struct {
									ID         string `json:"id"`
									ItemType   int64  `json:"itemType"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									Title      string `json:"title"`
									URL        string `json:"url"`
								} `json:"user"`
								UserDivider struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
								} `json:"userDivider"`
								Welcome struct {
									ID         string `json:"id"`
									TargetSelf bool   `json:"targetSelf"`
									Text       string `json:"text"`
									URL        string `json:"url"`
								} `json:"welcome"`
							} `json:"actions"`
							Available  bool `json:"available"`
							Properties struct {
								IdentityImageURI string `json:"IdentityImageUri"`
							} `json:"properties"`
						} `json:"rightMenuBar"`
						Search struct {
							Available  bool `json:"available"`
							Properties struct {
								CodeSearchAvailable                         bool `json:"codeSearchAvailable"`
								ShouldInvokeConditionalFaultInForWIS        bool `json:"shouldInvokeConditionalFaultInForWIS"`
								ShouldInvokeConditionalFaultInForWikiSearch bool `json:"shouldInvokeConditionalFaultInForWikiSearch"`
								WikiSearchAvailable                         bool `json:"wikiSearchAvailable"`
								WorkItemSearchAvailable                     bool `json:"workItemSearchAvailable"`
							} `json:"properties"`
						} `json:"search"`
					} `json:"rightMenu"`
				} `json:"ms.vss-tfs-web.navigation-context-data-provider"`
				Ms_vss_tfs_web_navigation_settings_data_provider struct {
					ResolvedURL string `json:"resolvedUrl"`
					User        struct {
						FirstRunDismissed     bool   `json:"firstRunDismissed"`
						Myworkdefaultpivot    string `json:"myworkdefaultpivot"`
						NavHeaderBannerHidden bool   `json:"navHeaderBannerHidden"`
						VerticalMenuExpanded  int64  `json:"verticalMenuExpanded"`
					} `json:"user"`
				} `json:"ms.vss-tfs-web.navigation-settings-data-provider"`
				Ms_vss_tfs_web_page_data                                    struct{} `json:"ms.vss-tfs-web.page-data"`
				Ms_vss_tfs_web_tfs_rest_error_handler_service_data_provider struct {
					ClosePopupURL string `json:"closePopupUrl"`
					OpenPopupURL  string `json:"openPopupUrl"`
				} `json:"ms.vss-tfs-web.tfs-rest-error-handler-service-data-provider"`
				Ms_vss_tfs_web_vertical_header_me_control_data_provider struct {
					AadIdpConfig struct {
						AadIdpForgetURL             string `json:"aadIdpForgetUrl"`
						AadIdpRememberedAccountsURL string `json:"aadIdpRememberedAccountsUrl"`
						AadIdpSignOutAndForgetURL   string `json:"aadIdpSignOutAndForgetUrl"`
						AadIdpSignOutURL            string `json:"aadIdpSignOutUrl"`
					} `json:"aadIdpConfig"`
					AuthProviderConfig struct {
						SignInToURL string `json:"signInToUrl"`
						SignInURL   string `json:"signInUrl"`
						SignOutURL  string `json:"signOutUrl"`
						SwitchToURL string `json:"switchToUrl"`
						SwitchURL   string `json:"switchUrl"`
					} `json:"authProviderConfig"`
					CurrentAccount struct {
						AadTenantDisplayName string `json:"aadTenantDisplayName"`
						Email                string `json:"email"`
						ID                   string `json:"id"`
						IsMsaOrFpmsa         bool   `json:"isMsaOrFpmsa"`
						Name                 string `json:"name"`
					} `json:"currentAccount"`
					MeControlScriptURL string `json:"meControlScriptUrl"`
					ProfileURL         string `json:"profileUrl"`
				} `json:"ms.vss-tfs-web.vertical-header-me-control-data-provider"`
				Ms_vss_web_component_data struct {
					Content []struct {
						ComponentID         string `json:"componentId"`
						ComponentProperties struct {
							DefaultPivot                string `json:"defaultPivot"`
							DisplayTabIcons             bool   `json:"displayTabIcons"`
							HubContributionDataProvider string `json:"hubContributionDataProvider"`
							ID                          string `json:"id"`
							Key                         string `json:"key"`
							Level                       string `json:"level"`
						} `json:"componentProperties"`
						ComponentType string `json:"componentType"`
					} `json:"content"`
					Content_header []struct {
						ComponentID         string `json:"componentId"`
						ComponentProperties struct {
							ComponentOrder int64 `json:"componentOrder"`
						} `json:"componentProperties"`
						ComponentType string `json:"componentType"`
					} `json:"content-header"`
					Header []struct {
						ComponentID         string `json:"componentId"`
						ComponentProperties struct {
							AdjustmentIgnore bool   `json:"adjustmentIgnore"`
							AlwaysExpanded   bool   `json:"alwaysExpanded"`
							ClassName        string `json:"className"`
							ComponentOrder   int64  `json:"componentOrder"`
							Components       []struct {
								ComponentProperties struct {
									Dependencies []string `json:"dependencies"`
									Key          string   `json:"key"`
								} `json:"componentProperties"`
								ComponentType string `json:"componentType"`
							} `json:"components"`
							Key         string `json:"key"`
							Orientation string `json:"orientation"`
							RouteID     string `json:"routeId"`
						} `json:"componentProperties"`
						ComponentType string `json:"componentType"`
					} `json:"header"`
					Page []struct {
						ComponentID         string `json:"componentId"`
						ComponentProperties struct {
							ClassName  string `json:"className"`
							Components []struct {
								ComponentProperties struct {
									ClassName string `json:"className"`
									Role      string `json:"role"`
								} `json:"componentProperties"`
								ComponentRegion    string `json:"componentRegion"`
								ComponentType      string `json:"componentType"`
								ErrorBoundaryProps struct {
									HandlerComponentType string `json:"handlerComponentType"`
								} `json:"errorBoundaryProps"`
							} `json:"components"`
							Orientation string `json:"orientation"`
						} `json:"componentProperties"`
						ComponentType string `json:"componentType"`
					} `json:"page"`
				} `json:"ms.vss-web.component-data"`
				Ms_vss_web_device_type_data struct {
					IsMobileDevice bool `json:"isMobileDevice"`
					IsTabletDevice bool `json:"isTabletDevice"`
				} `json:"ms.vss-web.device-type-data"`
				Ms_vss_web_navigation_data struct {
					Displayed []string `json:"displayed"`
					Hierarchy struct {
						Ms_vss_code_web_collection_code_hub_group     []string `json:"ms.vss-code-web.collection-code-hub-group"`
						Ms_vss_tfs_web_collection_users_hub_group     []string `json:"ms.vss-tfs-web.collection-users-hub-group"`
						Ms_vss_tfs_web_suite_me_page_hub_group        []string `json:"ms.vss-tfs-web.suite-me-page-hub-group"`
						Ms_vss_web_account_admin_hub_group            []string `json:"ms.vss-web.account-admin-hub-group"`
						Ms_vss_web_collection_admin_hub_group         []string `json:"ms.vss-web.collection-admin-hub-group"`
						Ms_vss_web_collection_hub_groups_collection   []string `json:"ms.vss-web.collection-hub-groups-collection"`
						Ms_vss_web_collection_page                    []string `json:"ms.vss-web.collection-page"`
						Ms_vss_web_collection_user_settings_hub_group []string `json:"ms.vss-web.collection-user-settings-hub-group"`
					} `json:"hierarchy"`
					Location struct {
						Ms_vss_admin_web_admin_collection_projects_hub          string      `json:"ms.vss-admin-web.admin-collection-projects-hub"`
						Ms_vss_admin_web_admin_collection_security_hub          string      `json:"ms.vss-admin-web.admin-collection-security-hub"`
						Ms_vss_admin_web_collection_admin_hub                   string      `json:"ms.vss-admin-web.collection-admin-hub"`
						Ms_vss_admin_web_user_admin_hub                         string      `json:"ms.vss-admin-web.user-admin-hub"`
						Ms_vss_build_web_agent_pools_hub                        string      `json:"ms.vss-build-web.agent-pools-hub"`
						Ms_vss_build_web_build_release_collection_admin_hub_new string      `json:"ms.vss-build-web.build-release-collection-admin-hub-new"`
						Ms_vss_code_web_collection_changesets_hub               string      `json:"ms.vss-code-web.collection-changesets-hub"`
						Ms_vss_code_web_collection_code_hub_group               interface{} `json:"ms.vss-code-web.collection-code-hub-group"`
						Ms_vss_code_web_collection_files_hub_tfvc               string      `json:"ms.vss-code-web.collection-files-hub-tfvc"`
						Ms_vss_code_web_collection_shelvesets_hub               string      `json:"ms.vss-code-web.collection-shelvesets-hub"`
						Ms_vss_distributed_task_oauth_configurations_hub        string      `json:"ms.vss-distributed-task.oauth-configurations-hub"`
						Ms_vss_notifications_web_collection_admin_hub           string      `json:"ms.vss-notifications-web.collection-admin-hub"`
						Ms_vss_notifications_web_user_admin_hub                 string      `json:"ms.vss-notifications-web.user-admin-hub"`
						Ms_vss_org_web_collection_admin_policy_hub              string      `json:"ms.vss-org-web.collection-admin-policy-hub"`
						Ms_vss_releaseManagement_web_hub_deploymentpools        string      `json:"ms.vss-releaseManagement-web.hub-deploymentpools"`
						Ms_vss_search_web_collection_hub                        string      `json:"ms.vss-search-web.collection-hub"`
						Ms_vss_search_web_collection_hub_group                  interface{} `json:"ms.vss-search-web.collection-hub-group"`
						Ms_vss_tfs_web_about_hub                                string      `json:"ms.vss-tfs-web.about-hub"`
						Ms_vss_tfs_web_collection_users_hub                     string      `json:"ms.vss-tfs-web.collection-users-hub"`
						Ms_vss_tfs_web_collection_users_hub_group               interface{} `json:"ms.vss-tfs-web.collection-users-hub-group"`
						Ms_vss_tfs_web_diagnostics_hub                          string      `json:"ms.vss-tfs-web.diagnostics-hub"`
						Ms_vss_tfs_web_suite_me_page_hub                        string      `json:"ms.vss-tfs-web.suite-me-page-hub"`
						Ms_vss_tfs_web_suite_me_page_hub_group                  interface{} `json:"ms.vss-tfs-web.suite-me-page-hub-group"`
						Ms_vss_tfs_web_user_settings_security_hub               string      `json:"ms.vss-tfs-web.user-settings-security-hub"`
						Ms_vss_tfs_web_utilization_usagesummary_hub             string      `json:"ms.vss-tfs-web.utilization-usagesummary-hub"`
						Ms_vss_tfs_web_utilization_userusagesummary_hub         string      `json:"ms.vss-tfs-web.utilization-userusagesummary-hub"`
						Ms_vss_web_account_admin_hub_group                      string      `json:"ms.vss-web.account-admin-hub-group"`
						Ms_vss_web_collection_admin_hub_group                   interface{} `json:"ms.vss-web.collection-admin-hub-group"`
						Ms_vss_web_collection_home_hub_group                    interface{} `json:"ms.vss-web.collection-home-hub-group"`
						Ms_vss_web_collection_hub_groups_collection             interface{} `json:"ms.vss-web.collection-hub-groups-collection"`
						Ms_vss_web_collection_page                              interface{} `json:"ms.vss-web.collection-page"`
						Ms_vss_web_collection_user_settings_hub_group           interface{} `json:"ms.vss-web.collection-user-settings-hub-group"`
						Ms_vss_work_web_work_customization_hub                  string      `json:"ms.vss-work-web.work-customization-hub"`
					} `json:"location"`
					Navigation struct {
						Ms_vss_admin_web_admin_collection_projects_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-admin-web.admin-collection-projects-hub"`
						Ms_vss_admin_web_admin_collection_security_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-admin-web.admin-collection-security-hub"`
						Ms_vss_admin_web_collection_admin_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-admin-web.collection-admin-hub"`
						Ms_vss_admin_web_user_admin_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-admin-web.user-admin-hub"`
						Ms_vss_build_web_agent_pools_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-build-web.agent-pools-hub"`
						Ms_vss_build_web_build_release_collection_admin_hub_new struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-build-web.build-release-collection-admin-hub-new"`
						Ms_vss_code_web_collection_changesets_hub struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-code-web.collection-changesets-hub"`
						Ms_vss_code_web_collection_code_hub_group struct {
							DisplayedDocTypes []string `json:"displayedDocTypes"`
							Flags             int64    `json:"flags"`
							ID                string   `json:"id"`
							Name              string   `json:"name"`
							Order             int64    `json:"order"`
						} `json:"ms.vss-code-web.collection-code-hub-group"`
						Ms_vss_code_web_collection_files_hub_tfvc struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-code-web.collection-files-hub-tfvc"`
						Ms_vss_code_web_collection_shelvesets_hub struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-code-web.collection-shelvesets-hub"`
						Ms_vss_distributed_task_oauth_configurations_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-distributed-task.oauth-configurations-hub"`
						Ms_vss_notifications_web_collection_admin_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-notifications-web.collection-admin-hub"`
						Ms_vss_notifications_web_user_admin_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-notifications-web.user-admin-hub"`
						Ms_vss_org_web_collection_admin_policy_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-org-web.collection-admin-policy-hub"`
						Ms_vss_releaseManagement_web_hub_deploymentpools struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-releaseManagement-web.hub-deploymentpools"`
						Ms_vss_search_web_collection_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-search-web.collection-hub"`
						Ms_vss_search_web_collection_hub_group struct {
							DisplayedDocTypes []string `json:"displayedDocTypes"`
							Flags             int64    `json:"flags"`
							ID                string   `json:"id"`
							Name              string   `json:"name"`
							Order             int64    `json:"order"`
						} `json:"ms.vss-search-web.collection-hub-group"`
						Ms_vss_tfs_web_about_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-tfs-web.about-hub"`
						Ms_vss_tfs_web_collection_users_hub struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-tfs-web.collection-users-hub"`
						Ms_vss_tfs_web_collection_users_hub_group struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-tfs-web.collection-users-hub-group"`
						Ms_vss_tfs_web_diagnostics_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-tfs-web.diagnostics-hub"`
						Ms_vss_tfs_web_suite_me_page_hub struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-tfs-web.suite-me-page-hub"`
						Ms_vss_tfs_web_suite_me_page_hub_group struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-tfs-web.suite-me-page-hub-group"`
						Ms_vss_tfs_web_user_settings_security_hub struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-tfs-web.user-settings-security-hub"`
						Ms_vss_tfs_web_utilization_usagesummary_hub struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-tfs-web.utilization-usagesummary-hub"`
						Ms_vss_tfs_web_utilization_userusagesummary_hub struct {
							Flags    int64  `json:"flags"`
							IconName string `json:"iconName"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-tfs-web.utilization-userusagesummary-hub"`
						Ms_vss_web_account_admin_hub_group struct {
							Flags    int64  `json:"flags"`
							Icon     string `json:"icon"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							NavGroup string `json:"navGroup"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-web.account-admin-hub-group"`
						Ms_vss_web_collection_admin_hub_group struct {
							Flags    int64  `json:"flags"`
							Icon     string `json:"icon"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							NavGroup string `json:"navGroup"`
							Order    int64  `json:"order"`
						} `json:"ms.vss-web.collection-admin-hub-group"`
						Ms_vss_web_collection_home_hub_group struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-web.collection-home-hub-group"`
						Ms_vss_web_collection_hub_groups_collection struct {
							Flags int64       `json:"flags"`
							ID    string      `json:"id"`
							Name  interface{} `json:"name"`
							Order int64       `json:"order"`
						} `json:"ms.vss-web.collection-hub-groups-collection"`
						Ms_vss_web_collection_page struct {
							Flags int64       `json:"flags"`
							ID    string      `json:"id"`
							Name  interface{} `json:"name"`
							Order int64       `json:"order"`
						} `json:"ms.vss-web.collection-page"`
						Ms_vss_web_collection_user_settings_hub_group struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-web.collection-user-settings-hub-group"`
						Ms_vss_work_web_work_customization_hub struct {
							Flags int64  `json:"flags"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-work-web.work-customization-hub"`
					} `json:"navigation"`
					PageTitle      string   `json:"pageTitle"`
					ResolvedURL    string   `json:"resolvedUrl"`
					RouteID        string   `json:"routeId"`
					RouteTemplates []string `json:"routeTemplates"`
					RouteValues    struct {
						Action      string `json:"action"`
						AdminPivot  string `json:"adminPivot"`
						Controller  string `json:"controller"`
						ServiceHost string `json:"serviceHost"`
					} `json:"routeValues"`
				} `json:"ms.vss-web.navigation-data"`
				Ms_vss_web_page_data struct {
					EsNextSupported               bool   `json:"esNextSupported"`
					HostID                        string `json:"hostId"`
					HostName                      string `json:"hostName"`
					HostPath                      string `json:"hostPath"`
					HostType                      int64  `json:"hostType"`
					IsBundlingEnabled             bool   `json:"isBundlingEnabled"`
					IsCdnAvailable                bool   `json:"isCdnAvailable"`
					IsCdnEnabled                  bool   `json:"isCdnEnabled"`
					IsDebugMode                   bool   `json:"isDebugMode"`
					IsDevfabric                   bool   `json:"isDevfabric"`
					IsHosted                      bool   `json:"isHosted"`
					IsProduction                  bool   `json:"isProduction"`
					IsSslOnly                     bool   `json:"isSslOnly"`
					IsTracePointCollectionEnabled bool   `json:"isTracePointCollectionEnabled"`
					ScriptType                    string `json:"scriptType"`
					ServiceInstanceType           string `json:"serviceInstanceType"`
					ServiceVersion                string `json:"serviceVersion"`
					SessionID                     string `json:"sessionId"`
					StyleType                     string `json:"styleType"`
					User                          struct {
						Links struct {
							Avatar struct {
								Href string `json:"href"`
							} `json:"avatar"`
						} `json:"_links"`
						Descriptor  string `json:"descriptor"`
						DisplayName string `json:"displayName"`
						ID          string `json:"id"`
						ImageURL    string `json:"imageUrl"`
						Name        string `json:"name"`
						UniqueName  string `json:"uniqueName"`
						URL         string `json:"url"`
					} `json:"user"`
					WebPlatformVersion int64 `json:"webPlatformVersion"`
				} `json:"ms.vss-web.page-data"`
				Ms_vss_web_shared_data interface{} `json:"ms.vss-web.shared-data"`
				Ms_vss_web_theme_data  struct {
					RequestedThemeID string `json:"requestedThemeId"`
				} `json:"ms.vss-web.theme-data"`
				Ms_vss_web_user_claims_data struct {
					Member                      bool `json:"member"`
					Ms_azure_artifacts_feature  bool `json:"ms.azure-artifacts.feature"`
					Ms_feed_feed                bool `json:"ms.feed.feed"`
					Ms_vss_build_pipelines      bool `json:"ms.vss-build.pipelines"`
					Ms_vss_code_version_control bool `json:"ms.vss-code.version-control"`
					Ms_vss_test_web_test        bool `json:"ms.vss-test-web.test"`
					Ms_vss_work_agile           bool `json:"ms.vss-work.agile"`
				} `json:"ms.vss-web.user-claims-data"`
			} `json:"data"`
			Exceptions        interface{} `json:"exceptions"`
			ResolvedProviders []struct {
				ID string `json:"id"`
			} `json:"resolvedProviders"`
			ScopeName  interface{} `json:"scopeName"`
			ScopeValue interface{} `json:"scopeValue"`
			SharedData struct {
				Assets struct {
					Ms_vss_build_web_admin_tabs_zerodata_png                                    string `json:"ms.vss-build-web/admin.tabs/zerodata.png"`
					Ms_vss_storage_web_artifacts_usage_content_Images_apache_svg                string `json:"ms.vss-storage-web/artifacts-usage-content/Images/apache.svg"`
					Ms_vss_storage_web_artifacts_usage_content_Images_npm_svg                   string `json:"ms.vss-storage-web/artifacts-usage-content/Images/npm.svg"`
					Ms_vss_storage_web_artifacts_usage_content_Images_nuget_svg                 string `json:"ms.vss-storage-web/artifacts-usage-content/Images/nuget.svg"`
					Ms_vss_storage_web_artifacts_usage_content_Images_python_svg                string `json:"ms.vss-storage-web/artifacts-usage-content/Images/python.svg"`
					Ms_vss_storage_web_artifacts_usage_content_Images_universal_package_svg     string `json:"ms.vss-storage-web/artifacts-usage-content/Images/universal-package.svg"`
					Ms_vss_tfs_web_platform_content_Illustrations_general_no_results_found_svg  string `json:"ms.vss-tfs-web/platform-content/Illustrations/general-no-results-found.svg"`
					Ms_vss_work_web_my_work_content_Content_zerodata_your_work_in_one_place_svg string `json:"ms.vss-work-web/my-work-content/Content/zerodata-your-work-in-one-place.svg"`
				} `json:"_assets"`
				Contributions struct {
					Contributions struct {
						Ms_feed_package_collection_search_provider struct {
							Attributes          int64  `json:"::Attributes"`
							ServiceInstanceType string `json:"::ServiceInstanceType"`
							Version             string `json:"::Version"`
							SharedData          struct {
								FeatureFlags []string `json:"featureFlags"`
							} `json:"_sharedData"`
							Dependencies      []string `json:"dependencies"`
							Placeholder       string   `json:"placeholder"`
							Priority          int64    `json:"priority"`
							ProviderName      string   `json:"providerName"`
							ServiceName       string   `json:"serviceName"`
							SupportedDocTypes []string `json:"supportedDocTypes"`
						} `json:"ms.feed.package-collection-search-provider"`
						Ms_vss_admin_web_admin_hub_header_breadcrumb_service struct {
							Attributes   int64    `json:"::Attributes"`
							Hashcode     string   `json:"::Hashcode"`
							BreadcrumbID string   `json:"breadcrumbId"`
							Dependencies []string `json:"dependencies"`
							ServiceName  string   `json:"serviceName"`
						} `json:"ms.vss-admin-web.admin-hub-header-breadcrumb-service"`
						Ms_vss_admin_web_collection_admin_hub_tab_group struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Name       string `json:"name"`
							Rank       int64  `json:"rank"`
						} `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
						Ms_vss_admin_web_collection_admin_security_tab_group struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Name       string `json:"name"`
							Rank       int64  `json:"rank"`
						} `json:"ms.vss-admin-web.collection-admin-security-tab-group"`
						Ms_vss_admin_web_create_hidden_org_admin_hubgroup struct {
							Attributes   int64  `json:"::Attributes"`
							Hashcode     string `json:"::Hashcode"`
							BreadcrumbID string `json:"breadcrumbId"`
							Hidden       bool   `json:"hidden"`
							Key          string `json:"key"`
							Priority     int64  `json:"priority"`
							Rank         int64  `json:"rank"`
							Text         string `json:"text"`
						} `json:"ms.vss-admin-web.create-hidden-org-admin-hubgroup"`
						Ms_vss_admin_web_org_admin_orggroups_tab struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								Assets    []string `json:"assets"`
								Features  []string `json:"features"`
								Locations []struct {
									ServiceInstanceID string `json:"ServiceInstanceId"`
									HostType          string `json:"hostType"`
								} `json:"locations"`
							} `json:"_sharedData"`
							ComponentType string `json:"componentType"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-admin-web.org-admin-orggroups-tab"`
						Ms_vss_admin_web_organization_admin_aad_tab struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
							} `json:"_sharedData"`
							ComponentType string `json:"componentType"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-admin-web.organization-admin-aad-tab"`
						Ms_vss_admin_web_organization_admin_new_overview_tab struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
							} `json:"_sharedData"`
							ComponentType string `json:"componentType"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-admin-web.organization-admin-new-overview-tab"`
						Ms_vss_admin_web_organization_admin_new_policies_tab struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								Assets       []string `json:"assets"`
								FeatureFlags []string `json:"featureFlags"`
								Routes       []string `json:"routes"`
							} `json:"_sharedData"`
							ComponentType string `json:"componentType"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-admin-web.organization-admin-new-policies-tab"`
						Ms_vss_admin_web_organization_admin_new_projects_tab struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								Assets       []string `json:"assets"`
								FeatureFlags []string `json:"featureFlags"`
								Routes       []string `json:"routes"`
							} `json:"_sharedData"`
							ComponentType string `json:"componentType"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-admin-web.organization-admin-new-projects-tab"`
						Ms_vss_aex_user_management_collection_user_management_pivot struct {
							Attributes int64  `json:"::Attributes"`
							Version    string `json:"::Version"`
							SharedData struct {
								Routes []string `json:"routes"`
							} `json:"_sharedData"`
							ComponentProperties struct {
								TelemetryArea    string `json:"telemetryArea"`
								TelemetryFeature string `json:"telemetryFeature"`
							} `json:"componentProperties"`
							ComponentRegion string   `json:"componentRegion"`
							ComponentType   string   `json:"componentType"`
							Dependencies    []string `json:"dependencies"`
							IconProps       struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-aex-user-management.collection-user-management-pivot"`
						Ms_vss_audit_web_collection_admin_auditing_hub struct {
							Attributes    int64    `json:"::Attributes"`
							Hashcode      string   `json:"::Hashcode"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-audit-web.collection-admin-auditing-hub"`
						Ms_vss_bill_web_user_admin_hub_bill_tab_new struct {
							Attributes int64  `json:"::Attributes"`
							Version    string `json:"::Version"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
							} `json:"_sharedData"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-bill-web.user-admin-hub-bill-tab-new"`
						Ms_vss_build_web_account_admin_agent_pools_tab struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								Assets       []string `json:"assets"`
								FeatureFlags []string `json:"featureFlags"`
								Locations    []struct {
									HostType          string `json:"HostType"`
									ServiceInstanceID string `json:"ServiceInstanceId"`
								} `json:"locations"`
								Routes []string `json:"routes"`
							} `json:"_sharedData"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-build-web.account-admin-agent-pools-tab"`
						Ms_vss_build_web_build_release_account_settings_tab_group struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Name       string `json:"name"`
							Rank       int64  `json:"rank"`
						} `json:"ms.vss-build-web.build-release-account-settings-tab-group"`
						Ms_vss_build_web_build_release_queue_admin_parallel_jobs__org struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
							} `json:"_sharedData"`
							IconProps struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey       string `json:"itemKey"`
							ModuleContent string `json:"moduleContent"`
							Name          string `json:"name"`
							Order         int64  `json:"order"`
						} `json:"ms.vss-build-web.build-release-queue-admin-parallel-jobs--org"`
						Ms_vss_build_web_pipeline_org_settings_tab struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
							} `json:"_sharedData"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							GroupKey      string   `json:"groupKey"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-build-web.pipeline-org-settings-tab"`
						Ms_vss_build_web_pools_admin_header_breadcrumb_service struct {
							Attributes   int64    `json:"::Attributes"`
							Hashcode     string   `json:"::Hashcode"`
							BreadcrumbID string   `json:"breadcrumbId"`
							Dependencies []string `json:"dependencies"`
							ServiceName  string   `json:"serviceName"`
						} `json:"ms.vss-build-web.pools-admin-header-breadcrumb-service"`
						Ms_vss_code_web_admin_account_version_control_pivot struct {
							Attributes    int64    `json:"::Attributes"`
							Hashcode      string   `json:"::Hashcode"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-code-web.admin-account-version-control-pivot"`
						Ms_vss_code_web_code_account_settings_tab_group struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Name       string `json:"name"`
							Rank       int64  `json:"rank"`
						} `json:"ms.vss-code-web.code-account-settings-tab-group"`
						Ms_vss_code_web_my_prs_flyout_tab struct {
							Attributes    int64    `json:"::Attributes"`
							Hashcode      string   `json:"::Hashcode"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							ID            string   `json:"id"`
							Key           string   `json:"key"`
							Name          string   `json:"name"`
							Order         int64    `json:"order"`
						} `json:"ms.vss-code-web.my-prs-flyout-tab"`
						Ms_vss_code_web_repository_dropdown_menu struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
						} `json:"ms.vss-code-web.repository-dropdown-menu"`
						Ms_vss_code_web_versioncontrol_collection_search_provider struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
								Features     []string `json:"features"`
							} `json:"_sharedData"`
							Dependencies      []string `json:"dependencies"`
							Placeholder       string   `json:"placeholder"`
							Priority          int64    `json:"priority"`
							ProviderName      string   `json:"providerName"`
							ServiceName       string   `json:"serviceName"`
							SupportedDocTypes []string `json:"supportedDocTypes"`
						} `json:"ms.vss-code-web.versioncontrol-collection-search-provider"`
						Ms_vss_distributed_task_oauth_configurations_pivot struct {
							Attributes int64  `json:"::Attributes"`
							Version    string `json:"::Version"`
							IconProps  struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey       string `json:"itemKey"`
							ModuleContent string `json:"moduleContent"`
							Name          string `json:"name"`
							Order         int64  `json:"order"`
						} `json:"ms.vss-distributed-task.oauth-configurations-pivot"`
						Ms_vss_extmgmt_web_manageExtensions_pivot struct {
							Attributes    int64    `json:"::Attributes"`
							Version       string   `json:"::Version"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ID    string `json:"id"`
							Name  string `json:"name"`
							Order int64  `json:"order"`
						} `json:"ms.vss-extmgmt-web.manageExtensions-pivot"`
						Ms_vss_favorites_vertical_header_favorites struct {
							Attributes    int64    `json:"::Attributes"`
							Hashcode      string   `json:"::Hashcode"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							ID            string   `json:"id"`
							Key           string   `json:"key"`
							Name          string   `json:"name"`
							Order         int64    `json:"order"`
						} `json:"ms.vss-favorites.vertical-header-favorites"`
						Ms_vss_notifications_web_collection_admin_pivot struct {
							Attributes  int64  `json:"::Attributes"`
							Hashcode    string `json:"::Hashcode"`
							FpsRequired bool   `json:"fpsRequired"`
							IconProps   struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey       string `json:"itemKey"`
							ModuleContent string `json:"moduleContent"`
							Name          string `json:"name"`
							Order         int64  `json:"order"`
						} `json:"ms.vss-notifications-web.collection-admin-pivot"`
						Ms_vss_releaseManagement_web_deployment_pools_pivot struct {
							Attributes   int64    `json:"::Attributes"`
							Version      string   `json:"::Version"`
							Dependencies []string `json:"dependencies"`
							IconProps    struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey       string `json:"itemKey"`
							ModuleContent string `json:"moduleContent"`
							Name          string `json:"name"`
							Order         int64  `json:"order"`
						} `json:"ms.vss-releaseManagement-web.deployment-pools-pivot"`
						Ms_vss_search_web_global_search_coordinator struct {
							Attributes          int64  `json:"::Attributes"`
							Hashcode            string `json:"::Hashcode"`
							ServiceInstanceType string `json:"::ServiceInstanceType"`
							SharedData          struct {
								Locations []struct {
									HostType          string `json:"HostType"`
									ServiceInstanceID string `json:"ServiceInstanceId"`
								} `json:"locations"`
								Routes []string `json:"routes"`
							} `json:"_sharedData"`
							Priority    int64  `json:"priority"`
							ServiceName string `json:"serviceName"`
						} `json:"ms.vss-search-web.global-search-coordinator"`
						Ms_vss_storage_web_collection_admin_artifacts_tab_group struct {
							Attributes int64  `json:"::Attributes"`
							Version    string `json:"::Version"`
							Name       string `json:"name"`
							Rank       int64  `json:"rank"`
						} `json:"ms.vss-storage-web.collection-admin-artifacts-tab-group"`
						Ms_vss_storage_web_org_admin_artifacts_usage_tab struct {
							Attributes int64  `json:"::Attributes"`
							Version    string `json:"::Version"`
							SharedData struct {
								Assets    []string `json:"assets"`
								Locations []struct {
									HostType          string `json:"hostType"`
									ServiceInstanceID string `json:"serviceInstanceId"`
								} `json:"locations"`
								Routes []string `json:"routes"`
							} `json:"_sharedData"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							IconProps     struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey string `json:"itemKey"`
							Name    string `json:"name"`
							Order   int64  `json:"order"`
						} `json:"ms.vss-storage-web.org-admin-artifacts-usage-tab"`
						Ms_vss_tfs_web_code_conduct struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							GroupKey   string `json:"groupKey"`
							Href       string `json:"href"`
							IconProps  struct {
								IconName string `json:"iconName"`
							} `json:"iconProps"`
							ID   string `json:"id"`
							Name string `json:"name"`
							Rank int64  `json:"rank"`
						} `json:"ms.vss-tfs-web.code-conduct"`
						Ms_vss_tfs_web_get_started_action struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							GroupKey   string `json:"groupKey"`
							Href       string `json:"href"`
							IconProps  struct {
								IconName string `json:"iconName"`
							} `json:"iconProps"`
							ID   string `json:"id"`
							Name string `json:"name"`
							Rank int64  `json:"rank"`
						} `json:"ms.vss-tfs-web.get-started-action"`
						Ms_vss_tfs_web_get_support struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Command    struct {
								Dependencies []string `json:"dependencies"`
								MethodName   string   `json:"methodName"`
								ServiceName  string   `json:"serviceName"`
							} `json:"command"`
							GroupKey  string `json:"groupKey"`
							IconProps struct {
								IconName string `json:"iconName"`
							} `json:"iconProps"`
							ID   string `json:"id"`
							Name string `json:"name"`
							Rank int64  `json:"rank"`
						} `json:"ms.vss-tfs-web.get-support"`
						Ms_vss_tfs_web_goto_home_shortcut struct {
							Attributes int64    `json:"::Attributes"`
							Hashcode   string   `json:"::Hashcode"`
							Combos     []string `json:"combos"`
							Command    struct {
								Dependencies []string `json:"dependencies"`
								MethodName   string   `json:"methodName"`
								ServiceName  string   `json:"serviceName"`
							} `json:"command"`
							Description              string `json:"description"`
							Group                    string `json:"group"`
							IsPageNavigationShortcut bool   `json:"isPageNavigationShortcut"`
						} `json:"ms.vss-tfs-web.goto-home-shortcut"`
						Ms_vss_tfs_web_goto_organization_admin_shortcut struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								Routes []string `json:"routes"`
							} `json:"_sharedData"`
							Combos  []string `json:"combos"`
							Command struct {
								Dependencies    []string      `json:"dependencies"`
								MethodArguments []interface{} `json:"methodArguments"`
								MethodName      string        `json:"methodName"`
								ServiceName     string        `json:"serviceName"`
							} `json:"command"`
							Description              string `json:"description"`
							Group                    string `json:"group"`
							IsPageNavigationShortcut bool   `json:"isPageNavigationShortcut"`
						} `json:"ms.vss-tfs-web.goto-organization-admin-shortcut"`
						Ms_vss_tfs_web_goto_projects_shortcut struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								Routes []string `json:"routes"`
							} `json:"_sharedData"`
							Combos  []string `json:"combos"`
							Command struct {
								Dependencies    []string `json:"dependencies"`
								MethodArguments []string `json:"methodArguments"`
								MethodName      string   `json:"methodName"`
								ServiceName     string   `json:"serviceName"`
							} `json:"command"`
							Description              string `json:"description"`
							Group                    string `json:"group"`
							IsPageNavigationShortcut bool   `json:"isPageNavigationShortcut"`
						} `json:"ms.vss-tfs-web.goto-projects-shortcut"`
						Ms_vss_tfs_web_header_breadcrumb_service struct {
							Attributes   int64    `json:"::Attributes"`
							Hashcode     string   `json:"::Hashcode"`
							BreadcrumbID string   `json:"breadcrumbId"`
							Dependencies []string `json:"dependencies"`
							ServiceName  string   `json:"serviceName"`
						} `json:"ms.vss-tfs-web.header-breadcrumb-service"`
						Ms_vss_tfs_web_help_menu struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Title      string `json:"title"`
						} `json:"ms.vss-tfs-web.help-menu"`
						Ms_vss_tfs_web_make_suggestion struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Command    struct {
								Dependencies []string `json:"dependencies"`
								MethodName   string   `json:"methodName"`
								ServiceName  string   `json:"serviceName"`
							} `json:"command"`
							GroupKey  string `json:"groupKey"`
							IconProps struct {
								IconName string `json:"iconName"`
							} `json:"iconProps"`
							ID   string `json:"id"`
							Name string `json:"name"`
							Rank int64  `json:"rank"`
						} `json:"ms.vss-tfs-web.make-suggestion"`
						Ms_vss_tfs_web_market_menu struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Title      string `json:"title"`
						} `json:"ms.vss-tfs-web.market-menu"`
						Ms_vss_tfs_web_my_work_pivot struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
						} `json:"ms.vss-tfs-web.my-work-pivot"`
						Ms_vss_tfs_web_privacy_policy struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							GroupKey   string `json:"groupKey"`
							Href       string `json:"href"`
							IconProps  struct {
								IconName string `json:"iconName"`
							} `json:"iconProps"`
							ID   string `json:"id"`
							Name string `json:"name"`
							Rank int64  `json:"rank"`
						} `json:"ms.vss-tfs-web.privacy-policy"`
						Ms_vss_tfs_web_report_problem struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Command    struct {
								Dependencies []string `json:"dependencies"`
								MethodName   string   `json:"methodName"`
								ServiceName  string   `json:"serviceName"`
							} `json:"command"`
							GroupKey  string `json:"groupKey"`
							IconProps struct {
								IconName string `json:"iconName"`
							} `json:"iconProps"`
							ID   string `json:"id"`
							Name string `json:"name"`
							Rank int64  `json:"rank"`
						} `json:"ms.vss-tfs-web.report-problem"`
						Ms_vss_tfs_web_search_menu struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
						} `json:"ms.vss-tfs-web.search-menu"`
						Ms_vss_tfs_web_service_status struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							GroupKey   string `json:"groupKey"`
							Href       string `json:"href"`
							IconProps  struct {
								IconName string `json:"iconName"`
							} `json:"iconProps"`
							ID   string `json:"id"`
							Name string `json:"name"`
							Rank int64  `json:"rank"`
						} `json:"ms.vss-tfs-web.service-status"`
						Ms_vss_tfs_web_shortcut_dialog struct {
							Attributes int64    `json:"::Attributes"`
							Hashcode   string   `json:"::Hashcode"`
							Combos     []string `json:"combos"`
							Command    struct {
								Dependencies []string `json:"dependencies"`
								MethodName   string   `json:"methodName"`
								ServiceName  string   `json:"serviceName"`
							} `json:"command"`
							Description string `json:"description"`
							Group       string `json:"group"`
						} `json:"ms.vss-tfs-web.shortcut-dialog"`
						Ms_vss_tfs_web_shortcut_menu_provider struct {
							Attributes   int64    `json:"::Attributes"`
							Hashcode     string   `json:"::Hashcode"`
							Dependencies []string `json:"dependencies"`
							ServiceName  string   `json:"serviceName"`
						} `json:"ms.vss-tfs-web.shortcut-menu-provider"`
						Ms_vss_tfs_web_tfs_rest_error_handler_service struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
							} `json:"_sharedData"`
							ServiceName string `json:"serviceName"`
						} `json:"ms.vss-tfs-web.tfs-rest-error-handler-service"`
						Ms_vss_tfs_web_tfssearch_coordinator struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								Features []string `json:"features"`
								Routes   []string `json:"routes"`
							} `json:"_sharedData"`
							Dependencies []string `json:"dependencies"`
							Priority     int64    `json:"priority"`
							ServiceName  string   `json:"serviceName"`
						} `json:"ms.vss-tfs-web.tfssearch-coordinator"`
						Ms_vss_tfs_web_utilization_usagesummary_pivot struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							IconProps  struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey       string `json:"itemKey"`
							ModuleContent string `json:"moduleContent"`
							Name          string `json:"name"`
							Order         int64  `json:"order"`
						} `json:"ms.vss-tfs-web.utilization-usagesummary-pivot"`
						Ms_vss_wiki_web_wiki_collection_search_provider struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
							} `json:"_sharedData"`
							Dependencies      []string `json:"dependencies"`
							Placeholder       string   `json:"placeholder"`
							Priority          int64    `json:"priority"`
							ProviderName      string   `json:"providerName"`
							ServiceName       string   `json:"serviceName"`
							SupportedDocTypes []string `json:"supportedDocTypes"`
						} `json:"ms.vss-wiki-web.wiki-collection-search-provider"`
						Ms_vss_work_web_instant_search_page_interactive_event_subscription struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Command    struct {
								Dependencies []string `json:"dependencies"`
								MethodName   string   `json:"methodName"`
								ServiceName  string   `json:"serviceName"`
							} `json:"command"`
							EventName string `json:"eventName"`
						} `json:"ms.vss-work-web.instant-search-page-interactive-event-subscription"`
						Ms_vss_work_web_recent_workitem_activity_event_subscription struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Command    struct {
								Dependencies []string `json:"dependencies"`
								MethodName   string   `json:"methodName"`
								ServiceName  string   `json:"serviceName"`
							} `json:"command"`
							EventName string `json:"eventName"`
						} `json:"ms.vss-work-web.recent-workitem-activity-event-subscription"`
						Ms_vss_work_web_vertical_header_mywork_detailed struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								Assets       []string `json:"assets"`
								FeatureFlags []string `json:"featureFlags"`
								Routes       []string `json:"routes"`
							} `json:"_sharedData"`
							ComponentProperties struct {
								ClassName     string `json:"className"`
								IsCollapsible bool   `json:"isCollapsible"`
							} `json:"componentProperties"`
							ComponentType string   `json:"componentType"`
							Dependencies  []string `json:"dependencies"`
							ID            string   `json:"id"`
							Key           string   `json:"key"`
							Name          string   `json:"name"`
							Order         int64    `json:"order"`
						} `json:"ms.vss-work-web.vertical-header-mywork-detailed"`
						Ms_vss_work_web_work_account_settings_tab_group_new_branding struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							Name       string `json:"name"`
							Rank       int64  `json:"rank"`
						} `json:"ms.vss-work-web.work-account-settings-tab-group-new-branding"`
						Ms_vss_work_web_work_customization_admin_pivot struct {
							Attributes   int64    `json:"::Attributes"`
							Hashcode     string   `json:"::Hashcode"`
							Dependencies []string `json:"dependencies"`
							IconProps    struct {
								ClassName string `json:"className"`
								IconName  string `json:"iconName"`
							} `json:"iconProps"`
							ItemKey       string `json:"itemKey"`
							ModuleContent string `json:"moduleContent"`
							Name          string `json:"name"`
							Order         int64  `json:"order"`
						} `json:"ms.vss-work-web.work-customization-admin-pivot"`
						Ms_vss_work_web_workitem_collection_search_provider struct {
							Attributes int64  `json:"::Attributes"`
							Hashcode   string `json:"::Hashcode"`
							SharedData struct {
								FeatureFlags []string `json:"featureFlags"`
								Routes       []string `json:"routes"`
								Settings     []struct {
									Key       string `json:"key"`
									UserScope string `json:"userScope"`
								} `json:"settings"`
							} `json:"_sharedData"`
							Dependencies      []string `json:"dependencies"`
							Placeholder       string   `json:"placeholder"`
							Priority          int64    `json:"priority"`
							ProviderName      string   `json:"providerName"`
							ServiceName       string   `json:"serviceName"`
							SupportedDocTypes []string `json:"supportedDocTypes"`
						} `json:"ms.vss-work-web.workitem-collection-search-provider"`
					} `json:"contributions"`
					Providers     struct{} `json:"providers"`
					Relationships struct {
						Ms_feed_package_collection_search_provider struct {
							Parents struct {
								Ms_vss_web_collection_page bool `json:"ms.vss-web.collection-page"`
							} `json:"parents"`
						} `json:"ms.feed.package-collection-search-provider"`
						Ms_vss_admin_web_admin_hub_header_breadcrumb_service struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_component_collapsible bool `json:"ms.vss-admin-web.collection-admin-hub-component-collapsible"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.admin-hub-header-breadcrumb-service"`
						Ms_vss_admin_web_collection_admin_hub_tab_group struct {
							Children struct {
								Ms_vss_admin_web_organization_admin_aad_tab                 bool `json:"ms.vss-admin-web.organization-admin-aad-tab"`
								Ms_vss_admin_web_organization_admin_new_overview_tab        bool `json:"ms.vss-admin-web.organization-admin-new-overview-tab"`
								Ms_vss_admin_web_organization_admin_new_projects_tab        bool `json:"ms.vss-admin-web.organization-admin-new-projects-tab"`
								Ms_vss_aex_user_management_collection_user_management_pivot bool `json:"ms.vss-aex-user-management.collection-user-management-pivot"`
								Ms_vss_audit_web_collection_admin_auditing_hub              bool `json:"ms.vss-audit-web.collection-admin-auditing-hub"`
								Ms_vss_bill_web_user_admin_hub_bill_tab_new                 bool `json:"ms.vss-bill-web.user-admin-hub-bill-tab-new"`
								Ms_vss_extmgmt_web_manageExtensions_pivot                   bool `json:"ms.vss-extmgmt-web.manageExtensions-pivot"`
								Ms_vss_notifications_web_collection_admin_pivot             bool `json:"ms.vss-notifications-web.collection-admin-pivot"`
								Ms_vss_tfs_web_utilization_usagesummary_pivot               bool `json:"ms.vss-tfs-web.utilization-usagesummary-pivot"`
							} `json:"children"`
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_component_collapsible bool `json:"ms.vss-admin-web.collection-admin-hub-component-collapsible"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
						Ms_vss_admin_web_collection_admin_security_tab_group struct {
							Children struct {
								Ms_vss_admin_web_org_admin_orggroups_tab             bool `json:"ms.vss-admin-web.org-admin-orggroups-tab"`
								Ms_vss_admin_web_organization_admin_new_policies_tab bool `json:"ms.vss-admin-web.organization-admin-new-policies-tab"`
							} `json:"children"`
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_component_collapsible bool `json:"ms.vss-admin-web.collection-admin-hub-component-collapsible"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.collection-admin-security-tab-group"`
						Ms_vss_admin_web_create_hidden_org_admin_hubgroup struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_component_collapsible bool `json:"ms.vss-admin-web.collection-admin-hub-component-collapsible"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.create-hidden-org-admin-hubgroup"`
						Ms_vss_admin_web_org_admin_orggroups_tab struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_security_tab_group bool `json:"ms.vss-admin-web.collection-admin-security-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.org-admin-orggroups-tab"`
						Ms_vss_admin_web_organization_admin_aad_tab struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.organization-admin-aad-tab"`
						Ms_vss_admin_web_organization_admin_new_overview_tab struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.organization-admin-new-overview-tab"`
						Ms_vss_admin_web_organization_admin_new_policies_tab struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_security_tab_group bool `json:"ms.vss-admin-web.collection-admin-security-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.organization-admin-new-policies-tab"`
						Ms_vss_admin_web_organization_admin_new_projects_tab struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-admin-web.organization-admin-new-projects-tab"`
						Ms_vss_aex_user_management_collection_user_management_pivot struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-aex-user-management.collection-user-management-pivot"`
						Ms_vss_audit_web_collection_admin_auditing_hub struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-audit-web.collection-admin-auditing-hub"`
						Ms_vss_bill_web_user_admin_hub_bill_tab_new struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-bill-web.user-admin-hub-bill-tab-new"`
						Ms_vss_build_web_account_admin_agent_pools_tab struct {
							Children struct {
								Ms_vss_build_web_admin_agent_definitions_data_provider bool `json:"ms.vss-build-web.admin-agent-definitions-data-provider"`
								Ms_vss_build_web_agent_clouds_list_data_provider       bool `json:"ms.vss-build-web.agent-clouds-list-data-provider"`
								Ms_vss_build_web_agent_jobs_data_provider              bool `json:"ms.vss-build-web.agent-jobs-data-provider"`
								Ms_vss_build_web_agent_pool_data_provider              bool `json:"ms.vss-build-web.agent-pool-data-provider"`
								Ms_vss_build_web_agent_pool_security_data_provider     bool `json:"ms.vss-build-web.agent-pool-security-data-provider"`
								Ms_vss_build_web_elastic_pool_data                     bool `json:"ms.vss-build-web.elastic-pool-data"`
								Ms_vss_build_web_pools_admin_header_breadcrumb_service bool `json:"ms.vss-build-web.pools-admin-header-breadcrumb-service"`
							} `json:"children"`
							Parents struct {
								Ms_vss_build_web_build_release_account_settings_tab_group bool `json:"ms.vss-build-web.build-release-account-settings-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-build-web.account-admin-agent-pools-tab"`
						Ms_vss_build_web_build_release_account_settings_tab_group struct {
							Children struct {
								Ms_vss_build_web_account_admin_agent_pools_tab                bool `json:"ms.vss-build-web.account-admin-agent-pools-tab"`
								Ms_vss_build_web_build_release_queue_admin_parallel_jobs__org bool `json:"ms.vss-build-web.build-release-queue-admin-parallel-jobs--org"`
								Ms_vss_build_web_pipeline_org_settings_tab                    bool `json:"ms.vss-build-web.pipeline-org-settings-tab"`
								Ms_vss_distributed_task_oauth_configurations_pivot            bool `json:"ms.vss-distributed-task.oauth-configurations-pivot"`
								Ms_vss_releaseManagement_web_deployment_pools_pivot           bool `json:"ms.vss-releaseManagement-web.deployment-pools-pivot"`
							} `json:"children"`
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_component_collapsible bool `json:"ms.vss-admin-web.collection-admin-hub-component-collapsible"`
							} `json:"parents"`
						} `json:"ms.vss-build-web.build-release-account-settings-tab-group"`
						Ms_vss_build_web_build_release_queue_admin_parallel_jobs__org struct {
							Parents struct {
								Ms_vss_build_web_build_release_account_settings_tab_group bool `json:"ms.vss-build-web.build-release-account-settings-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-build-web.build-release-queue-admin-parallel-jobs--org"`
						Ms_vss_build_web_pipeline_org_settings_tab struct {
							Parents struct {
								Ms_vss_build_web_build_release_account_settings_tab_group bool `json:"ms.vss-build-web.build-release-account-settings-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-build-web.pipeline-org-settings-tab"`
						Ms_vss_build_web_pools_admin_header_breadcrumb_service struct {
							Parents struct {
								Ms_vss_build_web_account_admin_agent_pools_tab bool `json:"ms.vss-build-web.account-admin-agent-pools-tab"`
							} `json:"parents"`
						} `json:"ms.vss-build-web.pools-admin-header-breadcrumb-service"`
						Ms_vss_code_web_admin_account_version_control_pivot struct {
							Parents struct {
								Ms_vss_code_web_code_account_settings_tab_group bool `json:"ms.vss-code-web.code-account-settings-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-code-web.admin-account-version-control-pivot"`
						Ms_vss_code_web_code_account_settings_tab_group struct {
							Children struct {
								Ms_vss_code_web_admin_account_version_control_pivot bool `json:"ms.vss-code-web.admin-account-version-control-pivot"`
							} `json:"children"`
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_component_collapsible bool `json:"ms.vss-admin-web.collection-admin-hub-component-collapsible"`
							} `json:"parents"`
						} `json:"ms.vss-code-web.code-account-settings-tab-group"`
						Ms_vss_code_web_my_prs_flyout_tab struct {
							Parents struct {
								Ms_vss_tfs_web_my_work_pivot bool `json:"ms.vss-tfs-web.my-work-pivot"`
							} `json:"parents"`
						} `json:"ms.vss-code-web.my-prs-flyout-tab"`
						Ms_vss_code_web_repository_dropdown_menu struct {
							Parents struct {
								Ms_vss_tfs_web_vertical_header_composite bool `json:"ms.vss-tfs-web.vertical-header-composite"`
							} `json:"parents"`
						} `json:"ms.vss-code-web.repository-dropdown-menu"`
						Ms_vss_code_web_versioncontrol_collection_search_provider struct {
							Parents struct {
								Ms_vss_web_collection_page bool `json:"ms.vss-web.collection-page"`
							} `json:"parents"`
						} `json:"ms.vss-code-web.versioncontrol-collection-search-provider"`
						Ms_vss_distributed_task_oauth_configurations_pivot struct {
							Parents struct {
								Ms_vss_build_web_build_release_account_settings_tab_group bool `json:"ms.vss-build-web.build-release-account-settings-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-distributed-task.oauth-configurations-pivot"`
						Ms_vss_extmgmt_web_manageExtensions_pivot struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-extmgmt-web.manageExtensions-pivot"`
						Ms_vss_favorites_vertical_header_favorites struct {
							Parents struct {
								Ms_vss_tfs_web_my_work_pivot bool `json:"ms.vss-tfs-web.my-work-pivot"`
							} `json:"parents"`
						} `json:"ms.vss-favorites.vertical-header-favorites"`
						Ms_vss_notifications_web_collection_admin_pivot struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-notifications-web.collection-admin-pivot"`
						Ms_vss_releaseManagement_web_deployment_pools_pivot struct {
							Parents struct {
								Ms_vss_build_web_build_release_account_settings_tab_group bool `json:"ms.vss-build-web.build-release-account-settings-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-releaseManagement-web.deployment-pools-pivot"`
						Ms_vss_search_web_global_search_coordinator struct {
							Children struct {
								Ms_vss_search_web_search_coordinator bool `json:"ms.vss-search-web.search-coordinator"`
							} `json:"children"`
							Parents struct {
								Ms_vss_web_collection_page bool `json:"ms.vss-web.collection-page"`
							} `json:"parents"`
						} `json:"ms.vss-search-web.global-search-coordinator"`
						Ms_vss_storage_web_collection_admin_artifacts_tab_group struct {
							Children struct {
								Ms_vss_storage_web_org_admin_artifacts_usage_tab bool `json:"ms.vss-storage-web.org-admin-artifacts-usage-tab"`
							} `json:"children"`
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_component_collapsible bool `json:"ms.vss-admin-web.collection-admin-hub-component-collapsible"`
							} `json:"parents"`
						} `json:"ms.vss-storage-web.collection-admin-artifacts-tab-group"`
						Ms_vss_storage_web_org_admin_artifacts_usage_tab struct {
							Parents struct {
								Ms_vss_storage_web_collection_admin_artifacts_tab_group bool `json:"ms.vss-storage-web.collection-admin-artifacts-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-storage-web.org-admin-artifacts-usage-tab"`
						Ms_vss_tfs_web_code_conduct struct {
							Parents struct {
								Ms_vss_tfs_web_help_menu bool `json:"ms.vss-tfs-web.help-menu"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.code-conduct"`
						Ms_vss_tfs_web_get_started_action struct {
							Parents struct {
								Ms_vss_tfs_web_help_menu bool `json:"ms.vss-tfs-web.help-menu"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.get-started-action"`
						Ms_vss_tfs_web_get_support struct {
							Parents struct {
								Ms_vss_tfs_web_help_menu bool `json:"ms.vss-tfs-web.help-menu"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.get-support"`
						Ms_vss_tfs_web_goto_home_shortcut struct {
							Parents struct {
								Ms_vss_web_tfs_website bool `json:"ms.vss-web.tfs-website"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.goto-home-shortcut"`
						Ms_vss_tfs_web_goto_organization_admin_shortcut struct {
							Parents struct {
								Ms_vss_web_collection_page bool `json:"ms.vss-web.collection-page"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.goto-organization-admin-shortcut"`
						Ms_vss_tfs_web_goto_projects_shortcut struct {
							Parents struct {
								Ms_vss_web_tfs_website bool `json:"ms.vss-web.tfs-website"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.goto-projects-shortcut"`
						Ms_vss_tfs_web_header_breadcrumb_service struct {
							Parents struct {
								Ms_vss_tfs_web_vertical_header_breadcrumb bool `json:"ms.vss-tfs-web.vertical-header-breadcrumb"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.header-breadcrumb-service"`
						Ms_vss_tfs_web_help_menu struct {
							Children struct {
								Ms_vss_tfs_web_code_conduct           bool `json:"ms.vss-tfs-web.code-conduct"`
								Ms_vss_tfs_web_get_started_action     bool `json:"ms.vss-tfs-web.get-started-action"`
								Ms_vss_tfs_web_get_support            bool `json:"ms.vss-tfs-web.get-support"`
								Ms_vss_tfs_web_make_suggestion        bool `json:"ms.vss-tfs-web.make-suggestion"`
								Ms_vss_tfs_web_privacy_policy         bool `json:"ms.vss-tfs-web.privacy-policy"`
								Ms_vss_tfs_web_report_problem         bool `json:"ms.vss-tfs-web.report-problem"`
								Ms_vss_tfs_web_service_status         bool `json:"ms.vss-tfs-web.service-status"`
								Ms_vss_tfs_web_shortcut_menu_provider bool `json:"ms.vss-tfs-web.shortcut-menu-provider"`
							} `json:"children"`
							Parents struct {
								Ms_vss_tfs_web_vertical_header_composite     bool `json:"ms.vss-tfs-web.vertical-header-composite"`
								Ms_vss_tfs_web_vertical_header_help          bool `json:"ms.vss-tfs-web.vertical-header-help"`
								Ms_vss_web_account_admin_hub_group_component bool `json:"ms.vss-web.account-admin-hub-group-component"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.help-menu"`
						Ms_vss_tfs_web_make_suggestion struct {
							Parents struct {
								Ms_vss_tfs_web_help_menu bool `json:"ms.vss-tfs-web.help-menu"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.make-suggestion"`
						Ms_vss_tfs_web_market_menu struct {
							Parents struct {
								Ms_vss_tfs_web_vertical_header_composite bool `json:"ms.vss-tfs-web.vertical-header-composite"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.market-menu"`
						Ms_vss_tfs_web_my_work_pivot struct {
							Children struct {
								Ms_vss_code_web_my_prs_flyout_tab               bool `json:"ms.vss-code-web.my-prs-flyout-tab"`
								Ms_vss_favorites_vertical_header_favorites      bool `json:"ms.vss-favorites.vertical-header-favorites"`
								Ms_vss_work_web_vertical_header_mywork_detailed bool `json:"ms.vss-work-web.vertical-header-mywork-detailed"`
							} `json:"children"`
							Parents struct {
								Ms_vss_tfs_web_me_vertical_header bool `json:"ms.vss-tfs-web.me-vertical-header"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.my-work-pivot"`
						Ms_vss_tfs_web_privacy_policy struct {
							Parents struct {
								Ms_vss_tfs_web_help_menu bool `json:"ms.vss-tfs-web.help-menu"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.privacy-policy"`
						Ms_vss_tfs_web_report_problem struct {
							Parents struct {
								Ms_vss_tfs_web_help_menu bool `json:"ms.vss-tfs-web.help-menu"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.report-problem"`
						Ms_vss_tfs_web_search_menu struct {
							Parents struct {
								Ms_vss_tfs_web_vertical_header_composite bool `json:"ms.vss-tfs-web.vertical-header-composite"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.search-menu"`
						Ms_vss_tfs_web_service_status struct {
							Parents struct {
								Ms_vss_tfs_web_help_menu bool `json:"ms.vss-tfs-web.help-menu"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.service-status"`
						Ms_vss_tfs_web_shortcut_dialog struct {
							Parents struct {
								Ms_vss_web_tfs_website bool `json:"ms.vss-web.tfs-website"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.shortcut-dialog"`
						Ms_vss_tfs_web_shortcut_menu_provider struct {
							Parents struct {
								Ms_vss_tfs_web_help_menu bool `json:"ms.vss-tfs-web.help-menu"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.shortcut-menu-provider"`
						Ms_vss_tfs_web_tfs_rest_error_handler_service struct {
							Children struct {
								Ms_vss_tfs_web_services_content                             bool `json:"ms.vss-tfs-web.services-content"`
								Ms_vss_tfs_web_tfs_rest_error_handler_service_data_provider bool `json:"ms.vss-tfs-web.tfs-rest-error-handler-service-data-provider"`
							} `json:"children"`
							Parents struct {
								Ms_vss_web_rest_error_handler_root bool `json:"ms.vss-web.rest-error-handler-root"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.tfs-rest-error-handler-service"`
						Ms_vss_tfs_web_tfssearch_coordinator struct {
							Parents struct {
								Ms_vss_web_collection_page bool `json:"ms.vss-web.collection-page"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.tfssearch-coordinator"`
						Ms_vss_tfs_web_utilization_usagesummary_pivot struct {
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_tab_group bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							} `json:"parents"`
						} `json:"ms.vss-tfs-web.utilization-usagesummary-pivot"`
						Ms_vss_wiki_web_wiki_collection_search_provider struct {
							Parents struct {
								Ms_vss_web_collection_page bool `json:"ms.vss-web.collection-page"`
							} `json:"parents"`
						} `json:"ms.vss-wiki-web.wiki-collection-search-provider"`
						Ms_vss_work_web_instant_search_page_interactive_event_subscription struct {
							Parents struct {
								Ms_vss_work_web_workitem_collection_search_provider bool `json:"ms.vss-work-web.workitem-collection-search-provider"`
							} `json:"parents"`
						} `json:"ms.vss-work-web.instant-search-page-interactive-event-subscription"`
						Ms_vss_work_web_recent_workitem_activity_event_subscription struct {
							Parents struct {
								Ms_vss_work_web_workitem_collection_search_provider bool `json:"ms.vss-work-web.workitem-collection-search-provider"`
							} `json:"parents"`
						} `json:"ms.vss-work-web.recent-workitem-activity-event-subscription"`
						Ms_vss_work_web_vertical_header_mywork_detailed struct {
							Parents struct {
								Ms_vss_tfs_web_my_work_pivot bool `json:"ms.vss-tfs-web.my-work-pivot"`
							} `json:"parents"`
						} `json:"ms.vss-work-web.vertical-header-mywork-detailed"`
						Ms_vss_work_web_work_account_settings_tab_group_new_branding struct {
							Children struct {
								Ms_vss_work_web_work_customization_admin_pivot bool `json:"ms.vss-work-web.work-customization-admin-pivot"`
							} `json:"children"`
							Parents struct {
								Ms_vss_admin_web_collection_admin_hub_component_collapsible bool `json:"ms.vss-admin-web.collection-admin-hub-component-collapsible"`
							} `json:"parents"`
						} `json:"ms.vss-work-web.work-account-settings-tab-group-new-branding"`
						Ms_vss_work_web_work_customization_admin_pivot struct {
							Parents struct {
								Ms_vss_work_web_work_account_settings_tab_group_new_branding bool `json:"ms.vss-work-web.work-account-settings-tab-group-new-branding"`
							} `json:"parents"`
						} `json:"ms.vss-work-web.work-customization-admin-pivot"`
						Ms_vss_work_web_workitem_collection_search_provider struct {
							Children struct {
								Ms_vss_work_web_instant_search_page_interactive_event_subscription bool `json:"ms.vss-work-web.instant-search-page-interactive-event-subscription"`
								Ms_vss_work_web_recent_workitem_activity_event_subscription        bool `json:"ms.vss-work-web.recent-workitem-activity-event-subscription"`
							} `json:"children"`
							Parents struct {
								Ms_vss_web_collection_page bool `json:"ms.vss-web.collection-page"`
							} `json:"parents"`
						} `json:"ms.vss-work-web.workitem-collection-search-provider"`
					} `json:"relationships"`
					Types struct {
						Ms_vss_web_breadcrumb_item struct {
							Ms_vss_admin_web_create_hidden_org_admin_hubgroup bool `json:"ms.vss-admin-web.create-hidden-org-admin-hubgroup"`
						} `json:"ms.vss-web.breadcrumb-item"`
						Ms_vss_web_breadcrumb_provider struct {
							Ms_vss_admin_web_admin_hub_header_breadcrumb_service   bool `json:"ms.vss-admin-web.admin-hub-header-breadcrumb-service"`
							Ms_vss_build_web_pools_admin_header_breadcrumb_service bool `json:"ms.vss-build-web.pools-admin-header-breadcrumb-service"`
							Ms_vss_tfs_web_header_breadcrumb_service               bool `json:"ms.vss-tfs-web.header-breadcrumb-service"`
						} `json:"ms.vss-web.breadcrumb-provider"`
						Ms_vss_web_event_subscription struct {
							Ms_vss_work_web_instant_search_page_interactive_event_subscription bool `json:"ms.vss-work-web.instant-search-page-interactive-event-subscription"`
							Ms_vss_work_web_recent_workitem_activity_event_subscription        bool `json:"ms.vss-work-web.recent-workitem-activity-event-subscription"`
						} `json:"ms.vss-web.event-subscription"`
						Ms_vss_web_menu struct {
							Ms_vss_code_web_repository_dropdown_menu bool `json:"ms.vss-code-web.repository-dropdown-menu"`
							Ms_vss_tfs_web_help_menu                 bool `json:"ms.vss-tfs-web.help-menu"`
							Ms_vss_tfs_web_market_menu               bool `json:"ms.vss-tfs-web.market-menu"`
							Ms_vss_tfs_web_search_menu               bool `json:"ms.vss-tfs-web.search-menu"`
						} `json:"ms.vss-web.menu"`
						Ms_vss_web_menu_item struct {
							Ms_vss_tfs_web_code_conduct       bool `json:"ms.vss-tfs-web.code-conduct"`
							Ms_vss_tfs_web_get_started_action bool `json:"ms.vss-tfs-web.get-started-action"`
							Ms_vss_tfs_web_get_support        bool `json:"ms.vss-tfs-web.get-support"`
							Ms_vss_tfs_web_make_suggestion    bool `json:"ms.vss-tfs-web.make-suggestion"`
							Ms_vss_tfs_web_privacy_policy     bool `json:"ms.vss-tfs-web.privacy-policy"`
							Ms_vss_tfs_web_report_problem     bool `json:"ms.vss-tfs-web.report-problem"`
							Ms_vss_tfs_web_service_status     bool `json:"ms.vss-tfs-web.service-status"`
						} `json:"ms.vss-web.menu-item"`
						Ms_vss_web_menu_provider struct {
							Ms_vss_tfs_web_shortcut_menu_provider bool `json:"ms.vss-tfs-web.shortcut-menu-provider"`
						} `json:"ms.vss-web.menu-provider"`
						Ms_vss_web_search_coordinator struct {
							Ms_vss_search_web_global_search_coordinator bool `json:"ms.vss-search-web.global-search-coordinator"`
							Ms_vss_tfs_web_tfssearch_coordinator        bool `json:"ms.vss-tfs-web.tfssearch-coordinator"`
						} `json:"ms.vss-web.search-coordinator"`
						Ms_vss_web_search_provider struct {
							Ms_feed_package_collection_search_provider                bool `json:"ms.feed.package-collection-search-provider"`
							Ms_vss_code_web_versioncontrol_collection_search_provider bool `json:"ms.vss-code-web.versioncontrol-collection-search-provider"`
							Ms_vss_wiki_web_wiki_collection_search_provider           bool `json:"ms.vss-wiki-web.wiki-collection-search-provider"`
							Ms_vss_work_web_workitem_collection_search_provider       bool `json:"ms.vss-work-web.workitem-collection-search-provider"`
						} `json:"ms.vss-web.search-provider"`
						Ms_vss_web_service struct {
							Ms_vss_tfs_web_tfs_rest_error_handler_service bool `json:"ms.vss-tfs-web.tfs-rest-error-handler-service"`
						} `json:"ms.vss-web.service"`
						Ms_vss_web_shortcut struct {
							Ms_vss_tfs_web_goto_home_shortcut               bool `json:"ms.vss-tfs-web.goto-home-shortcut"`
							Ms_vss_tfs_web_goto_organization_admin_shortcut bool `json:"ms.vss-tfs-web.goto-organization-admin-shortcut"`
							Ms_vss_tfs_web_goto_projects_shortcut           bool `json:"ms.vss-tfs-web.goto-projects-shortcut"`
							Ms_vss_tfs_web_shortcut_dialog                  bool `json:"ms.vss-tfs-web.shortcut-dialog"`
						} `json:"ms.vss-web.shortcut"`
						Ms_vss_web_tab struct {
							Ms_vss_admin_web_org_admin_orggroups_tab                      bool `json:"ms.vss-admin-web.org-admin-orggroups-tab"`
							Ms_vss_admin_web_organization_admin_aad_tab                   bool `json:"ms.vss-admin-web.organization-admin-aad-tab"`
							Ms_vss_admin_web_organization_admin_new_overview_tab          bool `json:"ms.vss-admin-web.organization-admin-new-overview-tab"`
							Ms_vss_admin_web_organization_admin_new_policies_tab          bool `json:"ms.vss-admin-web.organization-admin-new-policies-tab"`
							Ms_vss_admin_web_organization_admin_new_projects_tab          bool `json:"ms.vss-admin-web.organization-admin-new-projects-tab"`
							Ms_vss_aex_user_management_collection_user_management_pivot   bool `json:"ms.vss-aex-user-management.collection-user-management-pivot"`
							Ms_vss_audit_web_collection_admin_auditing_hub                bool `json:"ms.vss-audit-web.collection-admin-auditing-hub"`
							Ms_vss_bill_web_user_admin_hub_bill_tab_new                   bool `json:"ms.vss-bill-web.user-admin-hub-bill-tab-new"`
							Ms_vss_build_web_account_admin_agent_pools_tab                bool `json:"ms.vss-build-web.account-admin-agent-pools-tab"`
							Ms_vss_build_web_build_release_queue_admin_parallel_jobs__org bool `json:"ms.vss-build-web.build-release-queue-admin-parallel-jobs--org"`
							Ms_vss_build_web_pipeline_org_settings_tab                    bool `json:"ms.vss-build-web.pipeline-org-settings-tab"`
							Ms_vss_code_web_admin_account_version_control_pivot           bool `json:"ms.vss-code-web.admin-account-version-control-pivot"`
							Ms_vss_code_web_my_prs_flyout_tab                             bool `json:"ms.vss-code-web.my-prs-flyout-tab"`
							Ms_vss_distributed_task_oauth_configurations_pivot            bool `json:"ms.vss-distributed-task.oauth-configurations-pivot"`
							Ms_vss_extmgmt_web_manageExtensions_pivot                     bool `json:"ms.vss-extmgmt-web.manageExtensions-pivot"`
							Ms_vss_favorites_vertical_header_favorites                    bool `json:"ms.vss-favorites.vertical-header-favorites"`
							Ms_vss_notifications_web_collection_admin_pivot               bool `json:"ms.vss-notifications-web.collection-admin-pivot"`
							Ms_vss_releaseManagement_web_deployment_pools_pivot           bool `json:"ms.vss-releaseManagement-web.deployment-pools-pivot"`
							Ms_vss_storage_web_org_admin_artifacts_usage_tab              bool `json:"ms.vss-storage-web.org-admin-artifacts-usage-tab"`
							Ms_vss_tfs_web_utilization_usagesummary_pivot                 bool `json:"ms.vss-tfs-web.utilization-usagesummary-pivot"`
							Ms_vss_work_web_vertical_header_mywork_detailed               bool `json:"ms.vss-work-web.vertical-header-mywork-detailed"`
							Ms_vss_work_web_work_customization_admin_pivot                bool `json:"ms.vss-work-web.work-customization-admin-pivot"`
						} `json:"ms.vss-web.tab"`
						Ms_vss_web_tab_group struct {
							Ms_vss_admin_web_collection_admin_hub_tab_group              bool `json:"ms.vss-admin-web.collection-admin-hub-tab-group"`
							Ms_vss_admin_web_collection_admin_security_tab_group         bool `json:"ms.vss-admin-web.collection-admin-security-tab-group"`
							Ms_vss_build_web_build_release_account_settings_tab_group    bool `json:"ms.vss-build-web.build-release-account-settings-tab-group"`
							Ms_vss_code_web_code_account_settings_tab_group              bool `json:"ms.vss-code-web.code-account-settings-tab-group"`
							Ms_vss_storage_web_collection_admin_artifacts_tab_group      bool `json:"ms.vss-storage-web.collection-admin-artifacts-tab-group"`
							Ms_vss_tfs_web_my_work_pivot                                 bool `json:"ms.vss-tfs-web.my-work-pivot"`
							Ms_vss_work_web_work_account_settings_tab_group_new_branding bool `json:"ms.vss-work-web.work-account-settings-tab-group-new-branding"`
						} `json:"ms.vss-web.tab-group"`
					} `json:"types"`
				} `json:"_contributions"`
				FeatureFlags struct {
					Azure_DevOps_AdminEngagement_OrganizationAad_TooManyDisconnectedUsersContactSupport             bool `json:"Azure.DevOps.AdminEngagement.OrganizationAad.TooManyDisconnectedUsersContactSupport"`
					AzureDevOps_ReleaseManagement_EnforceJobAuthScopeForReleases                                    bool `json:"AzureDevOps.ReleaseManagement.EnforceJobAuthScopeForReleases"`
					Build2_HostedImageDeprication                                                                   bool `json:"Build2.HostedImageDeprication"`
					Build2_Retention_FilePathArtifactsAndSymbolsDelete                                              bool `json:"Build2.Retention.FilePathArtifactsAndSymbolsDelete"`
					DistributedTask_Node6LockdownAllowed                                                            bool `json:"DistributedTask.Node6LockdownAllowed"`
					DistributedTask_PipelineBillingModel2                                                           bool `json:"DistributedTask.PipelineBillingModel2"`
					Microsoft_Azure_DevOps_Commerce_DisableIncrementalArtifactsMaxQuantityInBillingTab              bool `json:"Microsoft.Azure.DevOps.Commerce.DisableIncrementalArtifactsMaxQuantityInBillingTab"`
					Microsoft_Azure_DevOps_Commerce_DisableInstallTestManager                                       bool `json:"Microsoft.Azure.DevOps.Commerce.DisableInstallTestManager"`
					Microsoft_Azure_DevOps_Commerce_EnableAzCommIntegration                                         bool `json:"Microsoft.Azure.DevOps.Commerce.EnableAzCommIntegration"`
					Microsoft_Azure_DevOps_Commerce_EnableAzCommRead                                                bool `json:"Microsoft.Azure.DevOps.Commerce.EnableAzCommRead"`
					Microsoft_Azure_DevOps_Commerce_EnableDefaultAccessLevelSelectorForAllOrgs                      bool `json:"Microsoft.Azure.DevOps.Commerce.EnableDefaultAccessLevelSelectorForAllOrgs"`
					Microsoft_Azure_DevOps_Commerce_HideCloudBasedLoadTesting                                       bool `json:"Microsoft.Azure.DevOps.Commerce.HideCloudBasedLoadTesting"`
					NWP_Mention_PR                                                                                  bool `json:"NWP.Mention.PR"`
					NWP_Mention_People                                                                              bool `json:"NWP.Mention.People"`
					NWP_Mention_WorkItem                                                                            bool `json:"NWP.Mention.WorkItem"`
					TeamFoundationRecentActivityService_DisplayUserProjectActivities                                bool `json:"TeamFoundationRecentActivityService.DisplayUserProjectActivities"`
					VisualStudio_Service_WebPlatform_ContributedRestErrorHandlerQueueing                            bool `json:"VisualStudio.Service.WebPlatform.ContributedRestErrorHandlerQueueing"`
					VisualStudio_Service_WebPlatform_ContributedRestErrorHandlers                                   bool `json:"VisualStudio.Service.WebPlatform.ContributedRestErrorHandlers"`
					VisualStudio_Services_AdminEngagement_MeControl_MeControlTenantPicker                           bool `json:"VisualStudio.Services.AdminEngagement.MeControl.MeControlTenantPicker"`
					VisualStudio_Services_AdminEngagement_OrganizationAzureActiveDirectory_EnableOrganizationPolicy bool `json:"VisualStudio.Services.AdminEngagement.OrganizationAzureActiveDirectory.EnableOrganizationPolicy"`
					VisualStudio_Services_AdminEngagement_OrganizationAzureActiveDirectory_EnableSwitchTenant       bool `json:"VisualStudio.Services.AdminEngagement.OrganizationAzureActiveDirectory.EnableSwitchTenant"`
					VisualStudio_Services_AdminEngagement_OrganizationOverview_EnableTakeOverUI                     bool `json:"VisualStudio.Services.AdminEngagement.OrganizationOverview.EnableTakeOverUI"`
					VisualStudio_Services_AdminEngagement_OrganizationPolicies_DisallowEnableAlterCred              bool `json:"VisualStudio.Services.AdminEngagement.OrganizationPolicies.DisallowEnableAlterCred"`
					VisualStudio_Services_AdminEngagement_ProfileSettingsView_EnableRefreshPermissionsTab           bool `json:"VisualStudio.Services.AdminEngagement.ProfileSettingsView.EnableRefreshPermissionsTab"`
					VisualStudio_Services_AdminEngagement_Projects_EnableProjectRestoreTab                          bool `json:"VisualStudio.Services.AdminEngagement.Projects.EnableProjectRestoreTab"`
					VisualStudio_Services_AdminEngagement_ShowAlternateCredentialsInfo                              bool `json:"VisualStudio.Services.AdminEngagement.ShowAlternateCredentialsInfo"`
					VisualStudio_Services_AdminEngagement_ShowAlternateCredentialsInfoPublic                        bool `json:"VisualStudio.Services.AdminEngagement.ShowAlternateCredentialsInfoPublic"`
					VisualStudio_Services_Identity_AnonymousAccess                                                  bool `json:"VisualStudio.Services.Identity.AnonymousAccess"`
					VisualStudio_Services_TenantPolicy_Admin_PATs_RestrictFullScopeTokenCreationPolicy              bool `json:"VisualStudio.Services.TenantPolicy.Admin.PATs.RestrictFullScopeTokenCreationPolicy"`
					VisualStudio_Services_TenantPolicy_Admin_PATs_RestrictGlobalTokenCreationPolicy                 bool `json:"VisualStudio.Services.TenantPolicy.Admin.PATs.RestrictGlobalTokenCreationPolicy"`
					VisualStudio_Services_TenantPolicy_Admin_PATs_RestrictTokenLifespanPolicy                       bool `json:"VisualStudio.Services.TenantPolicy.Admin.PATs.RestrictTokenLifespanPolicy"`
					VisualStudio_Services_WebAccess_RestErrorHandler_FetchSignout                                   bool `json:"VisualStudio.Services.WebAccess.RestErrorHandler.FetchSignout"`
					VisualStudio_Services_WebPlatform_DisableNewAvatarImage                                         bool `json:"VisualStudio.Services.WebPlatform.DisableNewAvatarImage"`
					WebAccess_DistributedTask_AgentPool_AnalyticsTab                                                bool `json:"WebAccess.DistributedTask.AgentPool.AnalyticsTab"`
					WebAccess_DistributedTask_NewPoolAdminHub_PoolsListSignalR                                      bool `json:"WebAccess.DistributedTask.NewPoolAdminHub.PoolsListSignalR"`
					WebAccess_Mention_WorkItem_RestSearch                                                           bool `json:"WebAccess.Mention.WorkItem.RestSearch"`
					WebAccess_MyExperiences_MyWorkItems                                                             bool `json:"WebAccess.MyExperiences.MyWorkItems"`
					WebAccess_Search_Code_DisableAggregations                                                       bool `json:"WebAccess.Search.Code.DisableAggregations"`
					WebAccess_Search_Code_EnableRationalizedListOfCodeElementFilters                                bool `json:"WebAccess.Search.Code.EnableRationalizedListOfCodeElementFilters"`
					WebAccess_Search_Code_InfiniteScroll                                                            bool `json:"WebAccess.Search.Code.InfiniteScroll"`
					WebAccess_Search_EnableNewRoute                                                                 bool `json:"WebAccess.Search.EnableNewRoute"`
					WebAccess_Search_Package_InfiniteScroll                                                         bool `json:"WebAccess.Search.Package.InfiniteScroll"`
					WebAccess_Search_Wiki_InfiniteScroll                                                            bool `json:"WebAccess.Search.Wiki.InfiniteScroll"`
					WebAccess_Search_WorkItem_DisableAggregations                                                   bool `json:"WebAccess.Search.WorkItem.DisableAggregations"`
					WebAccess_Search_WorkItem_InfiniteScroll                                                        bool `json:"WebAccess.Search.WorkItem.InfiniteScroll"`
					WebAccess_Search_WorkItem_InstantSearch_Delay                                                   bool `json:"WebAccess.Search.WorkItem.InstantSearch.Delay"`
					WebAccess_Search_WorkItem_QuickNavigation                                                       bool `json:"WebAccess.Search.WorkItem.QuickNavigation"`
					WebAccess_Search_WorkItem_SearchWarmers                                                         bool `json:"WebAccess.Search.WorkItem.SearchWarmers"`
					WebAccess_SignalR_AppPool                                                                       bool `json:"WebAccess.SignalR.AppPool"`
					WebAccess_WorkItemTracking_Form_CommentService_SharedUI_Enabled                                 bool `json:"WebAccess.WorkItemTracking.Form.CommentService.SharedUI.Enabled"`
					WebAccess_WorkItemTracking_Form_Comments_QuoteReply_Enabled                                     bool `json:"WebAccess.WorkItemTracking.Form.Comments.QuoteReply.Enabled"`
					WebAccess_WorkItemTracking_Form_DarkThemeColorRemap                                             bool `json:"WebAccess.WorkItemTracking.Form.DarkThemeColorRemap"`
					WebAccess_WorkItemTracking_MentionsReplyAll                                                     bool `json:"WebAccess.WorkItemTracking.MentionsReplyAll"`
					Wiki_InstantSearch                                                                              bool `json:"Wiki.InstantSearch"`
				} `json:"_featureFlags"`
				Features struct {
					Ms_vss_admin_web_organization_admin_project_scope_users_group bool `json:"ms.vss-admin-web.organization-admin-project-scope-users-group"`
					Ms_vss_admin_web_organization_admin_settings_search_feature   bool `json:"ms.vss-admin-web.organization-admin-settings-search-feature"`
					Ms_vss_build_web_historical_graph_for_pools_feature           bool `json:"ms.vss-build-web.historical-graph-for-pools-feature"`
					Ms_vss_search_web_enable_search_dialog                        bool `json:"ms.vss-search-web.enable-search-dialog"`
					Ms_vss_tfs_web_enable_search_suggestions                      bool `json:"ms.vss-tfs-web.enable-search-suggestions"`
					Ms_vss_tfs_web_vertical_navigation                            bool `json:"ms.vss-tfs-web.vertical-navigation"`
				} `json:"_features"`
				IdentityDetails struct {
					Six8ca9068_6894_4845_8320_68a8bdea930b struct {
						DisplayName     string `json:"displayName"`
						IdentityImageID string `json:"identityImageId"`
					} `json:"68ca9068-6894-4845-8320-68a8bdea930b"`
				} `json:"_identityDetails"`
				Locations struct {
					Zero0000010_0000_8888_8000_000000000000 struct {
						Four string `json:"4"`
					} `json:"00000010-0000-8888-8000-000000000000"`
					Zero0000030_0000_8888_8000_000000000000 struct {
						Four string `json:"4"`
					} `json:"00000030-0000-8888-8000-000000000000"`
					Zero000003c_0000_8888_8000_000000000000 struct {
						Four string `json:"4"`
					} `json:"0000003c-0000-8888-8000-000000000000"`
					Nine51917ac_a960_4999_8464_e3f0aa25b381 struct {
						Four string `json:"4"`
					} `json:"951917ac-a960-4999-8464-e3f0aa25b381"`
				} `json:"_locations"`
				MenuItems struct {
					Ms_vss_tfs_web_market_menu []struct {
						GroupKey string `json:"groupKey"`
						Href     string `json:"href"`
						Key      string `json:"key"`
						Name     string `json:"name"`
					} `json:"ms.vss-tfs-web.market-menu"`
					Ms_vss_tfs_web_profile_menu []struct {
						Command struct {
							Dependencies []string `json:"dependencies"`
							MethodName   string   `json:"methodName"`
							ServiceName  string   `json:"serviceName"`
						} `json:"command"`
						CommandID     string `json:"commandId"`
						ComponentType string `json:"componentType"`
						GroupKey      string `json:"groupKey"`
						Href          string `json:"href"`
						IconProps     struct {
							ClassName string `json:"className"`
							IconName  string `json:"iconName"`
						} `json:"iconProps"`
						ItemType  int64  `json:"itemType"`
						Key       string `json:"key"`
						Name      string `json:"name"`
						SubMenuID string `json:"subMenuId"`
						Title     string `json:"title"`
					} `json:"ms.vss-tfs-web.profile-menu"`
					Ms_vss_tfs_web_search_menu []struct {
						GroupKey  string `json:"groupKey"`
						Href      string `json:"href"`
						IconProps struct {
							ClassName string `json:"className"`
							IconName  string `json:"iconName"`
						} `json:"iconProps"`
						Key  string `json:"key"`
						Name string `json:"name"`
						Rank int64  `json:"rank"`
					} `json:"ms.vss-tfs-web.search-menu"`
				} `json:"_menuItems"`
				Permissions struct {
					One01eae8c_1709_47f9_b228_0e476c35b3ba struct {
						AgentPools    int64 `json:"AgentPools"`
						AgentPools_13 int64 `json:"AgentPools/13"`
					} `json:"101eae8c-1709-47f9-b228-0e476c35b3ba"`
					Fa557b48_b5bf_458a_bb2b_1b680426fe8b struct {
						Favorites int64 `json:"/Favorites"`
					} `json:"fa557b48-b5bf-458a-bb2b-1b680426fe8b"`
				} `json:"_permissions"`
				Routes struct {
					Ms_feed_packaging_hub_route                        []string `json:"ms.feed.packaging-hub-route"`
					Ms_vss_admin_web_collection_admin_hub_route        []string `json:"ms.vss-admin-web.collection-admin-hub-route"`
					Ms_vss_admin_web_organization_admin_security_route []string `json:"ms.vss-admin-web.organization-admin-security-route"`
					Ms_vss_admin_web_project_admin_hub_route           []string `json:"ms.vss-admin-web.project-admin-hub-route"`
					Ms_vss_admin_web_user_admin_hub_route              []string `json:"ms.vss-admin-web.user-admin-hub-route"`
					Ms_vss_build_web_ci_results_hub_route              []string `json:"ms.vss-build-web.ci-results-hub-route"`
					Ms_vss_search_web_collection_route                 []string `json:"ms.vss-search-web.collection-route"`
					Ms_vss_search_web_project_route                    []string `json:"ms.vss-search-web.project-route"`
					Ms_vss_tfs_web_project_overview_route              []string `json:"ms.vss-tfs-web.project-overview-route"`
					Ms_vss_tfs_web_suite_me_page_route                 []string `json:"ms.vss-tfs-web.suite-me-page-route"`
					Ms_vss_work_web_work_customization_route           []string `json:"ms.vss-work-web.work-customization-route"`
					Ms_vss_work_web_work_items_form_route_with_id      []string `json:"ms.vss-work-web.work-items-form-route-with-id"`
				} `json:"_routes"`
				Settings struct {
					Host struct {
						Search_InstantSearch struct{} `json:"Search/InstantSearch"`
						Search_Telemetry     struct{} `json:"Search/Telemetry"`
						Search_Wit_Warmer    struct{} `json:"Search/Wit/Warmer"`
					} `json:"host"`
				} `json:"_settings"`
				Themes struct {
					Ms_vss_web_vsts_theme struct {
						Data struct {
							Background_color                                string `json:"background-color"`
							Border_subtle_color                             string `json:"border-subtle-color"`
							Build_output_error_color                        string `json:"build-output-error-color"`
							Callout_background_color                        string `json:"callout-background-color"`
							Callout_filtered_background_color               string `json:"callout-filtered-background-color"`
							Callout_shadow_color                            string `json:"callout-shadow-color"`
							Callout_shadow_secondary_color                  string `json:"callout-shadow-secondary-color"`
							Communication_background                        string `json:"communication-background"`
							Communication_foreground                        string `json:"communication-foreground"`
							Component_errorBoundary_background_color        string `json:"component-errorBoundary-background-color"`
							Component_errorBoundary_border_color            string `json:"component-errorBoundary-border-color"`
							Component_grid_action_hover_color               string `json:"component-grid-action-hover-color"`
							Component_grid_action_selected_cell_hover_color string `json:"component-grid-action-selected-cell-hover-color"`
							Component_grid_cell_bottom_border_color         string `json:"component-grid-cell-bottom-border-color"`
							Component_grid_drag_source_color                string `json:"component-grid-drag-source-color"`
							Component_grid_focus_border_color               string `json:"component-grid-focus-border-color"`
							Component_grid_link_hover_color                 string `json:"component-grid-link-hover-color"`
							Component_grid_link_selected_row_color          string `json:"component-grid-link-selected-row-color"`
							Component_grid_row_hover_color                  string `json:"component-grid-row-hover-color"`
							Component_grid_selected_row_color               string `json:"component-grid-selected-row-color"`
							Component_htmlEditor_background_color           string `json:"component-htmlEditor-background-color"`
							Component_htmlEditor_foreground_color           string `json:"component-htmlEditor-foreground-color"`
							Component_label_default_color                   string `json:"component-label-default-color"`
							Component_label_default_color_hover             string `json:"component-label-default-color-hover"`
							Component_menu_selected_item_background         string `json:"component-menu-selected-item-background"`
							Diff_color_modified                             string `json:"diff-color-modified"`
							Diff_color_original                             string `json:"diff-color-original"`
							Focus_border_color                              string `json:"focus-border-color"`
							Focus_pulse_max_color                           string `json:"focus-pulse-max-color"`
							Focus_pulse_min_color                           string `json:"focus-pulse-min-color"`
							Icon_folder_color                               string `json:"icon-folder-color"`
							Nav_header_active_item_background               string `json:"nav-header-active-item-background"`
							Nav_header_background                           string `json:"nav-header-background"`
							Nav_header_item_hover_background                string `json:"nav-header-item-hover-background"`
							Nav_header_product_color                        string `json:"nav-header-product-color"`
							Nav_header_text_disabled_color                  string `json:"nav-header-text-disabled-color"`
							Nav_header_text_primary_color                   string `json:"nav-header-text-primary-color"`
							Nav_header_text_secondary_color                 string `json:"nav-header-text-secondary-color"`
							Nav_vertical_active_group_background            string `json:"nav-vertical-active-group-background"`
							Nav_vertical_active_item_background             string `json:"nav-vertical-active-item-background"`
							Nav_vertical_background_color                   string `json:"nav-vertical-background-color"`
							Nav_vertical_item_hover_background              string `json:"nav-vertical-item-hover-background"`
							Nav_vertical_text_primary_color                 string `json:"nav-vertical-text-primary-color"`
							Nav_vertical_text_secondary_color               string `json:"nav-vertical-text-secondary-color"`
							Palette_accent1                                 string `json:"palette-accent1"`
							Palette_accent1_dark                            string `json:"palette-accent1-dark"`
							Palette_accent1_light                           string `json:"palette-accent1-light"`
							Palette_accent2                                 string `json:"palette-accent2"`
							Palette_accent2_dark                            string `json:"palette-accent2-dark"`
							Palette_accent2_light                           string `json:"palette-accent2-light"`
							Palette_accent3                                 string `json:"palette-accent3"`
							Palette_accent3_dark                            string `json:"palette-accent3-dark"`
							Palette_accent3_light                           string `json:"palette-accent3-light"`
							Palette_black_alpha_0                           string `json:"palette-black-alpha-0"`
							Palette_black_alpha_10                          string `json:"palette-black-alpha-10"`
							Palette_black_alpha_100                         string `json:"palette-black-alpha-100"`
							Palette_black_alpha_2                           string `json:"palette-black-alpha-2"`
							Palette_black_alpha_20                          string `json:"palette-black-alpha-20"`
							Palette_black_alpha_30                          string `json:"palette-black-alpha-30"`
							Palette_black_alpha_4                           string `json:"palette-black-alpha-4"`
							Palette_black_alpha_6                           string `json:"palette-black-alpha-6"`
							Palette_black_alpha_60                          string `json:"palette-black-alpha-60"`
							Palette_black_alpha_70                          string `json:"palette-black-alpha-70"`
							Palette_black_alpha_8                           string `json:"palette-black-alpha-8"`
							Palette_black_alpha_80                          string `json:"palette-black-alpha-80"`
							Palette_error                                   string `json:"palette-error"`
							Palette_error_10                                string `json:"palette-error-10"`
							Palette_error_6                                 string `json:"palette-error-6"`
							Palette_neutral_0                               string `json:"palette-neutral-0"`
							Palette_neutral_10                              string `json:"palette-neutral-10"`
							Palette_neutral_100                             string `json:"palette-neutral-100"`
							Palette_neutral_2                               string `json:"palette-neutral-2"`
							Palette_neutral_20                              string `json:"palette-neutral-20"`
							Palette_neutral_30                              string `json:"palette-neutral-30"`
							Palette_neutral_4                               string `json:"palette-neutral-4"`
							Palette_neutral_6                               string `json:"palette-neutral-6"`
							Palette_neutral_60                              string `json:"palette-neutral-60"`
							Palette_neutral_70                              string `json:"palette-neutral-70"`
							Palette_neutral_8                               string `json:"palette-neutral-8"`
							Palette_neutral_80                              string `json:"palette-neutral-80"`
							Palette_primary                                 string `json:"palette-primary"`
							Palette_primary_darken_10                       string `json:"palette-primary-darken-10"`
							Palette_primary_darken_6                        string `json:"palette-primary-darken-6"`
							Palette_primary_darkened_10                     string `json:"palette-primary-darkened-10"`
							Palette_primary_darkened_6                      string `json:"palette-primary-darkened-6"`
							Palette_primary_shade_10                        string `json:"palette-primary-shade-10"`
							Palette_primary_shade_20                        string `json:"palette-primary-shade-20"`
							Palette_primary_shade_30                        string `json:"palette-primary-shade-30"`
							Palette_primary_tint_10                         string `json:"palette-primary-tint-10"`
							Palette_primary_tint_20                         string `json:"palette-primary-tint-20"`
							Palette_primary_tint_30                         string `json:"palette-primary-tint-30"`
							Palette_primary_tint_40                         string `json:"palette-primary-tint-40"`
							Panel_shadow_color                              string `json:"panel-shadow-color"`
							Panel_shadow_secondary_color                    string `json:"panel-shadow-secondary-color"`
							Search_match_background                         string `json:"search-match-background"`
							Search_selected_match_background                string `json:"search-selected-match-background"`
							Status_error_background                         string `json:"status-error-background"`
							Status_error_foreground                         string `json:"status-error-foreground"`
							Status_error_strong                             string `json:"status-error-strong"`
							Status_error_text                               string `json:"status-error-text"`
							Status_info_background                          string `json:"status-info-background"`
							Status_info_foreground                          string `json:"status-info-foreground"`
							Status_success_background                       string `json:"status-success-background"`
							Status_success_foreground                       string `json:"status-success-foreground"`
							Status_success_text                             string `json:"status-success-text"`
							Status_warning_background                       string `json:"status-warning-background"`
							Status_warning_foreground                       string `json:"status-warning-foreground"`
							Status_warning_icon_foreground                  string `json:"status-warning-icon-foreground"`
							Status_warning_text                             string `json:"status-warning-text"`
							Text_disabled_color                             string `json:"text-disabled-color"`
							Text_on_communication_background                string `json:"text-on-communication-background"`
							Text_primary_color                              string `json:"text-primary-color"`
							Text_secondary_color                            string `json:"text-secondary-color"`
							Third_party_icon_filter                         string `json:"third-party-icon-filter"`
						} `json:"data"`
						Extends interface{} `json:"extends"`
						ID      string      `json:"id"`
						Name    string      `json:"name"`
					} `json:"ms.vss-web.vsts-theme"`
				} `json:"_themes"`
			} `json:"sharedData"`
		} `json:"dataProviders"`
	} `json:"fps"`
}
