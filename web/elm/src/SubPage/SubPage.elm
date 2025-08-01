module SubPage.SubPage exposing
    ( Model(..)
    , handleCallback
    , handleDelivery
    , handleNotFound
    , init
    , subscriptions
    , tooltip
    , update
    , urlUpdate
    , view
    )

import Application.Models exposing (Session)
import Build.Build as Build
import Build.Header.Models
import Build.Models
import Causality.Causality as Causality
import Dashboard.Dashboard as Dashboard
import Dashboard.Models
import DownloadFly.DownloadFly as DownloadFly
import DownloadFly.Model
import EffectTransformer exposing (ET)
import FlySuccess.FlySuccess as FlySuccess
import FlySuccess.Models
import Html exposing (Html)
import Job.Job as Job
import Login.Login as Login
import Message.Callback exposing (Callback(..))
import Message.Effects exposing (Effect(..))
import Message.Message exposing (Message(..))
import Message.Subscription exposing (Delivery(..), Interval(..), Subscription)
import Message.TopLevelMessage exposing (TopLevelMessage(..))
import NotFound.Model
import NotFound.NotFound as NotFound
import Pipeline.Pipeline as Pipeline
import Resource.Models
import Resource.Resource as Resource
import Routes
import Tooltip
import UpdateMsg exposing (UpdateMsg)


type Model
    = BuildModel Build.Models.Model
    | JobModel Job.Model
    | ResourceModel Resource.Models.Model
    | PipelineModel Pipeline.Model
    | NotFoundModel NotFound.Model.Model
    | DashboardModel Dashboard.Models.Model
    | FlySuccessModel FlySuccess.Models.Model
    | CausalityModel Causality.Model
    | DownloadFlyModel DownloadFly.Model.Model


init : Session -> Routes.Route -> ( Model, List Effect )
init session route =
    case route of
        Routes.Build { id, highlight } ->
            Build.init
                { highlight = highlight
                , pageType = Build.Header.Models.JobBuildPage id
                , fromBuildPage = Nothing
                }
                |> Tuple.mapFirst BuildModel

        Routes.OneOffBuild { id, highlight } ->
            Build.init
                { highlight = highlight
                , pageType = Build.Header.Models.OneOffBuildPage id
                , fromBuildPage = Nothing
                }
                |> Tuple.mapFirst BuildModel

        Routes.Resource { id, page, version } ->
            Resource.init
                { resourceId = id
                , paging = page
                , highlightVersion = version
                }
                |> Tuple.mapFirst ResourceModel

        Routes.Job { id, page } ->
            Job.init
                { jobId = id
                , paging = page
                }
                |> Tuple.mapFirst JobModel

        Routes.Pipeline { id, groups } ->
            Pipeline.init
                { pipelineLocator = id
                , turbulenceImgSrc = session.turbulenceImgSrc
                , selectedGroups = groups
                }
                |> Tuple.mapFirst PipelineModel

        Routes.Dashboard { searchType, dashboardView } ->
            Dashboard.init
                { searchType = searchType
                , dashboardView = dashboardView
                }
                |> Tuple.mapFirst DashboardModel

        Routes.FlySuccess noop flyPort ->
            FlySuccess.init
                { authToken = session.authToken
                , flyPort = flyPort
                , noop = noop
                }
                |> Tuple.mapFirst FlySuccessModel

        Routes.DownloadFly ->
            DownloadFly.init
                session.route
                |> Tuple.mapFirst DownloadFlyModel

        Routes.Causality { id, direction } ->
            if session.featureFlags.resource_causality then
                Causality.init
                    { versionId = id
                    , direction = direction
                    }
                    |> Tuple.mapFirst CausalityModel

            else
                NotFound.init { notFoundImgSrc = session.notFoundImgSrc, route = session.route }
                    |> Tuple.mapFirst NotFoundModel


handleNotFound : Session -> ET Model
handleNotFound session ( model, effects ) =
    case getUpdateMessage model of
        UpdateMsg.NotFound ->
            let
                ( newModel, newEffects ) =
                    NotFound.init { notFoundImgSrc = session.notFoundImgSrc, route = session.route }
            in
            ( NotFoundModel newModel, effects ++ newEffects )

        UpdateMsg.AOK ->
            ( model, effects )


