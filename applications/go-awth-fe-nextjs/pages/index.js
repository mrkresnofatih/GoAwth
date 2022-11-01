import SeoHead from "../templates/SeoHead";
import PageTemplate from "../templates/PageTemplate";
import { css, cx } from "@emotion/css";
import goAwthLogo from '../public/goawthlogo.png'
import Image from 'next/image'

export default function Home() {
  const seoHeadData = {
    title: "GoAwth | Home",
    description: "PoC OpenIDConnect & Oauth2.0 Authorization Server using Golang Echo, GORM, JWT, & MySQL"
  }
  return (
    <PageTemplate seoHead={seoHeadData}>
      <div className={homeStyles.welcomeBody}>
        <Image src={goAwthLogo} alt="goAwthLogo" width={72} height={72} />
        <p className={homeStyles.welcomeSubtitle}>Welcome to</p>
        <h1 className={homeStyles.welcomeTitle}>GoAwth</h1>
        <p className={cx(homeStyles.welcomeSubtitle, homeStyles.withBorderTop)}>PoC OpenIDConnect Identity Provider & OAuth2.0 Authorization Server using Golang Echo, GORM, JWT, MySQL</p>
      </div>
    </PageTemplate>
  )
}

const homeStyles = {
  welcomeBody: css`
    height: calc(100vh - 72px);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  `,
  welcomeSubtitle: css`
    font-size: 24px;
    font-weight: 100;
    color: lightgray;
    transform: translate(0, 25%);
    margin-top: 24px;
  `,
  welcomeTitle: css`
    font-size: 72px;
    font-weight: 600;
  `,
  withBorderTop: css`
    padding: 24px;
    border-top: 1px solid rgba(255, 255, 255, 0.2);
    margin-top: 0;
    width: 600px;
    text-align: center;
  `
}