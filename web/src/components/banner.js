import styles from '../styles/Banner.module.css';

function Banner() {
    return ( 
        <div className={styles.mainHeader}>
            <a href="https://passage.id/" target='_blank'  rel="noopener"><div className={styles.passageLogo}></div></a>
            <div className={styles.headerText}>Connect Demo </div>
            <div className={styles.spacer}></div>
            <a href="https://passage.id/" className={styles.link} target='_blank'  rel="noopener">Go to Passage</a>
        </div>
    );
}
export default Banner;
