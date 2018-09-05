import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';
import IAccount from './account';
import IOrganizationProfile from './organization-profile';

export default class IOrganization extends IAccount {
	@prop({ required: true })
	preferred_name: string;

	@prop({ required: true })
	is_verified: boolean;

	@prop()
	opportunities?: Array<string>;

	@prop({ required: true })
	org_info: IOrganizationProfile;

	@prop()
	experience_validations?: Array<[string, string]>
}
