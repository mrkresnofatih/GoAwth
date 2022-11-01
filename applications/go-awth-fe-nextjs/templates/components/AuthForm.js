import { css } from '@emotion/css'
import Image from 'next/image'
import React from 'react'
import goAwthLogo from '../../public/goawthlogo.png'

const AuthForm = ({subtitle, children}) => {
  return (
    <div className={authFormStyles.loginBox}>
        <Image src={goAwthLogo} alt="goAwthLogo" width={50} height={50} />
        <h3 className={authFormStyles.loginTitle}>GoAwth</h3>
        <h4 className={authFormStyles.loginSubtitle}>{subtitle}</h4>
        <div className={authFormStyles.loginDivider} />
        {children}
    </div>
  )
}

export default AuthForm

const authFormStyles = {
    loginBox: css`
        height: 100%;
        background-color: rgb(25,25,25);
        padding: 24px;
        border-top: 10px solid #4CAF50;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    `,
    loginTitle: css`
        font-size: 48px;
        font-weight: 600;
    `,
    loginSubtitle: css`
        font-size: 18px;
        font-weight: 200;
    `,
    loginDivider: css`
        border-top: 1px solid grey;
        width: 250px;
        margin-top: 20px;
    `
}