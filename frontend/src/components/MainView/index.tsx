import React, { useCallback, useState } from "react";
import styled from "styled-components";
import { callAPI } from "../../services/api-service";
import MainViewHeader from "./MainViewHeader";
import {
  AboutThisProjectBottom,
  AboutThisProjectHeader,
} from "./AboutThisProject";
import MyYaml from "./YamlBox/ExistingYaml";
import NewYaml from "./YamlBox/NewYaml";

const TextaresContainer = styled.div`
  display: grid;
  grid-column-gap: 4rem;
  grid-row-gap: 3rem;
  align-items: self-end;
  @media (min-width: 74rem) {
    grid-template-columns: 1fr 1fr;
    justify-content: center;
    grid-auto-flow: column;
  }
  padding: 2% 8% 3% 8%;
`;

const MainViewBodyContainer = styled.div`
  overflow-y: auto;
`;

export const BrOnlyOnPc = styled.br`
  @media (max-width: 74rem) {
    display: none;
  }
`;

const MainView: React.FC = () => {
  const [existingYamlTextArea, setExistingYamlTextArea] = useState("");
  const [output, setOutput] = useState<any>();
  const [fetching, setFetching] = useState(false);
  const [err, setErr] = useState<any>();
  const [curTab, setCurTab] = useState<number>(0);

  const callApiCallabck = useCallback(
    (endpoint: string) => {
      if (!existingYamlTextArea || existingYamlTextArea === "") return;
      callAPI(
        endpoint,
        {
          hcl: existingYamlTextArea,
        },
        setOutput,
        setFetching,
        setErr
      );
    },
    [existingYamlTextArea]
  );

  return (
    <>
      <MainViewHeader />
      <MainViewBodyContainer>
        <BrOnlyOnPc />
        <TextaresContainer>
          <MyYaml
            callApiCallabck={callApiCallabck}
            setExistingYamlTextArea={setExistingYamlTextArea}
            curTab={curTab}
          />
          <NewYaml
            yamlOutput={output?.toString()}
            fetching={fetching}
            err={err}
            callApiCallabck={callApiCallabck}
            curTab={curTab}
            setCurTab={setCurTab}
          />
        </TextaresContainer>
        <BrOnlyOnPc />
      </MainViewBodyContainer>
      <AboutThisProjectHeader />
      <AboutThisProjectBottom />
    </>
  );
};

export default MainView;
