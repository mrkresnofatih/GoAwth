import { css } from '@emotion/css'
import axios from 'axios'
import Image from 'next/image'
import React, { useEffect, useState } from 'react'
import AppButton from '../templates/components/AppButton'
import AuthForm from '../templates/components/AuthForm'
import TextInput from '../templates/components/TextInput'
import PageTemplate from '../templates/PageTemplate'

const player = () => {
    const [login, setLogin] = useState(false)

    const [username, setUsername] = useState("")
    const [fullName, setFullName] = useState("")
    const [password, setPassword] = useState("")
    const [imageUrl, setImageUrl] = useState("")

    const [playerToken, setPlayerToken] = useState("")

    const resetState = () => {
        setUsername("")
        setFullName("")
        setPassword("")
        setImageUrl("")
        setPlayerToken("")
    }

    const seoHeadData = {
        title: "GoAwth | Player Portal",
        description: "PoC OpenIDConnect & Oauth2.0 Authorization Server using Golang Echo, GORM, JWT, & MySQL"
    }
    const authFormSubtitle = (login) => login ? "Login to Player Portal." : "Signup to Player Portal."

    const LoginModeLink = ({login}) => {
        if (login) {
            return (
                <label onClick={() => setLogin(false)} className={playerStyles.loginAlt}>Don't have an account? <p className={playerStyles.loginAltBold}>Signup</p></label>
            )
        }
        return (
            <label onClick={() => setLogin(true)} className={playerStyles.loginAlt}>Already have an account? <p className={playerStyles.loginAltBold}>Login</p></label>
        )
    }

    return (
        <PageTemplate seoHead={seoHeadData}>
            {playerToken === "" ? (
                <div className={playerStyles.body}>
                    <div className={playerStyles.documentationBody}></div>
                    <div className={playerStyles.loginBody}>
                        <AuthForm subtitle={authFormSubtitle(login)}>
                            {login ? (
                                <>
                                    <TextInput
                                        password={false}
                                        placeholder="Username"
                                        style={{width: 225}}
                                        onChange={(e) => setUsername(e.target.value)}
                                        value={username}
                                    />
                                    <TextInput
                                        password={true}
                                        placeholder="Password"
                                        style={{width: 225}}
                                        onChange={(e) => setPassword(e.target.value)}
                                        value={password}
                                    />
                                    <AppButton 
                                        text={"Login"} 
                                        style={{marginTop: 20, width: 215}} 
                                        onClick={() => {
                                            loginApi({username, password}, (response) => {
                                                const {accessToken} = response;
                                                console.log(accessToken)
                                                setPlayerToken(accessToken)
                                            })
                                        }}
                                    />
                                </>
                            ): (
                                <>
                                    <TextInput
                                        password={false}
                                        placeholder="Fullname"
                                        style={{width: 225}}
                                        onChange={(e) => setFullName(e.target.value)}
                                        value={fullName}
                                    />
                                    <TextInput
                                        password={false}
                                        placeholder="ImageUrl"
                                        style={{width: 225}}
                                        onChange={(e) => setImageUrl(e.target.value)}
                                        value={imageUrl}
                                    />
                                    <TextInput
                                        password={false}
                                        placeholder="Username"
                                        style={{width: 225}}
                                        onChange={(e) => setUsername(e.target.value)}
                                        value={username}
                                    />
                                    <TextInput
                                        password={true}
                                        placeholder="Password"
                                        style={{width: 225}}
                                        onChange={(e) => setPassword(e.target.value)}
                                        value={password}
                                    />
                                    <AppButton 
                                        text={"Signup"} 
                                        style={{marginTop: 20, width: 215}} 
                                        onClick={() => {
                                            signupApi({username, password, fullName, imageUrl}, (response) => {
                                                resetState();
                                                setLogin(true);
                                            })
                                        }}
                                    />
                                </>
                            )}
                            <LoginModeLink login={login}/>
                        </AuthForm>
                    </div>
                </div>
            ) : (
                <PlayerPortal accessToken={playerToken} />
            )}
        </PageTemplate>
    )
}

export default player

const playerStyles = {
    body: css`
        height: calc(100vh - 72px);
        display: flex;
    `,
    documentationBody: css`
        display: flex;
        flex-direction: column;
        padding: 24px;
        flex: 1;
        padding: 24px;
    `,
    loginBody: css`
        padding: 24px;
        display: flex;
        flex-direction: column;
        width: 400px;
    `,
    loginAlt: css`
        margin-top: 20px;
        font-size: 14px;
        font-weight: 200;
        display: flex;

        :hover {
            font-style: italic;
            cursor: pointer;
        }
    `,
    loginAltBold: css`
        font-weight: 600;
        color: #4CAF50;
        margin-left: 4px;
    `
}

const loginApi = ({username, password}, callback) => {
    axios.post("http://localhost:1323/player/login", {
        username,
        password
    }).then((response) => {
        console.log(response.data)
        callback(response.data.data)
    }).catch((response) => {
        console.log(response)
    })
}

const signupApi = ({username, fullName, password, imageUrl}, callback) => {
    axios.post("http://localhost:1323/player/sign-up", {
        username,
        fullName,
        password,
        imageUrl
    }).then((response) => {
        console.log(response.data)
        callback(response.data.data)
    }).catch((response) => {
        console.log(response)
    })
}

const PlayerPortal = ({accessToken}) => {
    const [playerData, setPlayerData] = useState(undefined)

    useEffect(() => {
        getProfileApi({token: accessToken}, (data) => {
            setPlayerData(data)
        })
    }, [accessToken])

    if (playerData === undefined) {
        return (
            <div className={playerPortalStyles.body}></div>
        )
    }

    return (
        <div className={playerPortalStyles.body}>
            <img
                src={playerData.imageUrl}
                alt="playerImg"
                style={{borderRadius: 100, width: 80, height: 80}}
            />
            <h5 className={playerPortalStyles.playerTitle}>{playerData.username}</h5>
            <h5 className={playerPortalStyles.playerSubtitle}>{playerData.fullName}</h5>
        </div>
    )
}

const playerPortalStyles = {
    body: css`
        height: calc(100vh - 72px);
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    `,
    playerTitle: css`
        font-size: 24px;
        font-weight: 600;
    `,
    playerSubtitle: css`
        font-size: 18px;
        font-weight: 200;
    `,
}

const getProfileApi = ({token}, callback) => {
    axios.get("http://localhost:1323/player/get-my-profile", {
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response) => {
        console.log(response.data)
        callback(response.data.data)
    }).catch((response) => {
        console.log(response)
    })
}