getUpdateMessage : Model -> UpdateMsg
getUpdateMessage model =
    case model of
        BuildModel mdl ->
            Build.getUpdateMessage mdl

        JobModel mdl ->
            Job.getUpdateMessage mdl

        ResourceModel mdl ->
            Resource.getUpdateMessage mdl

        PipelineModel mdl ->
            Pipeline.getUpdateMessage mdl

        CausalityModel mdl ->
            Causality.getUpdateMessage mdl

        _ ->
            UpdateMsg.AOK


genericUpdate :
    ET Build.Models.Model
    -> ET Job.Model
    -> ET Resource.Models.Model
    -> ET Pipeline.Model
    -> ET Dashboard.Models.Model
    -> ET Causality.Model
    -> ET NotFound.Model.Model
    -> ET FlySuccess.Models.Model
    -> ET DownloadFly.Model.Model
    -> ET Model
genericUpdate fBuild fJob fRes fPipe fDash fCaus fNF fFS dFly ( model, effects ) =
    case model of
        BuildModel buildModel ->
            fBuild ( buildModel, effects )
                |> Tuple.mapFirst BuildModel

        JobModel jobModel ->
            fJob ( jobModel, effects )
                |> Tuple.mapFirst JobModel

        PipelineModel pipelineModel ->
            fPipe ( pipelineModel, effects )
                |> Tuple.mapFirst PipelineModel

        ResourceModel resourceModel ->
            fRes ( resourceModel, effects )
                |> Tuple.mapFirst ResourceModel

        DashboardModel dashboardModel ->
            fDash ( dashboardModel, effects )
                |> Tuple.mapFirst DashboardModel

        CausalityModel causalityModel ->
            fCaus ( causalityModel, effects )
                |> Tuple.mapFirst CausalityModel

        FlySuccessModel flySuccessModel ->
            fFS ( flySuccessModel, effects )
                |> Tuple.mapFirst FlySuccessModel

        NotFoundModel notFoundModel ->
            fNF ( notFoundModel, effects )
                |> Tuple.mapFirst NotFoundModel

        DownloadFlyModel downloadFlyModel ->
            dFly ( downloadFlyModel, effects )
                |> Tuple.mapFirst DownloadFlyModel


handleCallback : Callback -> Session -> ET Model
handleCallback callback session =
    genericUpdate
        (Build.handleCallback callback)
        (Job.handleCallback callback)
        (Resource.handleCallback callback session)
        (Pipeline.handleCallback callback)
        (Dashboard.handleCallback callback)
        (Causality.handleCallback callback)
        identity
        identity
        identity
        >> (case callback of
                LoggedOut (Ok ()) ->
                    genericUpdate
                        handleLoggedOut
                        handleLoggedOut
                        handleLoggedOut
                        handleLoggedOut
                        handleLoggedOut
                        handleLoggedOut
                        handleLoggedOut
                        handleLoggedOut
                        handleLoggedOut

                _ ->
                    identity
           )


handleLoggedOut : ET { a | isUserMenuExpanded : Bool }
handleLoggedOut ( m, effs ) =
    ( { m | isUserMenuExpanded = False }
    , effs
        ++ [ NavigateTo <|
                Routes.toString <|
                    Routes.Dashboard
                        { searchType = Routes.Normal ""
                        , dashboardView = Routes.ViewNonArchivedPipelines
                        }
           ]
    )


handleDelivery : Session -> Delivery -> ET Model
handleDelivery session delivery =
    genericUpdate
        (Build.handleDelivery session delivery)
        (Job.handleDelivery delivery)
        (Resource.handleDelivery session delivery)
        (Pipeline.handleDelivery delivery)
        (Dashboard.handleDelivery session delivery)
        (Causality.handleDelivery delivery)
        (NotFound.handleDelivery delivery)
        (FlySuccess.handleDelivery delivery)
        (DownloadFly.handleDelivery delivery)


update : Session -> Message -> ET Model
update session msg =
    genericUpdate
        (Login.update msg >> Build.update msg)
        (Login.update msg >> Job.update msg)
        (Login.update msg >> Resource.update msg)
        (Login.update msg >> Pipeline.update msg)
        (Login.update msg >> Dashboard.update session msg)
        (Login.update msg >> Causality.update msg)
        (Login.update msg)
        (Login.update msg >> FlySuccess.update msg)
        (Login.update msg >> DownloadFly.update msg)
        >> (case msg of
                GoToRoute route ->
                    handleGoToRoute route

                _ ->
                    identity
           )


handleGoToRoute : Routes.Route -> ET a
handleGoToRoute route ( a, effs ) =
    ( a, effs ++ [ NavigateTo <| Routes.toString route ] )


