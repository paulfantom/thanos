import React, { FC, useState, Fragment } from 'react';
import { Alert, Collapse, Table, Badge } from 'reactstrap';
import { RuleStatus } from './AlertContents';
import { Rule } from '../../types/types';
import { faChevronDown, faChevronRight } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { createExternalExpressionLink } from '../../utils/index';

interface CollapsibleAlertPanelProps {
  rule: Rule;
  showAnnotations: boolean;
}

const alertColors: RuleStatus<string> = {
  firing: 'danger',
  pending: 'warning',
  inactive: 'success',
};

const CollapsibleAlertPanel: FC<CollapsibleAlertPanelProps> = ({ rule, showAnnotations }) => {
  const [open, toggle] = useState(false);

  return (
    <>
      <Alert fade={false} onClick={() => toggle(!open)} color={alertColors[rule.state]} style={{ cursor: 'pointer' }}>
        <FontAwesomeIcon icon={open ? faChevronDown : faChevronRight} fixedWidth />
        <strong>{rule.name}</strong> ({`${rule.alerts.length} active`})
      </Alert>
      <Collapse isOpen={open} className="mb-2">
        <pre className="alert-cell">
          <code>
            <div>
              name: <a href={createExternalExpressionLink(`ALERTS{alertname="${rule.name}"}`)}>{rule.name}</a>
            </div>
            <div>
              expr: <a href={createExternalExpressionLink(rule.query)}>{rule.query}</a>
            </div>
            <div>
              <div>labels:</div>
              <div className="ml-5">severity: {rule.labels.severity}</div>
            </div>
            <div>
              <div>annotations:</div>
              <div className="ml-5">summary: {rule.annotations.summary}</div>
            </div>
          </code>
        </pre>
        {rule.alerts.length > 0 && (
          <Table bordered size="sm">
            <thead>
              <tr>
                <th>Labels</th>
                <th>State</th>
                <th>Active Since</th>
                <th>Value</th>
              </tr>
            </thead>
            <tbody>
              {rule.alerts.map((alert, i) => {
                return (
                  <Fragment key={i}>
                    <tr>
                      <td style={{ verticalAlign: 'middle' }}>
                        {Object.entries(alert.labels).map(([k, v], j) => {
                          return (
                            <Badge key={j} color="primary" className="mr-1">
                              {k}={v}
                            </Badge>
                          );
                        })}
                      </td>
                      <td>
                        <h5 className="m-0">
                          <Badge color={alertColors[alert.state] + ' text-uppercase'} className="px-3">
                            {alert.state}
                          </Badge>
                        </h5>
                      </td>
                      <td>{alert.activeAt}</td>
                      <td>{alert.value}</td>
                    </tr>
                    {showAnnotations && <Annotations annotations={alert.annotations} />}
                  </Fragment>
                );
              })}
            </tbody>
          </Table>
        )}
      </Collapse>
    </>
  );
};

interface AnnotationsProps {
  annotations: Record<string, string>;
}

export const Annotations: FC<AnnotationsProps> = ({ annotations }) => {
  return (
    <Fragment>
      <tr>
        <td colSpan={4}>
          <h5 className="font-weight-bold">Annotations</h5>
        </td>
      </tr>
      <tr>
        <td colSpan={4}>
          {Object.entries(annotations).map(([k, v], i) => {
            return (
              <div key={i}>
                <strong>{k}</strong>
                <div>{v}</div>
              </div>
            );
          })}
        </td>
      </tr>
    </Fragment>
  );
};

export default CollapsibleAlertPanel;
