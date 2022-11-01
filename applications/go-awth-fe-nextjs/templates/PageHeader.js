import { css, cx } from '@emotion/css'
import Image from 'next/image'
import React from 'react'
import RouterLink from './components/RouterLink'
import goAwthLogo from '../public/goawthlogo.png'

const PageHeader = ({title}) => {
  return (
    <div className={pageHeaderStyles.body}>
        <div className={pageHeaderStyles.innerBody}>
            <div className={pageHeaderStyles.leftContainer}>
                <Image src={goAwthLogo} alt="goAwthLogo" width={36} height={36} />
                <RouterLink href={"/"}>
                    <label className={pageHeaderStyles.headerMenuButton}>Home</label>
                </RouterLink>
                <RouterLink href={"/developer"}>
                    <label className={pageHeaderStyles.headerMenuButton}>Developer</label>
                </RouterLink>
            </div>
            <RouterLink href={"/player"}>
                <label className={cx(pageHeaderStyles.headerMenuButton, pageHeaderStyles.loginMenuButton)}>Login / Signup</label>
            </RouterLink>
        </div>
    </div>
  )
}

export default PageHeader

const pageHeaderStyles = {
    body: css`
        width: 100%;
        display: flex;
        background-color: rgb(25,25,25);
    `,
    innerBody: css`
        flex: 1;
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 18px;
    `,
    leftContainer: css`
        display: flex;
    `,
    headerMenuButton: css`
        font-size: 22px;
        font-weight: 500;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-left: 18px;

        :hover {
            cursor: pointer;
            opacity: 0.9;
        }
    `,
    loginMenuButton: css`
        font-size: 16px;
        margin: 0 8px;
        background-color: #4CAF50;
        border-radius: 4px;
        padding: 8px 12px;
    `
}