import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';
import Link from "@docusaurus/Link";

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: ReactNode;
  link: string;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Principles',
    Svg: require('@site/static/img/principles.svg').default,
    description: (
      <>
          The principles we adhere to when developing dsenv.
      </>
    ),
    link: '/docs/category/principles/',
  },
  {
    title: 'Architecture',
    Svg: require('@site/static/img/architecture.svg').default,
    description: (
      <>
          Understand the architecture of dsenv.
      </>
    ),
    link: '/docs/category/architecture/',
  },
  {
    title: 'User Guide',
    Svg: require('@site/static/img/guide.svg').default,
    description: (
      <>
          How to interact with dsenv via the cli.
      </>
    ),
    link: '/docs/category/user-guide/',
  },
];

function Feature({title, Svg, description, link}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
      <div className="text--center padding-horiz--md">
          <Link
              className="button button--secondary button--lg"
              to={link}>
              Read more
          </Link>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
