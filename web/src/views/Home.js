import "@passageidentity/passage-elements/passage-auth";
import styles from '../styles/App.module.css';

function Home() {
    return (
        <>
            <div className={styles.demoTitle}><a href="https://hashnode.com/hackathons/1password" target='_blank' rel="noopener">Hackathon By 1Password</a></div>
            <div>
                <passage-auth app-id={process.env.REACT_APP_PASSAGE_APP_ID}></passage-auth>
            </div>
        </>
    );
}

export default Home;
