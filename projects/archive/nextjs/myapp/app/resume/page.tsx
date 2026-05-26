import '../app.css';
import './style.css';

export default function ResumePage() {
  return (
    <div className="Scroll-pane">
      <h2 className='education'>Education & Technical Skills</h2>
      <p>Tufts University: Medford, MA. Quantitative Economics (BS) – 2019</p>
      <p>Technologies: AWS, Azure, Kubernetes, Linux, Docker, Terraform, Ansible, Istio, Prometheus, Grafana</p>
      <p>Languages: Go, Python, Bash, SQL, TypeScript/JS, HTML, CSS</p>
      <p>Certifications: CKAD, AZ104</p>
      <p>Interests: Distributed systems, Kubernetes, Linux, networking, security, cloud, full-stack development</p>
      <h2 className='education'>Experience</h2>
      <h3>Deloitte</h3>
      <div className="details">
        Solution Specialist - Kansas City, MO – Nov 2021 - present
      </div>
      <h4>
        FDA CBER
      </h4>
      <p>• Supported services across all environments and Kubernetes clusters</p>
      <p>• Contributed to implementation of monitoring and alerting systems to enable faster incident
        response times and provide insight into application health</p>
      <p>• Contributed to app and system security via service mesh and OAuth2.0/OIDC configurations</p>
      <h4>GPS Software Factory</h4>
      <p>• Contributed to projects aiming to enable teams to stand up CI/CD infrastructure and workflows
        quickly on AWS infra. Primary tooling includes EKS, Istio, GitLab, Jenkins, ArgoCD, and a suite
        of testing tools</p>
      <p>• Contributed to design architecture, EKS configuration, security, and usage patterns. Researched
        and tested EKS best practices, implemented OAuth2.0/OIDC authentication flow, automated
        process of deploying helm charts to Kubernetes</p>
      <p>• Wrote Ansible playbooks to automate installation, configuration of EC2 hosted services
      </p>
      <h3>Allscripts Healthcare</h3>
      <div className="details">
        DevOps Engineer - Raleigh, NC. Aug 2019 - Oct 2021
      </div>
      <p>
        • Built and manage cloud infrastructure with Terraform
      </p>
      <p>
        • Built CI/CD pipelines using Azure DevOps, Azure
      </p>
      <p>
        • Contributed to adoption of modern technologies - Git, Azure, Terraform, Docker, K8s
      </p>
      <p>
        • Worked on installers to support legacy product lines
      </p>
      <p>
        • Developed service using Angular, .NET Core, & SQL Server to simplify SQL queries and
        provide data visualization for QA - (done on 4-month ext. for internship, requested switch to dev)
      </p>
      <h3>Allscripts Healthcare</h3>
      <div className="details">
        Solutions Management - Intern - Raleigh, NC. June - Aug 2019
      </div>
      <p>
        • Created spreadsheets for senior level management to track client application usage
      </p>
      <p>
        • Wrote VBA scripts to automate data organization and manipulation for existing spreadsheets
      </p>
      <h3>Economic Development Unit, Legal Aid of W. MO</h3>
      <div className="details">
        Intern - Kansas City, MO. June - Aug 2017
      </div>
      <p>
        • Worked on behalf of neighborhood associations to acquire rights to vacant properties
      </p>
      <p>
        • Used databases and legal records to create Excel spreadsheets for potential properties
      </p>
      <h2 className='education'>Volunteer</h2>
      <h3>
        Big Brothers Big Sisters of Massachusetts Bay
      </h3>
      <h4>Somerville, MA, 2015 - 2019</h4>
      <p>• Met with 'little' every Friday for mentorship</p>
      <h3>
        Care Beyond the Boulevard
      </h3>
      <h4>Kansas City, MO, 2018</h4>
      <p>• Worked with doctors and nurses to provide healthcare to people in KC area</p>  
      <div className="buffer">&nbsp;</div>
    </div>

  );
}