urlUpdate : Session -> Routes.Transition -> ET Model
urlUpdate session routes =
    case ( session.featureFlags.resource_causality, routes.to ) of
        ( False, Routes.Causality _ ) ->
            -- If the feature flag is disabled, you can't navigate to the
            -- causality page
            handleNotFound session

        _ ->
            urlUpdateValid routes



-- urlUpdateValid should be invoked only when we know the page can be navigated to


urlUpdateValid : Routes.Transition -> ET Model
urlUpdateValid routes =
    genericUpdate
        (case routes.to of
            Routes.Build { id, highlight } ->
                Build.changeToBuild
                    { pageType = Build.Header.Models.JobBuildPage id
                    , highlight = highlight
                    , fromBuildPage =
                        case routes.from of
                            Routes.Build params ->
                                Just <| Build.Header.Models.JobBuildPage params.id

                            _ ->
                                Nothing
                    }

            Routes.OneOffBuild { id, highlight } ->
                Build.changeToBuild
                    { pageType = Build.Header.Models.OneOffBuildPage id
                    , highlight = highlight
                    , fromBuildPage =
                        case routes.from of
                            Routes.OneOffBuild params ->
                                Just <| Build.Header.Models.OneOffBuildPage params.id

                            _ ->
                                Nothing
                    }

            _ ->
                identity
        )
        (case routes.to of
            Routes.Job { id, page } ->
                Job.changeToJob { jobId = id, paging = page }

            _ ->
                identity
        )
        (case routes.to of
            Routes.Resource { id, page, version } ->
                Resource.changeToResource { resourceId = id, paging = page, highlightVersion = version }

            _ ->
                identity
        )
        (case routes.to of
            Routes.Pipeline { id, groups } ->
                Pipeline.changeToPipelineAndGroups
                    { pipelineLocator = id
                    , selectedGroups = groups
                    }

            _ ->
                identity
        )
        (case routes.to of
            Routes.Dashboard f ->
                Dashboard.changeRoute f

            _ ->
                identity
        )
        (case routes.to of
            Routes.Causality { id, direction } ->
                Causality.changeToVersionedResource
                    { versionId = id
                    , direction = direction
                    }

            _ ->
                identity
        )
        identity
        identity
        identity


view : Session -> Model -> ( String, Html Message )
view ({ userState } as session) mdl =
    case mdl of
        BuildModel model ->
            ( Build.documentTitle model
            , Build.view session model
            )

        JobModel model ->
            ( Job.documentTitle model
            , Job.view session model
            )

        PipelineModel model ->
            ( Pipeline.documentTitle model
            , Pipeline.view session model
            )

        ResourceModel model ->
            ( Resource.documentTitle model
            , Resource.view session model
            )

        DashboardModel model ->
            ( Dashboard.documentTitle
            , Dashboard.view session model
            )

        NotFoundModel model ->
            ( NotFound.documentTitle
            , NotFound.view session model
            )

        FlySuccessModel model ->
            ( FlySuccess.documentTitle
            , FlySuccess.view userState model
            )

        DownloadFlyModel model ->
            ( DownloadFly.documentTitle
            , DownloadFly.view session model
            )

        CausalityModel model ->
            ( Causality.documentTitle model
            , Causality.view session model
            )


tooltip : Model -> Session -> Maybe Tooltip.Tooltip
tooltip mdl =
    case mdl of
        BuildModel model ->
            Build.tooltip model

        JobModel model ->
            Job.tooltip model

        PipelineModel model ->
            Pipeline.tooltip model

        ResourceModel model ->
            Resource.tooltip model

        DashboardModel _ ->
            Dashboard.tooltip

        NotFoundModel model ->
            NotFound.tooltip model

        FlySuccessModel model ->
            FlySuccess.tooltip model

        DownloadFlyModel model ->
            DownloadFly.tooltip model

        CausalityModel model ->
            Causality.tooltip model


subscriptions : Model -> List Subscription
subscriptions mdl =
    case mdl of
        BuildModel model ->
            Build.subscriptions model

        JobModel _ ->
            Job.subscriptions

        PipelineModel _ ->
            Pipeline.subscriptions

        ResourceModel model ->
            Resource.subscriptions model

        DashboardModel _ ->
            Dashboard.subscriptions

        NotFoundModel _ ->
            NotFound.subscriptions

        FlySuccessModel _ ->
            FlySuccess.subscriptions

        DownloadFlyModel _ ->
            DownloadFly.subscriptions

        CausalityModel _ ->
            Causality.subscriptions
